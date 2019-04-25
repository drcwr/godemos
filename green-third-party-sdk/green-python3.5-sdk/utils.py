#!/usr/bin/python3
# coding=utf-8

import hashlib
import time
import base64
import hmac
import json
from hashlib import sha1


def signature_md5(body):
    md5 = hashlib.md5()
    md5.update(body)
    return str(base64.encodebytes(md5.digest()), encoding="utf-8")[:-1]


def gmt_time():
    return time.strftime("%a, %d %b %Y %H:%M:%S GMT", time.gmtime(time.time()))


def signature_hmac(data, key):
    h = hmac.new(key, data, sha1)
    return str(base64.encodebytes(h.digest()), encoding="utf-8")[:-1]


def signature_header(header, path, client_info, key):
    str_list = ["POST", "application/json", header.get("Content-MD5"), "application/json", header.get("Date"), "x-acs-signature-method:HMAC-SHA1",
                "x-acs-signature-nonce:" + header.get("x-acs-signature-nonce"), "x-acs-signature-version:1.0", "x-acs-version:2017-01-12", path + "?clientInfo=" + json.dumps(client_info)]
    to_sign_str = "\n".join(str_list)

    return signature_hmac(to_sign_str.encode("utf-8"), key.encode("utf-8"))