using System;
using Model;
using AliYunCore;
using Utils;
using Service;

namespace greenImageCsharpsdk
{
    class MainClass
    {
        public static void Test()
        {
            // create your profile 
            DefaultProfile profile = new DefaultProfile("cn-shanghai", "<your access key id>", "<your access key secret>");
            AliYunClient client = new AliYunClient(profile);

            ClientInfo clientInfo = new ClientInfo();
            clientInfo.ip = "127.0.0.1";

            // image scan task
            Task imageTask = new ImageTask("https://xxx.png");
            imageTask.dataId = Guid.NewGuid().ToString();
            BizData imageBizData = new BizData("Image", new String[] { "porn" }, new Task[] { imageTask });

            string algoPath = "/green/image/scan";
            Request req = new Request(imageBizData, algoPath);
            req.addQueryParameter("clientInfo", JSON.stringify(clientInfo));

            string response = client.getResponse(req);

            // your biz code
            Console.WriteLine(response);

            Console.ReadKey();
        }
    }
}
