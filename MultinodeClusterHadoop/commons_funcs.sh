#!/bin/bash


export hadoop_version=3.3.1



extract_hadoop(){

    # extract and rename to make it simple name.
    echo "tar -xzvf hadoop-$hadoop_version.tar.gz"
    echo "
    
    "
    tar -xzf hadoop-$hadoop_version.tar.gz
    mv hadoop-$hadoop_version hadoop
    sudo mv hadoop /opt/hadoop 
}



write_env_variable(){

    echo "install java-8 ..."
    echo "
    
    "
    sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y


    export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
    export HADOOP_HOME=/opt/hadoop

    echo "cat bashrc.sh  >> ~/.bashrc"
    echo "
    
    "
    cat bashrc.sh  >> ~/.bashrc

    echo "# set master node"
    echo "# set master node" >> ~/.bashrc
    echo "export MASTER_NODE=${master_host}" >> ~/.bashrc

    source ~/.bashrc

    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64 >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh"
    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64" >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
}



send_setup(){

    target_host=$1
    target_folder=$2

    echo "coping to ${target_host} .... " \n\n\n

    ssh  ${target_host} "mkdir -p ${target_folder}"

    # cp hadoop to slaves
    scp -r hadoop-$hadoop_version.tar.gz ${target_host}:~/${target_folder}
    
    #config
    scp -r config ${target_host}:~/${target_folder}

    #cp utils functions to slaves
    scp commons_funcs.sh ${target_host}:~/${target_folder}

    #cp slave setup file to target host
    scp slave_setup.sh ${target_host}:~/${target_folder}

    #scp env config
    scp env_config.sh ${target_host}:~/${target_folder}

    scp hosts ${target_host}:~/${target_folder}

    scp bashrc.sh ${target_host}:~/${target_folder}
}