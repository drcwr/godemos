
#ifndef GREEN_C_SDK_BIZ_H
#define GREEN_C_SDK_BIZ_H

#include <string>
#include <map>
#include <list>
#include "rapidjson/document.h"
#include "rapidjson/writer.h"
#include "rapidjson/prettywriter.h"
#include "rapidjson/stringbuffer.h"

using namespace std;
using namespace rapidjson;

class ClientInfo{
public:
    map<string, string> clientInfoMap;
public:
    ClientInfo(){}

    void addClientInfoItem(string key, string value){
        this->clientInfoMap[key] = value;
    }

    string toJSONString(){
        Document d;
        d.SetObject();
        Document::AllocatorType& allocator = d.GetAllocator();

        for(map<string,string>::const_iterator iter=this->clientInfoMap.begin();iter!=this->clientInfoMap.end();iter++){
            string a = iter->first;
            Value key(a.c_str(), allocator);

            string b = iter->second;
            Value value(b.c_str(), allocator);

            d.AddMember(key, value, allocator);
        }

        StringBuffer buffer;
        Writer<StringBuffer> writer(buffer);
        d.Accept(writer);

        return buffer.GetString();
    }

};


class Task{
public:
    ClientInfo clientInfo;
    string dataId;
    long long time;
public:
    const ClientInfo &getClientInfo() const {
        return clientInfo;
    }

    void setClientInfo(const ClientInfo &clientInfo) {
        Task::clientInfo = clientInfo;
    }

    const string &getDataId() const {
        return dataId;
    }

    void setDataId(const string &dataId) {
        Task::dataId = dataId;
    }

    long long int getTime() const {
        return time;
    }

    void setTime(long long int time) {
        Task::time = time;
    }

    virtual void getClassType(){}
};

class ImageTask : public Task{
public:
    string url;

public:
    ImageTask(string url){
        this->url = url;
    }

    void getClassType(){}
};

class TextTask : public Task{
public:
    string content;

public:
    TextTask(string content){
        this->content = content;
    }

    void getClassType(){}
};

class BizData{
public:
    string bizType;
    list<string> scenes;
    list<Task*> tasks;

public:
    BizData(){}

    BizData(string bizType, list<string> scenes, list<Task*> tasks){
        this->bizType = bizType;
        this->scenes = scenes;
        this->tasks = tasks;
    }

    string toJSONString(){
        Document d;
        d.SetObject();
        Document::AllocatorType& allocator = d.GetAllocator();

        Value value(this->bizType.c_str(), allocator);
        d.AddMember(StringRef("bizType"), value, allocator);

        Value scenes(kArrayType);
        list<string>::iterator scenesIterator = this->scenes.begin();
        while(scenesIterator != this->scenes.end()){
            string scene = *scenesIterator++;
            Value sceneValue(scene.c_str(), allocator);
            scenes.PushBack(sceneValue, allocator);
        }
        d.AddMember(StringRef("scenes"), scenes, allocator);

        Value tasks(kArrayType);
        list<Task*>::iterator tasksIterator = this->tasks.begin();
        while(tasksIterator != this->tasks.end()){
            Task* task = *tasksIterator++;
            if(typeid(*task).name() == typeid(ImageTask).name()){
                ImageTask*  imageTask = dynamic_cast<ImageTask*>(task);
                Document documentTask;
                documentTask.SetObject();

                Value url(imageTask->url.c_str(), documentTask.GetAllocator());
                documentTask.AddMember(StringRef("url"), url, documentTask.GetAllocator());

                Value dataId(imageTask->dataId.c_str(), documentTask.GetAllocator());
                documentTask.AddMember(StringRef("dataId"), dataId, documentTask.GetAllocator());

                tasks.PushBack(documentTask, allocator);
            }

            if(typeid(*task).name() == typeid(TextTask).name()){
                TextTask*  textTask = dynamic_cast<TextTask*>(task);
                Document documentTask;
                documentTask.SetObject();

                Value url(textTask->content.c_str(), documentTask.GetAllocator());
                documentTask.AddMember(StringRef("content"), url, documentTask.GetAllocator());

                Value dataId(textTask->dataId.c_str(), documentTask.GetAllocator());
                documentTask.AddMember(StringRef("dataId"), dataId, documentTask.GetAllocator());

                tasks.PushBack(documentTask, allocator);
            }

        }
        d.AddMember(StringRef("tasks"), tasks, allocator);

        StringBuffer buffer;
        Writer<StringBuffer> writer(buffer);
        d.Accept(writer);

        return buffer.GetString();
    }

};




#endif //GREEN_C_SDK_BIZ_H
