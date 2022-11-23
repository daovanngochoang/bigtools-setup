#!/bin/bash


export hadoop_version=3.3.1



extract_hadoop(){

    # extract and rename to make it simple name.
    tar -xzvf hadoop-$hadoop_version.tar.gz
    mv hadoop-$hadoop_version hadoop
    sudo mv hadoop /opt/hadoop 
}



write_env_variable(){

    sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y

    mv hadoop /opt/hadoop


    JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
    HADOOP_HOME=/opt/hadoop

    cat bashrc.sh  >> ~/.bashrc
    echo "# set master node" >> ~/.bashrc
    echo "export MASTER_NODE=${master_host}" >> ~/.bashrc

    
    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64" >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
}