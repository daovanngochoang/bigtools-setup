#!/bin/bash

# import functions from commons_funcs.sh file
source commons_funcs.sh



generate_hosts_n_workers (){
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



cp_to_slaves(){

    #authorize key must be cp to the slaves and successfully tested first.
    #number of slaves must be config in the env file.
    target_folder=""

    for (( c=1; c<=$n_slaves; c++ ))
    do 
        target_host=slave_${c}_host

        echo "coping to ${!target_host} .... " \n\n\n

        # send setup.
        send_setup ${!target_host} $target_folder

    done
}



download_hadoop(){

    # download hadoop and extract
    echo "installing hadoop version $hadoop_version ..." \n\n\n

    # download hadoop 
    curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS
    wget https://dlcdn.apache.org/hadoop/common/hadoop-$hadoop_version/hadoop-$hadoop_version.tar.gz

}





#----------------RUN-----------------------------#

#download and extract hadoop
download_hadoop 
extract_hadoop  #from utils


#generate host file 
generate_setup_env
source env_config.sh

# add env to .bashrc
echo "setup env variables to .bashrc "\n\n\n
write_env_variable

# write config
./write_config.sh

# add config to hadoop config 
echo "add hadoop xml files config to hadoop etc .... "\n\n\n
cp config/* $HADOOP_HOME/etc/hadoop/ 

#generate host file
echo "generate host file" \n\n\n
generate_hosts_n_workers


# cp hosts to host file
echo "config hosts and workers ... "\n\n
cat hosts | sudo tee -a /etc/hosts  
cp workers $HADOOP_HOME/etc/hadoop/ 


cp_to_slaves

source ~/.bashrc




