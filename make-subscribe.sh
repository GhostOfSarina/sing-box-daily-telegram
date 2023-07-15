#!/bin/bash

rm -rf /var/www/html/subscribe.*
cp  /root/subscribe.* /var/www/html/

# Restart sing-box service
systemctl restart sing-box