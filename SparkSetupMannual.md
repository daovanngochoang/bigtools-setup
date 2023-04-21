# Spark installation.

## Spark installation
If you just want to run Spark on a single machine then this installation is enough!

<br />

### Step 1: Get Spark from apache.
1. Download spark
```bash
curl https://dlcdn.apache.org/spark/spark-3.3.2/spark-3.3.2-bin-hadoop3.tgz -o spark-3.3.2-bin-hadoop3.tgz
```
2. decompress Spark.

```bash
tar -xzf spark-3.3.2-bin-hadoop3.tgz
mv spark-3.3.2-bin-hadoop3 spark
sudo mv spark ~/bigdata/spark
```
### Step 2: Install java8 and scala 2.13
1. Install java 8.
```bash
sudo apt-get install openjdk-8-jdk \
		net-tools curl netcat \
		gnupg openssh-server \
		libsnappy-dev -y
```

2. Download scala version 2.13
```bash
 wget https://downloads.lightbend.com/scala/2.13.10/scala-2.13.10.deb
```

3. Install with apt or nala or dpkg
```bash
sudo apt install ./scala-2.13.10.deb
or
sudo nala install ./scala-2.13.10.deb
or 
sudo dpkg -i ./scala-2.13.10.deb

```
### Add ENV variables.
1. open .zshrc or .bashrc file depended on what shell you are using.
```bash
nano ~/.barhrc 
or 
nano ~/.zshrc
```
2. add this text to that file.
```bash
# config java home paths. 
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 
export PATH=$PATH:$JAVA_HOME/bin 

# Spark env.
export SPARK_HOME=~/bigtada/spark
export PATH=$PATH:$SPARK_HOME/bin:$SPARK_HOME/sbin
```

# Multi-nodes cluster Configuration

### Note:
If you want to install hadoop on multi-nodes cluster you need to have hadoop installation on all machines (master and slaves).

<br />
<br />

### Step 1: edit */etc/hosts* file. *(Both on the master and slaves)*
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


### Step 2:  Generate SSH key and copy to slaves.
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

### Step 3: Configure Master information *(both on master and slaves)*
```bash
cd $SPARK_HOME/conf
cp spark-env.sh.template spark-env.sh
nano spark-env.sh

# Add below lines
export SPARK_MASTER_HOST=master
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 
```

### Step 4: Add Workers
```bash
cd $SPARK_HOME/conf
cp workers.template workers
nano workers
```
I have 1 master and 3 slave with different public ip.
```bash
# Add bellow lines
master
slave-1
slave-2
slave-3
```



