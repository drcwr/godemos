项目使用到的开源库：
1. openssl:   主要使用其MD5 HMAC-SHA1加密算法
2. libcurl:   HTTP请求客户端
3. rapidJson: C++ json序列化与反序列化工具（http://rapidjson.org/zh-cn/index.html）

注意事项：
1. 本项目在使用cmake编译配置工具，主要设置了openssl与curl链接过程中的库文件，要确保测试环境中已经安装好openssl和curl
2. 请求样例中只调用了图片检测的同步鉴黄接口，如果需要调用其他接口，需要设置该接口的path；如果发现请求参数错误，可以直接修改Request.h Biz.h头文件中参数定义类
3. 将ImageSyncScanSample.cc中的profile的accessKeyId、accessKeySecret替换自己的账户AK