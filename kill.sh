#!/bin/bash
pid=`ps aux|grep gb|grep -v "grep"|awk '{print $2}'`
echo ${pid}
echo killing backend ${pid} ..
kill -9 ${pid}