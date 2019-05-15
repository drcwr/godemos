package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {

	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("oss.New Error:", err)
		// os.Exit(-1)
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("service.Client.Bucket Error:", err)
		// os.Exit(-1)
	}

	// Case 7: List object with paging , with prefix and max keys; return 2 items each time.
	pre := oss.Prefix("201905")
	marker := oss.Marker("")
	for {
		lor, err := bucket.ListObjects(oss.MaxKeys(600), marker, pre)
		if err != nil {
			HandleError(err)
		}
		pre = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		fmt.Println(getObjectsFormResponse(lor))
		if !lor.IsTruncated {
			break
		}

	}

}

func HandleError(err error) {
	fmt.Println("occurred error:", err)
	os.Exit(-1)
}

func getObjectsFormResponse(lor oss.ListObjectsResult) string {
	var output string
	for _, object := range lor.Objects {
		output += object.Key + "  "
	}
	return output
}
