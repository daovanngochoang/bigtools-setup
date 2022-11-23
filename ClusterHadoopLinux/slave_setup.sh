#!/bin/bash
source utils.sh
source env_config.sh

# add env to .bashrc
env_setup

# cp hadoop to /opt/hadoop
mv hadoop /opt/hadoop

# add config to hadoop config 
cp config/* $HADOOP_HOME/etc/hadoop/ 

# add env to .bashrc
env_setup

# add config to hadoop config 
cp config/* $HADOOP_HOME/etc/hadoop/ 

# cp hosts to host file
sudo cat hosts >> /etc/hosts
cp workers $HADOOP_HOME/etc/hadoop/ 

