#!/bin/bash

pid=`ps -ef | grep "/home/rock/rpi_scada_compiled/collector/scada" | grep -v 'grep' | awk '{print $2}'`
if [ -z $pid ]; then
	echo $(date)
	 sh /home/rock/scada.sh 
	 echo ""
fi
