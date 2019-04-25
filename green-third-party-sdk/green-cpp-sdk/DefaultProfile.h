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

#ifndef ALIYUNCORE_DEFAULTPROFILE_H_
#define ALIYUNCORE_DEFAULTPROFILE_H_

#include <string>
#include "DefaultProfile.h"
using namespace std;

namespace AliYunCore{
	class DefaultProfile {
	public:
		string RegionId;
		string AccessKeyId;
		string AccessKeySecret;
	public:
		DefaultProfile();
		DefaultProfile(const string &RegionId, const string &AccessKeyId, const string &AccessKeySecret);
		void setRegionId(const string &RegionId);
		void setAccessKeyId(const string &AccessKeyId);
		void setAccessKeySecret(const string &AccessKeySecret);
	};

	DefaultProfile::DefaultProfile(){

	}

	DefaultProfile::DefaultProfile(const string &RegionId, const string &AccessKeyId, const string &AccessKeySecret):RegionId(RegionId),AccessKeyId(AccessKeyId),AccessKeySecret(AccessKeySecret) {
		// TODO Auto-generated constructor stub

	}

	void DefaultProfile::setRegionId(const string &RegionId){
		this->RegionId = RegionId;
	}

	void DefaultProfile::setAccessKeyId(const string &AccessKeyId){
		this->AccessKeyId = AccessKeyId;
	}

	void DefaultProfile::setAccessKeySecret(const string &AccessKeySecret){
		this->AccessKeySecret = AccessKeySecret;
	}
}
#endif /* ALIYUNCORE_DEFAULTPROFILE_H_ */
