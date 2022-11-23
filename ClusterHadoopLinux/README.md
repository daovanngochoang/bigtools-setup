This setup aim to simplify the installation of hadoop in a distributed cluster of computers.

## Preparation steps.

1.  Generate the ssh keys in the master machine. make sure the username must be the same in all machine or you can create a new user in slave machine that match with the user name in the master machine in order to establish ssh connection without username.
```bash
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
```

2. Copy to all the slaves.
- Change the *slave_username* with the user name of the target slave.
- Change the *slave_ip* with the ip of the target slave.
```bash
scp ~/.ssh/authorized_keys slave_username@splave_ip:~/.ssh/authorized_keys
```

Example: 
```bash
scp .ssh/authorized_keys slave-1@192.168.1.2:~/.ssh/authorized_keys
```

Then try to connect to the target slave with ssh.
```bash
ssh slave_ip 
or
ssh slave-host
```
this shoudld be successful without password!

## Installation steps

1. Clone the repository.
```bash
git clone https://github.com/daovanngochoang/bigdata-tools-setup.git
```

2. Modify the ***cluster_infor.txt*** file. This file should contain the folowing infor and in the order of master infor first then the slave infor. They must have the same username.
```txt
# master ip should be first
#IP                     hostname                 
172.16.96.93            master      
172.16.96.92            slave-1     
172.16.96.91            slave-2     
172.16.96.94            slave-3     

```

3.  Run master_setup.sh
```bash
./master_setup.sh
```

4. Run slave setup in the slaves
```bash
./slave_setup.sh
```

## Run

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


