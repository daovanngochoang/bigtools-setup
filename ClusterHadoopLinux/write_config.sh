#!/bin/bash
source utils.sh
source env_config.sh


core_site="""<configuration>
    <property>
        <name>fs.default.name</name>
        <value>hdfs://${master_host}:9000</value>
    </property>
</configuration>"""


hdfs_site="""
<configuration>
    <property>
        <name>dfs.namenode.name.dir</name>
        <value>file://${HOME}/hadoop/data/nameNode</value>
    </property>
    <property>
        <name>dfs.datanode.data.dir</name>
        <value>file://${HOME}/hadoop/data/dataNode</value>
    </property>
    <property>
        <name>dfs.replication</name>
        <value>2</value>
    </property>
</configuration>
"""

mapred_site="""
<configuration>
	<property>
		<name>mapreduce.framework.name</name>
		<value>yarn</value>
	</property>
	<property>
		<name>yarn.app.mapreduce.am.env</name>
		<value>HADOOP_MAPRED_HOME=${HADOOP_HOME}</value>
	</property>
	<property>
		<name>mapreduce.map.env</name>
		<value>HADOOP_MAPRED_HOME=${HADOOP_HOME}</value>
	</property>
	<property>
		<name>mapreduce.reduce.env</name>
		<value>HADOOP_MAPRED_HOME=${HADOOP_HOME}</value>
	</property>
</configuration>
"""

yarn_site="""
<configuration>
    <property>
            <name>yarn.acl.enable</name>
            <value>0</value>
    </property>

    <property>
            <name>yarn.resourcemanager.hostname</name>
            <value>${MASTER_NODE}</value>
    </property>

    <property>
            <name>yarn.nodemanager.aux-services</name>
            <value>mapreduce_shuffle</value>
    </property>

	<property>
		<name>yarn.nodemanager.aux-services.mapreduce.shuffle.class</name>
		<value>org.apache.hadoop.mapred.ShuffleHandler</value>
	</property>
</configuration>
"""


target=config
mkdir $target

echo $core_site >> $target/core-site.xml
echo $hdfs_site >> $target/hdfs-site.xml
echo $mapred_site >> $target/mapred-site.xml
echo $yarn_site >> $target/yarn-site.xml

