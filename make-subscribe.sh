#!/bin/bash

rm -rf /var/www/html/subscribe.*
cp  /root/subscribe.* /var/www/html/


rm -rf /var/www/html/aggregate.*
cp  /root/aggregate.* /var/www/html/

# Restart sing-box service
systemctl restart sing-box