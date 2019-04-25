using System;
using System.Runtime.Serialization;

namespace Model{

    [DataContract]
    class ClientInfo {
		[DataMember(Order = 0, IsRequired = false, EmitDefaultValue = false)]
		public string sdkVersion { set; get; }

        [DataMember(Order = 1, IsRequired = false, EmitDefaultValue = false)]
        public string cfgVersion { set; get; }

        [DataMember(Order = 2, IsRequired = false, EmitDefaultValue = false)]
        public string userType { set; get; }

		[DataMember(Order = 3, IsRequired = false, EmitDefaultValue = false)]
		public string userId { set; get; }

		[DataMember(Order = 4, IsRequired = false, EmitDefaultValue = false)]
		public string userNick { set; get; }

		[DataMember(Order = 5, IsRequired = false, EmitDefaultValue = false)]
		public string avatar { set; get; }

		[DataMember(Order = 6, IsRequired = false, EmitDefaultValue = false)]
		public string imei { set; get; }

		[DataMember(Order = 7, IsRequired = false, EmitDefaultValue = false)]
		public string imsi { set; get; }

		[DataMember(Order = 8, IsRequired = false, EmitDefaultValue = false)]
		public string umid { set; get; }

		[DataMember(Order = 9, IsRequired = false, EmitDefaultValue = false)]
		public string ip { set; get; }

		[DataMember(Order = 10, IsRequired = false, EmitDefaultValue = false)]
		public string os { set; get; }

		[DataMember(Order = 11, IsRequired = false, EmitDefaultValue = false)]
		public string channel { set; get; }

		[DataMember(Order = 12, IsRequired = false, EmitDefaultValue = false)]
		public string hostAppName { set; get; }

		[DataMember(Order = 13, IsRequired = false, EmitDefaultValue = false)]
		public string hostPackage { set; get; }

		[DataMember(Order = 14, IsRequired = false, EmitDefaultValue = false)]
		public string hostVersion { set; get; }
    }

    [DataContract]
    class BizData{
        [DataMember(Order = 0, IsRequired = false, EmitDefaultValue = false)]
        public string bizType { set; get; }

        [DataMember(Order = 1, IsRequired = true)]
        public string[] scenes { set; get; }

        [DataMember(Order = 2, IsRequired = true)]
        public Task[] tasks { set; get; }

        public BizData(){}

        public BizData(string bizType, string[] scenes, Task[] tasks){
            this.bizType = bizType;
            this.scenes = scenes;
            this.tasks = tasks;
        }

    }

    [KnownType("KnownTypes")]
    [DataContract]
    class Task{
        [DataMember(Order = 0, IsRequired = false, EmitDefaultValue = false)]
        public ClientInfo clientInfo { set; get; }

        [DataMember(Order = 1, IsRequired = false, EmitDefaultValue = false)]
        public string dataId { set; get; }

        [DataMember(Order = 2, IsRequired = false, EmitDefaultValue = false)]
        public long time { set; get; }

        public Task(){}

        public Task(ClientInfo clientInfo, string dataId, long time){
            this.clientInfo = clientInfo;
            this.dataId = dataId;
            this.time = time;
        }

		static Type[] KnownTypes()
		{
			return new Type[] { typeof(ImageTask), typeof(TextTask)};
		}
    }

    [DataContract]
    class ImageTask : Task{
        [DataMember(Order = 0, IsRequired = true)]
        public string url { set; get; }

        public ImageTask(string url){
            this.url = url;
        }
    }

    [DataContract]
    class TextTask : Task{
        [DataMember(Order = 0, IsRequired = true)]
        public string content { set; get; }

        public TextTask(string content){
            this.content = content;
        }
    }
}