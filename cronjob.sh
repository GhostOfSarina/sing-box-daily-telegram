#!/bin/bash

croncmd="/root/sing-box-telegram > /root/cronjob.log 2>&1"
cronjob="0 10 * * * $croncmd"
( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -