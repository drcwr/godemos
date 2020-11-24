package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	filename       string            // 配置文件名
	data           map[string]string // kv 配置
	lastModefyTime int64             // 上一次更新时间
	rwLock         sync.RWMutex      // 读写互斥锁
	notifyList     []Notifyer        // 观察者模式
}

type Notifyer interface {
	CallBack(*Config)
}

type AppConf struct {
	hostAddr string
	port     int
}

type AppConfMgr struct {
	config atomic.Value
}

var appConfMgr = &AppConfMgr{}

func main() {
	log.Println("main reloadconfig")
	// NewConfig("test.cfg")
	InitConf("test.cfg")
	run()
	time.Sleep(1 * time.Second)
}

func NewConfig(file string) (conf *Config, err error) {
	conf = &Config{
		filename: file,
		data:     make(map[string]string, 1024),
	}

	f, err := os.Open(conf.filename)
	if err != nil {
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return
	}

	conf.lastModefyTime = fileInfo.ModTime().Unix()

	log.Printf("NewConfig,conf:%#v\n", conf)

	m, err := conf.parse()
	if err != nil {
		log.Printf("NewConfig,parse error:%s\n", err.Error())
	}

	conf.rwLock.Lock()
	conf.data = m
	conf.rwLock.Unlock()

	log.Printf("NewConfig,conf:%#v\n", conf)

	go conf.reload()

	return
}

func (c *Config) parse() (m map[string]string, err error) {
	m = make(map[string]string, 1024)

	f, err := os.Open(c.filename)
	if err != nil {
		log.Printf("parse error:%s\n", err.Error())
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var lineNo int
	for {
		line, errReader := reader.ReadString('\n')
		if errReader == io.EOF {
			lineParse(&lineNo, &line, &m)
			break
		}
		if errReader != nil {
			log.Printf("parse,readstring err:%s\n", err.Error())
			return
		}
		lineParse(&lineNo, &line, &m)
	}

	return
}

func lineParse(lineNo *int, line *string, m *map[string]string) {
	*lineNo++

	l := strings.TrimSpace(*line)
	if len(l) == 0 || l[0] == '\n' || l[0] == '#' || l[0] == ';' {
		log.Printf("lineParse,TrimSpace line invalid,line=%d\n", *lineNo)
		return
	}

	itemSlice := strings.Split(l, "=")
	// log.Printf("lineParse,Split line %s,s=%#v,len=%d\n", l, itemSlice, len(itemSlice))
	if len(itemSlice) != 2 {
		log.Printf("lineParse,Split line invalid,line=%d\n", *lineNo)
		return
	}

	key := strings.TrimSpace(itemSlice[0])
	value := strings.TrimSpace(itemSlice[1])

	if len(key) == 0 {
		log.Printf("key len=%d\n", len(key))
		return
	}
	(*m)[key] = value

}

func (c *Config) GetValue(key string) (value string, ok bool) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	value, ok = c.data[key]
	return
}

func (c *Config) GetString(key string) (value string, err error) {
	value, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("key [%s] not found", key)
	}
	return
}

func (c *Config) GetInt(key string) (value int, err error) {
	return
}

func (c *Config) GetIntDefault(key string, defaultValue int) (value int, err error) {
	return
}

func (c *Config) GetStringDefault(key, defaultValue string) (value string, err error) {
	return
}

func (c *Config) reload() {
	log.Printf("Config reload...\n")
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		func() {
			f, err := os.Open(c.filename)
			if err != nil {
				return
			}
			defer f.Close()

			fileInfo, err := f.Stat()
			if err != nil {
				return
			}

			curModifyTime := fileInfo.ModTime().Unix()
			if curModifyTime > c.lastModefyTime {
				m, err := c.parse()
				if err != nil {
					return
				}

				c.rwLock.Lock()
				c.data = m
				c.rwLock.Unlock()

				c.lastModefyTime = curModifyTime

				for _, n := range c.notifyList {
					n.CallBack(c)
				}
			}
		}()

	}
}

func (c *Config) AddObserver(n Notifyer) {
	c.notifyList = append(c.notifyList, n)
}

func (a *AppConfMgr) CallBack(c *Config) {
	appConf := &AppConf{}

	hostAddr, err := c.GetString("hostaddr")
	// if
	_ = err
	appConf.hostAddr = hostAddr

	appConfMgr.config.Store(appConf)
}

func InitConf(filename string) {
	conf, err := NewConfig(filename)
	// if
	_ = err
	log.Printf("InitConf conf=%#v\n", conf)
	conf.AddObserver(appConfMgr)
	appConfMgr.CallBack(conf)

}

func run() {
	for {
		appConf := appConfMgr.config.Load().(*AppConf)
		log.Printf("run appConf=%#v\n", appConf)
		time.Sleep(1 * time.Second)
	}
}
