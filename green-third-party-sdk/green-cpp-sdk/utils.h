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

#ifndef ALIYUNCORE_UTILS_H_
#define ALIYUNCORE_UTILS_H_
#include <string>
#include <sstream>
#include <map>
#include <openssl/hmac.h>
#include <openssl/sha.h>
#include <openssl/evp.h>
#include <openssl/md5.h>
#include <cassert>
#include <sys/time.h>
#include "utils.h"
#include "Request.h"

using namespace std;

namespace AliYunCore{
	string Base64Encode(const unsigned char* Data,int DataByte);
	string Base64Decode(const char* Data,int DataByte,int& OutByte);
	string MakeRand();
	string UTCTime();
	string UrlEncode(const string &str);
	string UrlDecode(const string &str);
	string EncodeValue(const string &str);
	string MakeQueryStr(map<string,string> &container);
	int HashHMAC(const char * algo,const char * key, unsigned int key_length,const char * input, unsigned int input_length,unsigned char * &output, unsigned int &output_length);
	void EncodeParamValue(map<string,string> &container);
	int CurlCallBack(void *buffer, size_t size, size_t nmemb, void *stream);
	unsigned char ToHex(unsigned char x);
	unsigned char FromHex(unsigned char x);
	string getMD5(const string &str);
	string getSignature(const Request &request, const string &str);

	int HashHMAC(const char * algo,const char * key, unsigned int key_length,const char * input, unsigned int input_length,unsigned char * &output, unsigned int &output_length) {
		const EVP_MD * engine = NULL;
		if(strcasecmp("sha512", algo) == 0) {
			engine = EVP_sha512();
		}
		else if(strcasecmp("sha256", algo) == 0) {
			engine = EVP_sha256();
		}
		else if(strcasecmp("sha1", algo) == 0) {
			engine = EVP_sha1();
		}
		else if(strcasecmp("md5", algo) == 0) {
			engine = EVP_md5();
		}
		else if(strcasecmp("sha224", algo) == 0) {
			engine = EVP_sha224();
		}
		else if(strcasecmp("sha384", algo) == 0) {
			engine = EVP_sha384();
		}
		else if(strcasecmp("sha", algo) == 0) {
			engine = EVP_sha1();
		}

		output = new unsigned char[EVP_MAX_MD_SIZE]();

		HMAC_CTX *ctx = HMAC_CTX_new();
		HMAC_Init_ex(ctx, key, strlen(key), engine, NULL);
		HMAC_Update(ctx, (unsigned char*)input, strlen(input));

		HMAC_Final(ctx, output, &output_length);

		HMAC_CTX_free(ctx);

		return output_length;
	}

	string Base64Encode(const unsigned char* Data,int DataByte) {
		const char EncodeTable[]="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
		string strEncode;
		unsigned char Tmp[4]={0};
		int LineLength=0;
		for(int i=0;i<(int)(DataByte / 3);i++) {
			Tmp[1] = *Data++;
			Tmp[2] = *Data++;
			Tmp[3] = *Data++;
			strEncode+= EncodeTable[Tmp[1] >> 2];
			strEncode+= EncodeTable[((Tmp[1] << 4) | (Tmp[2] >> 4)) & 0x3F];
			strEncode+= EncodeTable[((Tmp[2] << 2) | (Tmp[3] >> 6)) & 0x3F];
			strEncode+= EncodeTable[Tmp[3] & 0x3F];
			if(LineLength+=4,LineLength==76) {strEncode+="\r\n";LineLength=0;}
		}
		int Mod=DataByte % 3;
		if(Mod==1) {
			Tmp[1] = *Data++;
			strEncode+= EncodeTable[(Tmp[1] & 0xFC) >> 2];
			strEncode+= EncodeTable[((Tmp[1] & 0x03) << 4)];
			strEncode+= "==";
		}
		else if(Mod==2) {
			Tmp[1] = *Data++;
			Tmp[2] = *Data++;
			strEncode+= EncodeTable[(Tmp[1] & 0xFC) >> 2];
			strEncode+= EncodeTable[((Tmp[1] & 0x03) << 4) | ((Tmp[2] & 0xF0) >> 4)];
			strEncode+= EncodeTable[((Tmp[2] & 0x0F) << 2)];
			strEncode+= "=";
		}

		return strEncode;
	}
	string Base64Decode(const char* Data,int DataByte,int& OutByte) {
		const char DecodeTable[] = {
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				62, // '+'
				0, 0, 0,
				63, // '/'
				52, 53, 54, 55, 56, 57, 58, 59, 60, 61, // '0'-'9'
				0, 0, 0, 0, 0, 0, 0,
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, // 'A'-'Z'
				0, 0, 0, 0, 0, 0,
				26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
				39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, // 'a'-'z'
		};
		string strDecode;
		int nValue;
		int i= 0;
		while (i < DataByte) {
			if (*Data != '\r' && *Data!='\n') {
				nValue = DecodeTable[(int)*Data++] << 18;
				nValue += DecodeTable[(int)*Data++] << 12;
				strDecode+=(nValue & 0x00FF0000) >> 16;
				OutByte++;
				if (*Data != '=') {
					nValue += DecodeTable[(int)*Data++] << 6;
					strDecode+=(nValue & 0x0000FF00) >> 8;
					OutByte++;
					if (*Data != '=') {
						nValue += DecodeTable[(int)*Data++];
						strDecode+=nValue & 0x000000FF;
						OutByte++;
					}
				}
				i += 4;
			} else {
				// ignore \r or \n
				Data++;
				i++;
			}
		}
		return strDecode;
	}

