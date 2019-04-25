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

#ifndef ALIYUNCORE_BASEREQUEST_H_
#define ALIYUNCORE_BASEREQUEST_H_

#include <string>
#include <map>
#include "Biz.h"

using namespace std;


namespace AliYunCore{
	class Request{
	public:
		string httpMethod;
		string protocol;
		string domain;
		string path;
		string version;
		string accessKeyId;
		string accessKeySecret;
		string signatureMethod;
		string signatureVersion;
		ClientInfo clientInfo;
	public:
		Request();
		virtual void getClassType(){};
	};

	Request::Request() {
		this->httpMethod = "POST";
		this->protocol = "HTTP";
		this->domain = "green.cn-shanghai.aliyuncs.com";
		this->version = "2017-01-12";
		this->signatureMethod = "HMAC-SHA1";
		this->signatureVersion = "1.0";
	}


	class ImageRequest : public Request {
	public:
        BizData bizData;
	public:
		ImageRequest(BizData bizData){
			// 默认调用鉴黄同步接口
			this->bizData = bizData;
			this->path = "/green/image/scan";
			this->clientInfo = ClientInfo();
		}

		ImageRequest(BizData bizData, string path){
			this->bizData = bizData;
			this->path = path;
			this->clientInfo = ClientInfo();
		}

		ImageRequest(BizData bizData, string path, ClientInfo clientInfo){
			this->bizData = bizData;
			this->path = path;
			this->clientInfo = clientInfo;
		}

        void getClassType(){};
	};

    class TextRequest : public Request {
    public:
        BizData bizData;
    public:
        TextRequest(BizData bizData){
            // 默认调用鉴黄同步接口
            this->bizData = bizData;
            this->path = "/green/image/scan";
            this->clientInfo = ClientInfo();
        }

        TextRequest(BizData bizData, string path){
            this->bizData = bizData;
            this->path = path;
            this->clientInfo = ClientInfo();
        }

        TextRequest(BizData bizData, string path, ClientInfo clientInfo){
            this->bizData = bizData;
            this->path = path;
            this->clientInfo = clientInfo;
        }

        void getClassType(){};
    };

}
#endif /* ALIYUNCORE_BASEREQUEST_H_ */
