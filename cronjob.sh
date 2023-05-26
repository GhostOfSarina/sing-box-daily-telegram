#!/bin/bash

croncmd="/root/renew.sh > /root/cronjob.log 2>&1"
cronjob="0 10 * * * $croncmd"
( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -