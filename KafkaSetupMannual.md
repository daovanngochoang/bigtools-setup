# Kafka Installation

### Step 1: Install java8 scala-2.13
```bash
sudo apt-get install openjdk-8-jdk \
		net-tools curl netcat \
		gnupg openssh-server \
		libsnappy-dev -y
```

### Step 2: Get Spark from Apache
```bash
wget https://downloads.apache.org/kafka/3.3.1/kafka_2.13-3.3.1.tgz
tar -xzf kafka_2.13-3.3.1.tgz
mv kafka_2.13-3.3.1 kafka
sudo mv kafka /opt/kafka
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

# Kafka env.
export KAFKA_HOME=/opt/kafka
export PATH=$PATH:$KAFKA_HOME/bin
```


## Test Kafka.
1. Start zookeeper and kafka server
```bash
zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties
kafka-server-start.sh $KAFKA_HOME/config/server.properties

```

2. create topic
```bash
kafka-topics.sh --bootstrap-server localhost:9092 --partitions 2 --replication-factor 1 --topic test_topic --create
```
3. List topic
```bash
kafka-topics.sh --list --bootstrap-server localhost:9092
```
4. describe topic
```bash
kafka-topics.sh --bootstrap-server localhost:9092 --topic test_topic --describe
```
5. Delete topic
```bash
kafka-topics.sh --bootstrap-server localhost:9092 --topic test_topic --delete
```
6. Send message.
```bash
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test_topic 
> hello bro!
```
7. Receive message.
```bash 
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test_topic
```
8. Receive message by order.
```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test_topic --from-beginning

```