	string MakeRand() {
		struct timeval tv;
		gettimeofday(&tv,NULL);
		size_t milliseconds = tv.tv_sec * 1000 + tv.tv_usec / 1000;
		srand((unsigned)milliseconds);
		ostringstream os;
		os << milliseconds << rand();
		return os.str();
	}

	int CurlCallBack(void *buffer, size_t size, size_t nmemb, void *stream) {
		size_t stDataLen = size * nmemb;
		string *pStr = (string*) stream;
		(*pStr) += ((char *) buffer);
		return stDataLen;
	}

	string UTCTime() {
		time_t t = time(NULL);
		char szBuf[128]={0};
		strftime( szBuf , 127 , "%a, %d %b %Y %H:%M:%S GMT" , gmtime( &t ) );
		return szBuf;
	}

	unsigned char ToHex(unsigned char x) {
		return  x > 9 ? x + 55 : x + 48;
	}

	unsigned char FromHex(unsigned char x) {
		unsigned char y;
		if (x >= 'A' && x <= 'Z') y = x - 'A' + 10;
		else if (x >= 'a' && x <= 'z') y = x - 'a' + 10;
		else if (x >= '0' && x <= '9') y = x - '0';
		else assert(0);
		return y;
	}

	string UrlEncode(const string &str) {
		string strTemp = "";
		size_t length = str.length();
		for (size_t i = 0; i < length; i++) {
			if (isalnum((unsigned char)str[i]) || (str[i] == '-') || (str[i] == '_') || (str[i] == '.') || (str[i] == '~')) {
				strTemp += str[i];
			} else if (str[i] == ' ') {
				strTemp += "+";
			} else{
				strTemp += '%';
				strTemp += ToHex((unsigned char)str[i] >> 4);
				strTemp += ToHex((unsigned char)str[i] % 16);
			}
		}
		return strTemp;
	}

	string UrlDecode(const string &str) {
		string strTemp = "";
		size_t length = str.length();
		for (size_t i = 0; i < length; i++) {
			if (str[i] == '+') {
				strTemp += ' ';
			} else if (str[i] == '%') {
				assert(i + 2 < length);
				unsigned char high = FromHex((unsigned char)str[++i]);
				unsigned char low = FromHex((unsigned char)str[++i]);
				strTemp += high*16 + low;
			}
			else strTemp += str[i];
		}
		return strTemp;
	}

	string EncodeValue(const string &str){
		string encStr = UrlEncode(str);
		string::size_type startpos = 0;
		while (startpos!= string::npos) {
			startpos = encStr.find('+');
			if( startpos != string::npos ) {
				encStr.replace(startpos,1,"%20");
			}
		}
		return encStr;
	}

	void EncodeParamValue(map<string,string> &container){
		for(map<string,string>::iterator map_it=container.begin();map_it!=container.end();++map_it){
			map_it->second = EncodeValue(map_it->second);
		}
		return;
	}

	string MakeQueryStr(map<string,string> &container){
		string queryStr;
		for(map<string,string>::iterator map_it=container.begin();map_it!=container.end();++map_it){
			queryStr +=  map_it->first+"="+map_it->second+"&";
		}
		return queryStr.substr(0,queryStr.length()-1);
	}

	string getMD5(const string &str){
		const char *s = str.c_str();
		unsigned char out[16];

		MD5_CTX ctx;
		MD5_Init(&ctx);
		MD5_Update(&ctx, str.c_str(), strlen(s));
		MD5_Final(out, &ctx);

		return Base64Encode(out, 16);
	}

	string getSignature(const Request &request, const string &str){
		string aks = request.accessKeySecret;
		const char *key = aks.c_str();
		unsigned char * mac = NULL;
		unsigned int mac_length = 0;
		HashHMAC("sha1", key, strlen(key), str.c_str(), str.length(), mac, mac_length);
		string signature = Base64Encode(reinterpret_cast<const unsigned char*>(mac), mac_length);
		delete [] mac;
		return signature;
	}
}
#endif /* ALIYUNCORE_UTILS_H_ */
