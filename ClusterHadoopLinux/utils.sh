#!/bin/bash


download_hadoop(){

    export hadoop_version=$1
    export java_version=$2

    # download hadoop and extract
    echo "installing hadoop version: $hadoop_version"

    # download hadoop 
    curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS
    wget https://dlcdn.apache.org/hadoop/common/hadoop-$hadoop_version/hadoop-$hadoop_version.tar.gz

    # extract and rename to make it simple name.
    mv hadoop-$hadoop_version.tar.gz hadoop.tar.gz

    tar -xzvf hadoop.tar.gz
    sudo mv hadoop /opt/hadoop 
}




cp_to_slaves(){

    #authorize key must be cp to the slaves and successfully tested first.
    #number of slaves must be config in the env file.
    target_folder=bigdata-tools-setup/hadoop

    for (( c=1; c<=$n_slaves; c++ ))
    do 
        target_host=slave_${c}_host

        echo "coping to ${!target_host} ...."

        ssh  ${!target_host} "mkdir -p ${target_folder}"

        # cp hadoop to slaves
        scp -r  ${!target_host}:~/${target_folder}
        
        #config
        scp -r config ${!target_host}:~/${target_folder}

        #cp utils functions to slaves
        scp utils.sh ${!target_host}:~/${target_folder}

        #cp slave setup file to target host
        scp slave_setup.sh ${!target_host}:~/${target_folder}

        #scp env config
        scp env_config.sh ${!target_host}:~/${target_folder}

        scp hosts ${!target_host}:~/${target_folder}

        scp workers ${!target_host}:~/${target_folder}

        scp bashrc.sh ${!target_host}:~/${target_folder}

    done
}


generate_hosts_n_workers ()
{
    echo $master_host >> workers
    echo "${master_ip}       ${master_host}" >> hosts

    for (( c=1; c<=$n_slaves; c++ ))
    do 
        target_host=slave_${c}_host
        target_ip=slave_${c}_ip

        echo ${!target_host} >> workers
        echo "${!target_ip}     ${!target_host}"  >> hosts


    done
}


generate_setup_env(){

    count=0
    target="env_config.sh"

    echo "#!/bin/bash" >> $target

    while read -r line
    do
        string=(${line//" "/ })
    
        # not include comments
        if ( [ ! ${string[0]} == "#IP" ] && [ ! ${string[0]:0:1} == "#" ] )
        then

            if [ $count == 0 ]
            then 
                echo "# master info" >> $target
                echo "export master_ip=${string[0]} " >> $target
                echo "export master_host=${string[1]} " >> $target
                echo "   
                
                ">> $target
                count=$((count + 1))
                
            else
                echo "# slave ${count} info" >> $target
                echo "export slave_${count}_ip=${string[0]} " >> $target
                echo "export slave_${count}_host=${string[1]} " >> $target
                echo "  
                
                " >> $target
                count=$((count + 1))
            fi
        fi
    done < "cluster_info.txt"


    count=$((count - 1))
    echo "export n_slaves=$count" >> $target

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