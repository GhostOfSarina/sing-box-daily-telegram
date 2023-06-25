#!/bin/bash

apt install tor
apt install obfs4proxy

cp ./torrc  /etc/tor/torrc 

systemctl enable --now tor.service

systemctl stop tor.service
systemctl start tor.service
systemctl status tor.service