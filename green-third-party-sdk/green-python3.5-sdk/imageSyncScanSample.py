#!/usr/bin/python3
# coding=utf-8
from profile import *
from service import *

access_key_id = "<your access key id>"
access_key_secret = "<your access key secret>"
profile = Profile(access_key_id, access_key_secret)

api_path = r'/green/image/scan'

client_info = {
    "ip": "127.0.0.1"
}

post_data = {
    "bizType": "Green",
    "scenes": ["porn"],
    "tasks": [
        {
            "dataId": str(uuid.uuid1()),
            "url": "https://xxx.png"
        }
    ]
}

service = GreenService(profile)
response = service.send_request(api_path, client_info, post_data)

# your biz code
print(response)


