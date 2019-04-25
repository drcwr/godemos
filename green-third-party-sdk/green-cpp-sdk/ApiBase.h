/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#ifndef ALIYUNCORE_APIBASE_H_
#define ALIYUNCORE_APIBASE_H_

#include <cstring>
#include <iostream>
#include <string>
#include <map>
#include <curl/curl.h>
#include <openssl/md5.h>
#include "utils.h"
#include "Request.h"

using namespace std;

namespace AliYunCore{
	class ApiBase{
	public:
		string sendRequest(AliYunCore::Request &request, string &finalUrl, string &postFields);
		string toSignHeader(Request &request, string contentMD5, string gmtTime, string rand);
		virtual string getResponse(Request* request);
		virtual ~ApiBase(){}
	};

	/**
	 *	get the reponse of aliyun according the request
	 */
	string ApiBase::getResponse(Request* request){
        string url = request->protocol+"://" + request->domain;
        url += request->path;
        url += "?clientInfo=" + UrlEncode(request->clientInfo.toJSONString());

        if(typeid(*request).name() == typeid(ImageRequest).name()){
            ImageRequest* imageRequest = dynamic_cast<ImageRequest*>(request);
            string postData = imageRequest->bizData.toJSONString();
			return this->sendRequest(*imageRequest, url, postData);
		} else if(typeid(*request).name() == typeid(TextRequest).name()){
            TextRequest* textRequest = dynamic_cast<TextRequest*>(request);
            string postData = textRequest->bizData.toJSONString();
            return this->sendRequest(*textRequest, url, postData);
        } else{
			cout << "request type error!" << endl;
			return NULL;
		}
	}

	string ApiBase::sendRequest(Request &request, string &url, string &postData){
		CURL *curl;
		string response("");

		string contentMD5 = getMD5(postData);
		string gmtTime = UTCTime();
		string rand = MakeRand();

		string toSignstr = toSignHeader(request, contentMD5, gmtTime, rand);
		string signature = getSignature(request, toSignstr);

		struct curl_slist *headers = NULL;
		headers = curl_slist_append(headers, "Accept:application/json");
		headers = curl_slist_append(headers, "Content-Type:application/json");
		headers = curl_slist_append(headers, ("Content-MD5:" + contentMD5).c_str());
		headers = curl_slist_append(headers, ("Date:" + gmtTime).c_str());
		headers = curl_slist_append(headers, ("x-acs-version:" + request.version).c_str());
		headers = curl_slist_append(headers, ("x-acs-signature-nonce:" + rand).c_str());
		headers = curl_slist_append(headers, ("x-acs-signature-version:" + request.signatureVersion).c_str());
		headers = curl_slist_append(headers, ("x-acs-signature-method:" + request.signatureMethod).c_str());
		headers = curl_slist_append(headers, ("Authorization:acs " + request.accessKeyId + ":" + signature).c_str());

		curl = curl_easy_init();    // 初始化
		if (curl) {
			curl_easy_setopt( curl, CURLOPT_CUSTOMREQUEST, request.httpMethod.c_str());
			if( request.httpMethod == "POST" ){
				curl_easy_setopt(curl, CURLOPT_POSTFIELDS, postData.c_str());
			}
			curl_easy_setopt( curl, CURLOPT_CONNECTTIMEOUT, 30L);
			curl_easy_setopt( curl, CURLOPT_HTTPHEADER, headers);// 改协议头
			curl_easy_setopt( curl, CURLOPT_URL, url.c_str() );

			//跳过服务器SSL验证，不使用CA证书
			curl_easy_setopt( curl, CURLOPT_SSL_VERIFYPEER, 0L);

			//验证服务器端发送的证书，默认是 2(高)，1（中），0（禁用）
			curl_easy_setopt( curl, CURLOPT_SSL_VERIFYHOST, 0L);

			curl_easy_setopt( curl, CURLOPT_WRITEFUNCTION, CurlCallBack);
			curl_easy_setopt( curl, CURLOPT_WRITEDATA, &response);

			curl_easy_perform(curl);// 执行
		}

		curl_slist_free_all(headers);
		curl_easy_cleanup(curl);
		return response;
	}

	string ApiBase::toSignHeader(Request &request, string contentMD5, string gmtTime, string rand){
		string toSignStr = "POST\n";
		toSignStr.append("application/json\n").append(contentMD5 + "\n").append("application/json\n").append(gmtTime + '\n').append("x-acs-signature-method:HMAC-SHA1\n");
		toSignStr.append("x-acs-signature-nonce:").append(rand).append("\n").append("x-acs-signature-version:1.0\n").append("x-acs-version:2017-01-12\n");
		toSignStr.append(request.path + "?clientInfo=" + request.clientInfo.toJSONString());
		return toSignStr;
	}


}
#endif /* ALIYUNCORE_APIBASE_H_ */
