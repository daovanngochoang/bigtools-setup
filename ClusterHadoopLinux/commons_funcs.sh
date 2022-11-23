#!/bin/bash


export hadoop_version=3.3.1



extract_hadoop(){

    # extract and rename to make it simple name.
    echo "tar -xzvf hadoop-$hadoop_version.tar.gz"
    echo "
    
    "
    tar -xzvf hadoop-$hadoop_version.tar.gz
    mv hadoop-$hadoop_version hadoop
    sudo mv hadoop /opt/hadoop 
}



write_env_variable(){

    echo "install java-8 ..."
    echo "
    
    "
    sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y


    JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
    HADOOP_HOME=/opt/hadoop

    echo "cat bashrc.sh  >> ~/.bashrc"
    echo "
    
    "
    cat bashrc.sh  >> ~/.bashrc

    echo "# set master node"
    echo "# set master node" >> ~/.bashrc
    echo "export MASTER_NODE=${master_host}" >> ~/.bashrc

    source .bashrc

    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh"
    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64" >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
}