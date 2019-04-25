var uuidV1 = require('uuid/v1');
var greenNodejs = require('./green-nodejs-invoker.js');

const accessKeyId = '<your accessKeyId>';
const accessKeySecret = '<your accessKeySecret>';
const greenVersion = '2017-01-12';
var hostname = 'green.cn-shanghai.aliyuncs.com';
var path = '/green/image/scan';

var clientInfo = {
	"ip":"127.0.0.1"
};


// 请求体,根据需要调用相应的算法
var requestBody = JSON.stringify({  
    bizType:'Green',
    scenes:['porn'],
    tasks:[{
    	'dataId':uuidV1(),
    	'url':'https://xxx.png'
    }]
}); 

var bizCfg = {
	'accessKeyId' : accessKeyId,
	'accessKeySecret' : accessKeySecret,
	'path' : path,
	'clientInfo' : clientInfo,
	'requestBody' : requestBody,
	'hostname' : hostname,
	'greenVersion' : greenVersion
}


greenNodejs(bizCfg, execute);


// 业务代码，根据不同算法的返回结果采取相应的业务流程
function execute(chunk){
	console.log('BODY: ' + chunk);
}
