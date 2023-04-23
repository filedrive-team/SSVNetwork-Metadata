#!/bin/bash
CURR_DIR=$(
    cd $(dirname $BASH_SOURCE)
    pwd
)

PID_FILES=$(ls $CURR_DIR/_*.pid)
echo "pid files : $PID_FILES"
echo "do kill"
cat $PID_FILES | xargs kill -9 
