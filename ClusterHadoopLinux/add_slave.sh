#!/bin/bash

ip=$1
host_name=$2


# add to /etc/hosts
echo "$ip      $host_name" | sudo tee /etc/hosts

source ~/.bashrc
# add to $HADOOP_HOME/etc/hadoop/workers
echo $host_name >> $HADOOP_HOME/etc/hadoop/workers

# send setup to new slave 
send_setup host_name ""    # target folder "" mean home

# restart all in master.
# stop-all.sh

# start-all.sh