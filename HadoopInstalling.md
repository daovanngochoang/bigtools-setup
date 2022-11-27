## Introduction.
1. This is a flexible tutorial that help you install hadoop on both single node and multiple nodes cluster.
2. There are several important file we are going to change that are the xml config file, workers and hosts file, read carefully my explaination before editing the file.

## Hadoop installation on single node cluster.

**Step 1:**  Generate SSH key and test ssh
```bash
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
```

```bash
ssh localhost
```

**Step 2: Install Java 8 and necessary libs**
```bash
sudo apt-get install openjdk-8-jdk \
		net-tools curl netcat \
		gnupg openssh-server \
		libsnappy-dev -y
```

**Step 3: Download hadoop.**
```bash
curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS
wget https://dlcdn.apache.org/hadoop/common/hadoop-3.3.1/hadoop-3.3.1.tar.gz

tar -xzf hadoop-$hadoop_version.tar.gz
mv hadoop-$hadoop_version hadoop
sudo mv hadoop /opt/hadoop
```

**Step 4: Edit .bashrc or .zshrc base on what shell you are using**
```bash
# config java home paths. 
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 
export PATH=$PATH:$JAVA_HOME/bin 

#config for hadoop paths.  
export HADOOP_HOME=/opt/hadoop
export HADOOP_INSTALL=$HADOOP_HOME 
export HADOOP_MAPRED_HOME=$HADOOP_HOME 
export HADOOP_COMMON_HOME=$HADOOP_HOME 
export HADOOP_HDFS_HOME=$HADOOP_HOME 
export YARN_HOME=$HADOOP_HOME 
export HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native 
export PATH=$PATH:$HADOOP_HOME/sbin:$HADOOP_HOME/bin 
export HADOOP_OPTS="-Djava.library.path=$HADOOP_HOME/lib/native"

# the master_node indicate the master node in the hadoop cluster
# in this case we install only 1 cluster on 1 machine then 
# the master node is the localhost.
export MASTER_NODE="localhost"

```

**Step 5:** **Edit Config file**
1. **core-site.xml**
```xml
<configuration>
	<property>
		<name>fs.default.name</name>
		<value>hdfs://${MASTER_NODE}:9000</value>
	</property>
</configuration>
```

2. **hdfs-site.xml**
```xml
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
```

3. **mapred-site.xml**
```xml
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
```
4. **yarn-site.xml**
```xml
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
```

5. add to ***$HADOOP_HOME/etc/hadoop/hadoop-env.sh***
```bash
echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 " >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
```


## Hadoop installation on multiple nodes cluster.

### Note
1. The master machine should be the machine that you are currently working on the installation process.
2. Both master machine and slave machines should have the same user, for example on master we have the user name hadoop then all the slaves should have the hadoop user to make things work. it's the the easiest way to install hadoop.

**Step 1: edit */etc/hosts* file.** *(Both on the master and slaves)*
1. open file with text editor as super user.
```bash
sudo nano /etc/hosts
```

2. add the ip and host name.
	For example: I have 1 master and 3 slave with different public ip.
```text
# PUBLIC IP             HOSTNAME

172.16.96.93            master 
172.16.96.92            slave-1     
172.16.96.91            slave-2     
172.16.96.94            slave-3
```


**Step 2:**  Generate SSH key and copy to slaves.
1. Generate ssh key.
```bash
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
```

2. Copy the key to the slaves machine.
```bash
scp ~/.ssh/authorized_keys slave_host_name:~/.ssh/authorized_keys
```

*example:* *if i have 1 master and 3 slaves*
```bash
scp ~/.ssh/authorized_keys slave-1:~/.ssh/authorized_keys
scp ~/.ssh/authorized_keys slave-2:~/.ssh/authorized_keys
scp ~/.ssh/authorized_keys slave-3:~/.ssh/authorized_keys
```

3. Test SSH and all shoud be work without 
```bash
ssh master
ssh slave-1
ssh slave-2
ssh slave-3
```


**Step 3: Install Java 8 and necessary libs** *(Both on the master and slaves)*
```bash
sudo apt-get install openjdk-8-jdk \
		net-tools curl netcat \
		gnupg openssh-server \
		libsnappy-dev -y
```

