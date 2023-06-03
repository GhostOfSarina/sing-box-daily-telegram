#!/bin/bash

apt-get install apache2

cd /var/www/html/
git clone https://github.com/codingstella/vCard-personal-portfolio.git
cp -ar ./vCard-personal-portfolio/  /var/www/html/
rm -rf ./vCard-personal-portfolio/


cp /root/subscribe.txt /var/www/html/subscribe.txt 
