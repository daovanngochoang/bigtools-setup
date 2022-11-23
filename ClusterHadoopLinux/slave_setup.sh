#!/bin/bash
source commons_funcs.sh
source env_config.sh

# add env to .bashrc
write_env_variable

source ~/.bashrc

extract_hadoop

# add config to hadoop config 
cp config/* $HADOOP_HOME/etc/hadoop/ 

# cp hosts to host file
cat hosts | sudo tee -a /etc/hosts  

