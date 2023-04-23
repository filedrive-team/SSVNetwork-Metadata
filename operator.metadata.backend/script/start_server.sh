#!/bin/bash

CURR_DIR=$(
    cd $(dirname $BASH_SOURCE)
    pwd
)
BASE_NAME=$(basename $BASH_SOURCE)
PIDFILE="${CURR_DIR}/_$BASE_NAME.pid"
function check_pid() {
    if [ -s ${PIDFILE} ]; then
        SPID=$(cat ${PIDFILE})
        PID_EXIST=$(ps aux | awk '{print $2}' | grep -w $SPID)
        if [ $PID_EXIST ]; then
            echo "script [$BASE_NAME] is running, do nothing and exit."
            exit 1
        fi
        cat /dev/null >${PIDFILE}
    fi
}
check_pid

# add what you to do --------begin
EXEC_BIN=$CURR_DIR/ssv_operator_server
LOG_FILE=$CURR_DIR/server.log
CONF=$CURR_DIR/app.toml

nohup $EXEC_BIN start --config=$CONF 2>&1 >>$LOG_FILE &

# update pid file
echo $! >${PIDFILE}
