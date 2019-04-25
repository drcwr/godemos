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
 
#ifndef ALIYUNOTS_OTSSERVICE_H_
#define ALIYUNOTS_OTSSERVICE_H_

#include <string>
#include <map>
#include <list>
#include "utils.h"
#include "DefaultProfile.h"
#include "Request.h"
#include "ApiBase.h"

using namespace AliYunCore;
using namespace std;

namespace AliGreen {

    class GreenClient : public ApiBase {
        private:
            DefaultProfile &Profile;
        public:
            GreenClient(DefaultProfile &pf);
            string getResponse(Request* request);
            ~GreenClient(){};
    };

    GreenClient::GreenClient(DefaultProfile &pf):Profile(pf){}

    string GreenClient::getResponse(Request* request){
        if(this->Profile.RegionId == "cn-shanghai"){
            request->domain = "green.cn-shanghai.aliyuncs.com";
        }
        request->accessKeyId = this->Profile.AccessKeyId;
        request->accessKeySecret = this->Profile.AccessKeySecret;
        return ApiBase::getResponse(request);
    }




}

#endif // ALIYUNOTS_OTSSERVICE_H_
