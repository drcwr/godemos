#!/usr/bin/python3
# coding=utf-8

import uuid
from urllib import parse
from utils import *

import http.client


class GreenService:
    def __init__(self, profile):
        self.profile = profile

    def send_request(self, api_path, client_info, post_data):
        data = json.dumps(post_data)
        content_md5 = signature_md5(data.encode("utf-8"))

        gmt = gmt_time()

        header = {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Content-MD5": content_md5,
            "Date": gmt,
            "x-acs-version": "2017-01-12",
            "x-acs-signature-nonce": str(uuid.uuid1()),
            "x-acs-signature-version": "1.0",
            "x-acs-signature-method": "HMAC-SHA1",
        }

        header["Authorization"] = "acs " + self.profile.access_key_id + ":" + signature_header(header, api_path, client_info, self.profile.access_key_secret)

        query = {
            "clientInfo": json.dumps(client_info)
        }
        query_data = parse.urlencode(query).encode('utf-8')

        conn = http.client.HTTPConnection("green.cn-shanghai.aliyuncs.com", 80)
        conn.request('POST', api_path + "?%s" % str(query_data, encoding="utf-8"), json.dumps(post_data), header)
        response = conn.getresponse()
        data = response.read().decode('utf-8')

        return data
