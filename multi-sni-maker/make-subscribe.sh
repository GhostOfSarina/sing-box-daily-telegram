#!/bin/bash


cp  /root/subscribe.txt /var/www/html/subscribe.txt

# Restart sing-box service
systemctl restart sing-box