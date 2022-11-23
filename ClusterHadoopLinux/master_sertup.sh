#!/bin/bash

# import functions from utils.sh file
source utils.sh


# # download hadoop
# download_hadoop 3.3.1 8


#generate host file 
generate_host_infor
source env_config.sh


# add env to .bashrc
echo "setup env variables to .bashrc "
env_setup

# add config to hadoop config 
echo "add hadoop xml files config to hadoop etc .... \n "
cp config/* $HADOOP_HOME/etc/hadoop/ 

#generate host file
generate_hosts_n_workers


# cp hosts to host file
echo "config hosts and workers ... \n"
cat hosts | sudo tee -a /etc/hosts  
cp workers $HADOOP_HOME/etc/hadoop/ 


cp_to_slaves

hdfs namenode -format
start-dfs.sh

