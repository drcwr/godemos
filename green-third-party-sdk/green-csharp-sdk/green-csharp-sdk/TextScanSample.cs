using System;
using Model;
using AliYunCore;
using Utils;
using Service;

namespace greenTextCsharpSdk
{
    class MainClass
    {
        public static void Main()
        {
            // create your profile 
            DefaultProfile profile = new DefaultProfile("cn-shanghai", "<your access key id>", "<your access key secret>");
            AliYunClient client = new AliYunClient(profile);

            ClientInfo clientInfo = new ClientInfo();
            clientInfo.ip = "127.0.0.1";

            // text scan task
            string content = "测试";
			Task textTask = new TextTask(content);

            textTask.dataId = Guid.NewGuid().ToString();
            textTask.time = DateTime.Now.Millisecond;
            BizData textBizData = new BizData("Text", new String[] { "antispam" }, new Task[] { textTask });
			

            string algoPath = "/green/text/scan";
            Request req = new Request(textBizData, algoPath);
            req.addQueryParameter("clientInfo", JSON.stringify(clientInfo));

            string response = client.getResponse(req);

            // your biz code
            Console.WriteLine(response);

            Console.ReadKey();
        }
    }
}
