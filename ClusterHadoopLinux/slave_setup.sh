#!/bin/bash
source utils.sh
source env_config.sh

# add env to .bashrc
write_env_variable

source ~/.bashrc

# cp hadoop to /opt/hadoop
sudo mv hadoop /opt/hadoop

# add config to hadoop config 
cp config/* $HADOOP_HOME/etc/hadoop/ 

# cp hosts to host file
cat hosts | sudo tee -a /etc/hosts  

