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
    tar -xzvf hadoop-$hadoop_version.tar.gz
    mv hadoop-$hadoop_version hadoop

    sudo mv hadoop /opt/hadoop 
}




cp_to_slaves(){

    #authorize key must be cp to the slaves and successfully tested first.
    #number of slaves must be config in the env file.

    for (( c=1; c<=$n_slaves; c++ ))
    do 
        target_host=slave_${c}_host
        target_user=slave_${c}_username

        echo "coping to ${!target_host} ...."

        # cp hadoop to slaves
        scp -r $HADOOP_HOME ${!target_user}@${!target_host}:~/hadoop_installation/

        #cp utils functions to slaves
        scp -r utils.sh ${!target_user}@${!target_host}:~/hadoop_installation/

        #cp config to slaves
        scp -r config ${!target_user}@${!target_host}:~/hadoop_installation/

        #cp slave setup file to target host
        scp -r slave_setup.sh ${!target_user}@${!target_host}:~/hadoop_installation/

        #scp env config
        scp -r env_config.sh ${!target_user}@${!target_host}:~/hadoop_installation/

        scp -r hosts ${!target_user}@${!target_host}:~/hadoop_installation/

        scp -r workers ${!target_user}@${!target_host}:~/hadoop_installation/

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

        echo ${!target_host}
        echo "${!target_ip}     ${!target_host}"
        echo ${!target_host} >> workers
        echo "${!target_ip}     ${!target_host}"  >> hosts


    done
}


generate_host_infor(){


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
                echo "export master_username=${string[2]} " >> $target
                echo "   
                
                ">> $target
                count=$((count + 1))
                
            else
                echo "# slave ${count} info" >> $target
                echo "export slave_${count}_ip=${string[0]} " >> $target
                echo "export slave_${count}_host=${string[1]} " >> $target
                echo "export slave_${count}_username=${string[2]}" >> $target
                echo "  
                
                " >> $target
                count=$((count + 1))
            fi
        fi
    done < "cluster_info.txt"


    count=$((count - 1))
    echo "export n_slaves=$count" >> $target

}



env_setup(){

    sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y

    mv hadoop /opt/hadoop


    JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
    HADOOP_HOME=/opt/hadoop

    cat bashrc.sh  >> ~/.bashrc

    source ~/.bashrc

    echo "export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64" >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
}