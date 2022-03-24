#!/bin/bash
cd /home/lighthouse/gatelligance
echo building backend...
go build -o gb
echo starting backend...
nohup ./gb &



#!/bin/bash
cd /home/admin/gatelligance/algo
echo building backend...
go build -o gb
echo starting backend...
nohup ./gb &
