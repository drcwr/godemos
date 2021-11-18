// a
package main

import (
	"log"
	"unsafe"
)

type LogEventStruct struct {
	Did           string      `form:"did"`              //设备编号
	Ddid          string      `form:"ddid"`             //设备编号
	Eid           string      `form:"eid"`              //事件ID
	Iso           string      `form:"iso"`              //国家编码
	P             string      `form:"p"`                //包名
	V             string      `form:"v"`                //sdk 版本
	Ec            string      `form:"ec"`               //错误编码
	Rt            string      `form:"rt"`               //请求耗时
	Luid          string      `form:"luid"`             //本地广告位
	As            string      `form:"as"`               //广告网络名称
	Asu           string      `form:"asu"`              //广告网络 unit id
	If            string      `form:"if"`               //是否填充
	Lt            string      `form:"lt"`               //加载时长
	Iar           string      `form:"iar"`              //是否有广告位ready
	Ic            string      `form:"ic"`               //是否播放完成
	Pr            string      `form:"pr"`               //播放事件比例
	Pt            string      `form:"pt"`               //播放时长
	Tt            string      `form:"tt"`               //广告总时长
	Suuid         string      `form:"ssuid"`            //唯一标识
	Emsg          string      `form:"Emsg"`             //错误信息
	Cf            string      `from:"cf"`               //配置来源
	Os            string      `form:"os"`               // '1-android，2-ios',
	AppId         string      `form:"appId"`            //广告位ID
	Bi            string      `from:"bi"`               //bidding info 格式：adSourceid_(adsource_placement_id) (广告网络ID_adsource_placement_id )例如:1_2040
	RequestId     string      `form:"requestId"`        //请求ID
	Bt            string      `from:"bt"`               //bidding time
	Ecpm          string      `form:"ecpm"`             //正价ecpm 数据
	MsgNum        int         `form:"msgNum"`           //消息次数
	APID          interface{} `form:"apid"`             //adsource_placement_id , 兼容android和ios字段类型不一致
	Ct            int64       `form:"ct"`               //创建时间
	ScenesId      string      `form:"scid" json:"scid"` //广告场景id
	Createtime    string      `json:"createtime"`       //创建时间
	AppVer        string      `json:"app_ver"`          //应用版本号
	SdkVer        string      `json:"sdk_ver"`          //SDK版本号, 可填多个SDK版本号，用英文逗号分隔
	DeviceIdfa    string      `json:"device_idfa"`      //idfa,  可填多个
	DeviceOaid    string      `json:"device_oaid"`      //oaid,  可填多个
	DeviceIdfv    string      `json:"device_idfv"`      //idfv,  可填多个
	DeviceGaid    string      `json:"device_gaid"`      //gaid,  可填多个
	DeviceOsv     string      `json:"device_osv"`       //系统版本
	DeviceType    string      `json:"device_type"`      //设备类型
	DeviceMake    string      `json:"device_make"`      //设备制造商
	DeviceContype string      `json:"device_contype"`   //网络连接类型
	UserId        string      `json:"user_id"`          //自定义用户ID
	UserAge       string      `json:"user_age"`         //年龄
	UserGender    string      `json:"user_gender"`      //性别
	Channel       string      `json:"channel"`          //渠道
	SubChannel    string      `json:"sub_channel"`      //子渠道
	SegmentId     string      `json:"segment_id"`       //流量分组ID
	BucketId      string      `json:"bucket_id"`        //AB测试分组ID
	Tpguid        string      `json:"tpguid"`           //TradPlus guid
	Ft            string      `json:"ft"`               //firstLoadTime
	Nbr           int         `json:"Nbr"`              //bidding 返回状态
	Lc            string      `json:"Lc"`               //bidding losscode
	UseTime       int64       `json:"use_time"`         //使用时长
	AttStatus     int         `json:"att_status"`       //苹果ATT状态(0：用户未决定；1：受限制的；2：拒绝的；3：授权的)
	Time          int64       //当前时间戳
	Cb            []cb        `json:"cb,omitempty"` //埋点预聚合（只聚合单次上报 700，800 事件）
}

type cb struct {
	Eid  string      `form:"eid"`  //事件ID
	Ec   string      `form:"ec"`   //错误编码
	As   string      `form:"as"`   //广告网络名称
	Asu  string      `form:"asu"`  //广告网络 unit id
	If   string      `form:"if"`   //是否填充
	Lt   string      `form:"lt"`   //加载时长
	APID interface{} `form:"apid"` //adsource_placement_id , 兼容android和ios字段类型不一致
}

func main() {
	log.Println("main")

	ch := make(chan LogEventStruct, 100)

	log.Printf("ch size=%d\n", unsafe.Sizeof(ch))

	ev := LogEventStruct{}

	log.Printf("ev size=%d\n", unsafe.Sizeof(ev))

}
