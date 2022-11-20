

config = {
    "hadoop_version": "3.3.1",
    "java_version":"8",
    "ubuntu_version":"20.04"
}


class BaseConfig:
    def __init__(self, master, all_domains) -> None:
        
        self.ubuntu_version = config["ubuntu_version"]
        self.java_version = config["java_version"]
        self.hadoop_version = config["hadoop_version"]

        self.DOCKER_FILE = """FROM ubuntu:{0}
                                \nRUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends openjdk-{1}-jdk net-tools curl netcat gnupg libsnappy-dev && rm -rf /var/lib/apt/lists/*
                                \nENV JAVA_HOME=/usr/lib/jvm/java-{1}-openjdk-amd64/
                                \nRUN curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS
                                \nRUN gpg --import KEYS
                                \nENV HADOOP_VERSION {2}
                                \nENV HADOOP_URL https://dlcdn.apache.org/hadoop/common/hadoop-$HADOOP_VERSION/hadoop-$HADOOP_VERSION.tar.gz
                                \nRUN set -x && curl -fSL "$HADOOP_URL" -o /tmp/hadoop.tar.gz && curl -fSL "$HADOOP_URL.asc" -o /tmp/hadoop.tar.gz.asc && gpg --verify /tmp/hadoop.tar.gz.asc && tar -xvf /tmp/hadoop.tar.gz -C /opt/ && rm /tmp/hadoop.tar.gz*
                                \nRUN ln -s /opt/hadoop-$HADOOP_VERSION/etc/hadoop /etc/hadoop
                                \nRUN mkdir /opt/hadoop-$HADOOP_VERSION/logs
                                \nRUN mkdir /hadoop-data
                                \nENV HADOOP_HOME=/opt/hadoop-$HADOOP_VERSION
                                \nENV HADOOP_CONF_DIR=/etc/hadoop
                                \nENV MULTIHOMED_NETWORK=1
                                \nENV USER=root
                                \nENV PATH $HADOOP_HOME/bin/:$PATH
                                \nRUN echo  "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/jre" >> "$HADOOP_HOME/etc/hadoop/hadoop-env.sh"
                                """.format(self.ubuntu_version, self.java_version,
                                                            self.hadoop_version)
        self.HADOOP_HOME = "/opt/hadoop-{0}".format(self.hadoop_version)

        self.core_site_config="""<configuration>\n<property>\n\t<name>fs.default.name</name>\n\t<value>hdfs://{0}:9000</value>\n</property>\n</configuration>""".format(master)
        self.hdfs_site_config="<configuration>\n<property>\n\t<name>dfs.namenode.name.dir</name>\n\t<value>hadoop/data/nameNode</value>\n</property>\n<property>\n\t<name>dfs.datanode.data.dir</name>\n\t<value>hadoop/data/dataNode</value>\n</property>\n<property>\n\t<name>dfs.replication</name>\n\t<value>2</value>\n</property>\n</configuration>"
        self.mapred_site_config=""" <configuration>\n<property>\n\t<name>mapreduce.framework.name</name>\n\t<value>yarn</value>\n</property>\n<property>\n\t<name>yarn.app.mapreduce.am.env</name>\n\t<value>HADOOP_MAPRED_HOME={0}</value>\n</property>\n<property>\n\t<name>mapreduce.map.env</name>\n\t<value>HADOOP_MAPRED_HOME={0}</value>\n</property>\n<property>\n\t<name>mapreduce.reduce.env</name>\n\t<value>HADOOP_MAPRED_HOME={0}</value>\n</property> </configuration>""".format(self.HADOOP_HOME)
        self.yarn_site="""<configuration>\n<property>\n\t<name>yarn.acl.enable</name>\n\t<value>0</value>\n</property>\n     <property>\n    <name>yarn.nodemanager.aux-services</name> <value>mapreduce_shuffle</value>\n</property>\n <property>\n    <name>yarn.nodemanager.aux-services.mapreduce.shuffle.class</name>\n     <value>org.apache.hadoop.mapred.ShuffleHandler</value>\n</property>\n <property>\n\t<name>yarn.resourcemanager.hostname</name>\n\t<value>{0}</value>\n</property>\n <property>\n\t<name>yarn.nodemanager.aux-services</name>\n\t<value>mapreduce_shuffle</value>\n</property> </configuration>""".format(master)
        

        self.compose = {
            "services": {
            },
            "volumes":[]

        }

        
        for node in all_domains:
            
            print(node)
            self.compose["services"][node] = {
                "build": {"dockerfile":"dockerfile/{0}Dockerfile".format(node)},
                "container_name": node,
                "restart": "always",
                "ports": ["9870:9870", "9000:9000"],
                "volumes":["{0}:/hadoop".format(node)]
            }

            self.compose["volumes"].append(node)       
        
