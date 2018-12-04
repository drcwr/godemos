package main
import (
	"os"
	"fmt"
	// "errors"
	// "time"

	"io/ioutil"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

	// const accessKeyId string =""								
	// const accessSecret string =""	
	// const region_QD string="cn-qingdao"										//华北 1	青岛	
	// const region_BJ string="cn-beijing"										//华北 2	北京
	// const region_ZJK string="cn-zhangjiakou"									//华北 3	张家口
	// const region_HHHT string="cn-huhehaote"									//华北 5	呼和浩特
	// const region_HZ string="cn-hangzhou"										//华东 1	杭州
	// const region_SH string="cn-shanghai"										//华东 2	上海
	// const region_SZ string="cn-shenzhen"										//华南 1	深圳

	type Conf struct {
		AccessKeyId			string `yaml:"accessKeyId"`
		AccessSecret 		string `yaml:"accessSecret"`
	}


	
	var accessKeyId string 
	var accessSecret string
	func main() {
		conf :=Conf{}
		path:="accesskeyconf.json"

		rangeList := []string{"cn-qingdao","cn-beijing","cn-zhangjiakou","cn-huhehaote","cn-hangzhou","cn-shanghai","cn-shanghai","cn-shenzhen"}
		// rangeList := []string{"cn-beijing"}

		if len(os.Args)>2 {
			fmt.Printf("Error：参数错误，请按照便准输入:onekeyreset [path]")
			return
		}
		
		if len(os.Args)==2 {
			_, err := os.Stat(os.Args[1])
			if err!=nil {
				fmt.Printf("Error：配置文件路径错误，请输入正确路径！")
				return
			}

			path = os.Args[1]
		}

		yamlFile, err := ioutil.ReadFile(path)
		err = json.Unmarshal(yamlFile, &conf)
        if err != nil {
            fmt.Printf("Error：%v",err)
        }

		accessKeyId=conf.AccessKeyId
		accessSecret=conf.AccessSecret
		fmt.Println(accessKeyId,accessSecret)

		for _, item := range rangeList {
			// fmt.Println(item)
			GetInstanceStatus(item)
		}

	}


	func GetInstanceStatus(region string ) (error){

		client, err := sdk.NewClientWithAccessKey(region, accessKeyId, accessSecret)
		if err != nil {
			panic(err)
		}
	


		request := requests.NewCommonRequest()
		request.Method = "POST"
		request.Domain = "ecs.aliyuncs.com"
		request.Version = "2014-05-26"
		request.ApiName = "DescribeInstances"
		request.QueryParams["RegionId"] =region
			
		response, err := client.ProcessCommonRequest(request)
		if err != nil {
			return err
		} 
		
		InstanceStatuses := make(map[string]interface{})
		err = json.Unmarshal([]byte(response.GetHttpContentString()), &InstanceStatuses)
		if err != nil {
			return err
		} 
		// fmt.Println(InstanceStatuses)
		// for _, data := range InstanceStatuses["Instances"] {
		// 	fmt.Printf("%v\n",data)
		// }

		for _, data := range InstanceStatuses["Instances"].(map[string]interface{})["Instance"].([]interface{}) {	
			//data := InstanceStatuses["Instances"].(map[string]interface{})["Instance"].([]interface{})[0]
			// fmt.Printf("%v\n",data)
			// fmt.Println("\n\n","\"",data.(map[string]interface{})["Status"])
			tabs := "\t\t\t\t"
			fmt.Printf("%s{\n",tabs)
			fmt.Printf("%s\"role\":\"%s\",\n",tabs,"xx")
			fmt.Printf("%s\"region\":\"%s\",\n",tabs,region)
			fmt.Printf("%s\"PublicIp\":\"%s\",\n",tabs,data.(map[string]interface{})["PublicIpAddress"].(map[string]interface{})["IpAddress"].([]interface {})[0])
			fmt.Printf("%s\"InstanceID\":\"%s\",\n",tabs,data.(map[string]interface{})["InstanceId"])
			fmt.Printf("%s\"ImageId\":\"%s\"\n",tabs,data.(map[string]interface{})["ImageId"])
			fmt.Printf("%s},\n",tabs)
			// fmt.Println(data.(map[string]interface{})["ImageId"])
			// fmt.Println(data.(map[string]interface{})["InstanceId"])
			// fmt.Println(data.(map[string]interface{})["InstanceId"])
			// fmt.Println(data.(map[string]interface{})["PublicIpAddress"])
		}
		// status:=data.(map[string]interface{})["Status"]
		// LockReason :=data.(map[string]interface{})["OperationLocks"].(map[string]interface{})["LockReason"].([]interface{})
		// fmt.Printf("===%v====\n",data)
		return nil
		

	}

	

	func PathExists(path string) (bool, error) {
		_, err := os.Stat(path)
		if err == nil {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
