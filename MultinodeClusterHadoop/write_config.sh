#!/bin/bash
source env_config.sh


core_site="""<configuration>
    <property>
        <name>fs.default.name</name>
        <value>hdfs://${master_host}:9000</value>
    </property>
</configuration>"""


hdfs_site="""

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
            <value>${master_host}</value>
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

