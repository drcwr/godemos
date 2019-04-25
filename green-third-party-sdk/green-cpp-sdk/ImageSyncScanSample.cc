#include <iostream>
#include "GreenService.h"

using namespace std;
using namespace AliGreen;
using namespace AliYunCore;

int main()
{
    DefaultProfile profile( "cn-shanghai", "<your access key id>",
        "<your access key secret>" );

    GreenClient client(profile);

    string algoPath = "/green/image/scan";

    ClientInfo clientInfo = ClientInfo();
    clientInfo.addClientInfoItem("ip", "127.0.0.1");
    clientInfo.addClientInfoItem("os", "Android");


    list<string> scenes;
    scenes.push_back("porn");
    scenes.push_back("ad");


    list<Task*> tasks;
    Task* task1 = new ImageTask("https://xxx.png");
    task1->setDataId(MakeRand());

    tasks.push_back(task1);

    BizData bizData = BizData("Test", scenes, tasks);

    Request* req = new ImageRequest(bizData, algoPath, clientInfo);

    string response = client.getResponse(req);

    // your biz code
    cout << "res:" << response << endl;

    return 0;
}
