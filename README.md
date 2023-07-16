# Sing Box Reality Daily Telegram
Sing box Reality with send configuration in the telegram channel every day.This project send sing box Reality configuration to your channel base on schedule.<br />
Also you can donate your configuration to Yebekhe Systems.

# Supported Version
Ubuntu 22.04

# Persian Articles

[Youtube video was made by AliAlma_GSM](https://www.youtube.com/watch?v=shZnK6Kn4qk)<br />
[Easy Installation](https://telegra.ph/process-of-install-sing-box-daily-telegram-07-08)<br />
[A to Z make a Sing-box VPN for all members of the family](https://telegra.ph/A-to-Z-make-a-Sing-box-VPN-for-all-members-of-the-family-06-01)<br />
[Explain sb-server-configer](https://telegra.ph/Small-family-servers-05-17)<br />
[Explain Sing-box](https://telegra.ph/How-run-Reality-protocol-with-Xray-or-Sing-box-Core-with-iSegaro-04-18)<br />

# Copywriting
This project is fork of [sing-REALITY-box](https://github.com/deathline94/sing-REALITY-Box).<br />
The main Idea is combine [sb-server-configer](https://github.com/hrostami/sb-server-configer) with bash script.<br />
It means that implement outstanding feature ```sb-server-configer``` with bash script. These feature added into the [reality-ezpz](https://github.com/aleskxyz/reality-ezpz).<br />

iSegaro sing-box Reality configuration [sing-box](https://raw.githubusercontent.com/iSegaro/Sing-Box/main/sing-box_config.json)<br />
iSegaro sing-box GRPC Reality configuration [sing-box](https://github.com/iSegaro/Sing-Box/blob/main/sing-box_config_GRPC.json)<br />
My sister-in-law project is [yebekhe](https://github.com/yebekhe/TelegramV2rayCollector). This project gathering configuration from the telegram channel<br />



# Fill bot token and chanel id files with your own information.


We have two configuration options. Get bot token and chat id from your telegram account and telegram bot father. <br />

get bot token from [BotFather](https://t.me/BotFather)<br />
get chat id from [Find Channel id](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35)<br />


Fill the configuration inside the setting.json.

```
bot_token =>  "2XXXXXXXX1:AXXXX_9XXXXXXXXXXXXXXXN-RXXXXXs"
chat_id =>  "-10000000000000" 
```

# Fill setting file with your values

Setting file is located in /root/settings.json and you can easily modify settings. After changing settings, it necessary to run again `./sing-box-telegram` after changing.  <b />

Edit this setting file base on your needs.<br />

```ports``` are the ports that you want to use in your server. <br />
```domains``` are the domains that you want to use in your SNI. <br />
```bot_token``` is the bot token that you get from the bot father. <br />
```chat_id``` is the chat id that you get from the channel. <br />
```donate_url``` is the url that you want to send your configuration. <br />
```dynamic_subscription``` is the boolean value that you want to have dynamic subscribe link like ```subscribe.txt``` or dynamic ones like ```subscribe.122.txt``` <br />
```channel_name``` is the channel name that you want to send your configuration. <br />
```send_vnstat``` is the boolean value that you want to send ( Bandwidth usage ) vnstat information to the channel. <br />


```
cd /root
touch /root/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2053, 2082, 8443, 6443, 2096 ],
    \"domains\": [
        \"www.datadoghq.com\",
        \"000webhost.ir\",
        \"speedtest.net\",
        \"speed.cloudflare.com\",
        \"fruitfulcode.com\",
        \"favakar.ir\",
        \"benecke.com\",
        \"tarhpro.ir\",
        \"fernandotrueba.com\",
        \"mathhub.info\"
    ],
    \"bot_token\" : \"627434621931:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1002343276432\",
    \"donate_url\" : \"https://where_ever_you_want.site",
    \"dynamic_subscription\" : true,
    \"channel_name\" : \"Sarina_Esmailzadeh\",
    \"send_vnstat\" : true
}">  /root/setting.json
```


You can also edit your file with nano editor. [youtube nano](https://www.youtube.com/watch?v=fJTPjWuyrIY) [youtube nano 2](https://www.youtube.com/watch?v=NV9PyPJKqH4) <br />

# Donate your server to the Yebekhe

You can also send your automatically to the  yebekhe server or what URL you wants. just fill `donate_url` with your desirable address.<br />
My sister-in-law project is [Yebekhe](https://github.com/yebekhe/TelegramV2rayCollector)<br />

Don't need to have telegram channel. <br />
Don't need to have telegram bot.<br />


```
cd /root
touch /root/setting.json
echo "{
    \"ports\": [443, 22, 10050],
    \"domains\": [
        \"www.datadoghq.com\",
        \"000webhost.ir\",
        \"speedtest.net\"
    ],
    \"bot_token\" : \"\",
    \"chat_id\" : \"\",
    \"donate_url\" : \"https://yebekhe.000webhostapp.com/donate/\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"Sarina\",
    \"send_vnstat\" : false
}">  /root/setting.json
```

After install that explain in below section. you can change scheduler time in the cronjob.
```
wget https://raw.githubusercontent.com/GhostOfSarina/sing-box-daily-telegram/main/first-time-install-sing-box.sh

```

for edit cronjob use these command:
```crontab -e```


put every hour schedule for update the configuration. 
```0 * * * *```



# How to install
For fast way install and run this service you need download below files and execute them. 
For security reason, I recommend you to change ssh port. change 9001 to any port that you want.

```
echo "Port 9001" >> /etc/ssh/sshd_config
systemctl restart sshd
service ssh restart
```
after you need ``` -p 9001 ``` for ssh connection.for example ```ssh root@ip -p 9001``` <br />


<b> 9001 is the default port for SSH connection. don't use this port in setting file. <b/>

If you had a below error please restart your server. <b />
```kex_exchange_identification: read: Connection reset by peer
Connection reset by x.x.x.x port 22
lost connection
```
or restart your service ```service ssh status``` and ```service ssh restart```




Download bash files and add permission for execute.
```
cd /root

wget https://raw.githubusercontent.com/GhostOfSarina/sing-box-daily-telegram/main/first-time-install-sing-box.sh

sudo chmod +x /root/first-time-install-sing-box.sh

bash /root/first-time-install-sing-box.sh
```


# For sending the new configuration to telegram channel. 

Check send the new configuration to telegram channel.

```
./sing-box-telegram
```

after command execution the configuration send to your telegram channel.


##  << Other options ( professional edit project) >>


# Modify the cronjob ( change time that want to send to the channel)
Cron job is the time scheduler for run the script automatically. after two days new configuration will send to your channel.


see the cronjob list
```crontab -l```

result:

```0 9 1-31/3 * * /root/sing-box-telegram > /root/cronjob.log 2>&1```



for edit cronjob use these command
```crontab -e```

[more information for cron job](https://www.youtube.com/watch?v=v952m13p-b4) 


show log of cronjob ``` cat cronjob.log ```

you can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


for example use ```30 9 * * 6``` for the “At 09:30 on Saturday.” 



# Fake HTML and subscribe to Sing-box 
This part for fake html and give url link to members of the Telegram channel.


you can share ```http://ip/subscribe.txt``` to members of the Telegram channel.
And Also you can use ```http://ip ``` for fake html.



# Install bach maker for different SNI with Golang 

for build go file use below command line:
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sing-box-telegram
```


# Install Obfs4proxy plugin (Use only your ip address is blocked by Iranian or chinese GFW)
If you don't need this server or you don't want renew the VPS, you can install this plugin to help tor project.

```
ch /root
sudo chmod +x /root/obfs4proxy.sh
bash ./obfs4proxy.sh
```