**Step 4: Download hadoop.** *(Both on the master and slaves)*
```bash
curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS
wget https://dlcdn.apache.org/hadoop/common/hadoop-3.3.1/hadoop-3.3.1.tar.gz

tar -xzf hadoop-$hadoop_version.tar.gz
mv hadoop-$hadoop_version hadoop
sudo mv hadoop /opt/hadoop
```

**Step 5: Edit .bashrc or .zshrc base on what shell you are using** *(Both on the master and slaves)*
```bash
# config java home paths. 
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 
export PATH=$PATH:$JAVA_HOME/bin 

#config for hadoop paths.  
export HADOOP_HOME=/opt/hadoop
export HADOOP_INSTALL=$HADOOP_HOME 
export HADOOP_MAPRED_HOME=$HADOOP_HOME 
export HADOOP_COMMON_HOME=$HADOOP_HOME 
export HADOOP_HDFS_HOME=$HADOOP_HOME 
export YARN_HOME=$HADOOP_HOME 
export HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native 
export PATH=$PATH:$HADOOP_HOME/sbin:$HADOOP_HOME/bin 
export HADOOP_OPTS="-Djava.library.path=$HADOOP_HOME/lib/native"

# the master_node indicate the master node in the hadoop cluster
# this variable should match with the hostname of the master node
# in my case its name is master. 
export MASTER_NODE="master"

```

**Step 6:** **Edit Config file** *(Both on the master and slaves)*

1. Edit **$HADOOP_HOME/etc/hadoop/core-site.xml** to decide what is the master node.
```xml
<configuration>
	<property>
		<name>fs.default.name</name>
		<value>hdfs://${MASTER_NODE}:9000</value>
	</property>
</configuration>
```

2. Edit **$HADOOP_HOME/etc/hadoop/hdfs-site.xml**
```xml
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
```

3.  Edit **$HADOOP_HOME/etc/hadoop/mapred-site.xml**
```xml
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
```

4. Edit **$HADOOP_HOME/etc/hadoop/yarn-site.xml** in the yanr-site, which node will manager all resources. We will use master node to manage all the other nodes.
```xml
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
```

5. Add to ***$HADOOP_HOME/etc/hadoop/hadoop-env.sh***
```bash
echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 " >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
```

**Step 7:** **add slaves to workers file, you need to add all slave hostname to this file.**
1. Open with editor.
```bash
nano $HADOOP_HOME/etc/hadoop/workers
```

2. Add the hostname to this file, in my case i have 1 master and 3 slaves.
```text
master
slave-1
slave-2
slave-3
```


## Run Hadoop.
1. run only on master node.
```bash
hdfs namenode -format
start-dfs.sh
```

2. run ***jps*** on master node and slave we should see.
```txt
Jps
NameNode
SecondaryNameNode
```

In the slaves should have:
```txt
Jps
Datanode
```

3. You can run http://node-master-IP:9870 this to see the web app, remember to replace the node-master-ip by the real ip.

## MapReduce test.

1. Create a _books_ directory
```bash
hdfs dfs -mkdir -p /user/hadoop
hdfs dfs -mkdir books
```

2.  Download a few txt file to test.
```bash
cd /home/hadoop 
wget -O alice.txt https://www.gutenberg.org/files/11/11-0.txt 
wget -O holmes.txt https://www.gutenberg.org/files/1661/1661-0.txt 
wget -O frankenstein.txt https://www.gutenberg.org/files/84/84-0.txt

```

3. Put the hadoop file system
```bash
hdfs dfs -put alice.txt holmes.txt frankenstein.txt books
```

4. HDFS is a distributed storage system, and doesn’t provide any services for running and scheduling tasks in the cluster. This is the role of the YARN framework. The following section is about starting, monitoring, and submitting jobs to YARN.
```bash
start-yarn.sh
```

5. get the node list.
```bash
yarn node -list
```

6. Submit a job with the sample `jar` to YARN. On **node-master** .
```bash
yarn jar $HADOOP_HOME/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.1.2.jar wordcount "books/*" output
```

7.  To see the result of hadoop after mapreduce, run.
```bash
hdfs dfs -cat output/part-r-00000 | less
```


