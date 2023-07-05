# sing-box-daily-telegram
sing box with send configuration in the telegram channel every day.


# Persian Articles
[A to Z make a Sing-box VPN for all members of the family](https://telegra.ph/A-to-Z-make-a-Sing-box-VPN-for-all-members-of-the-family-06-01)<br />
[Explain sb-server-configer](https://telegra.ph/Small-family-servers-05-17)<br />
[Explain Sing-box](https://telegra.ph/How-run-Reality-protocol-with-Xray-or-Sing-box-Core-with-iSegaro-04-18)

# Copywriting
This project is fork of [sing-REALITY-box](https://github.com/deathline94/sing-REALITY-Box).<br />
The main Idea is combine [sb-server-configer](https://github.com/hrostami/sb-server-configer) with bash script.
It means that implement outstanding feature [sb-server-configer] with bash script.
iSegaro sing-box configuration [sing-box](https://raw.githubusercontent.com/iSegaro/Sing-Box/main/sing-box_config.json)

# How to use
Clone the Project and run the sing-REALITY-box bash script

```
cd /root
git clone https://github.com/GostOfSarina/sing-box-daily-telegram.git
cp -ar ./sing-box-daily-telegram/* /root/
```

```
sudo chmod +x /root/sing-REALITY-box.sh
bash /root/sing-REALITY-box.sh
```



## Fill these files with your own information.


We have three configuration options. Get bot token and chat id from your telegram account and telegram bot father. <br />

get bot token from [BotFather](https://t.me/BotFather)<br />
get chat id from [Find Channel id](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35)


Fill the configuration inside the files.

```
touch /root/bot_token.txt
echo "2XXXXXXXX1:AXXXX_9XXXXXXXXXXXXXXXN-RXXXXXs" > /root/bot_token.txt

touch /root/chat_id.txt
echo "-10000000000000" > /root/chat_id.txt

```


only store the bot token and chat id in these files. noting more and then you can check the data is insert correctly.<br />


```cat /root/bot_token.txt```

```cat /root/chat_id.txt```




public key is automatically make with sing-Realty-Box script.<br />


I add store public key in the original project folder.
```/root/public_key.txt``` <br />



# Setup the cronjob
Cron job is the time scheduler for run the script automatically. after two days new configuration will send to your channel.


```
bash ./cronjob.sh
```



see the cronjob list
```crontab -l```

result:

```0 22 1-31/3 * * /root/sing-box-telegram > /root/cronjob.log 2>&1```



for edit cronjob use these command
```crontab -e```

[more information for cron job](https://www.youtube.com/watch?v=v952m13p-b4) 


show log of cronjob ``` cat cronjob.log ```

you can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


for example use ```30 9 * * 6``` for the “At 09:30 on Saturday.” 


# Get New Configuration

For sending the new configuration to telegram channel. 

```

wget https://github.com/GostOfSarina/sing-box-daily-telegram/releases/download/v.1.1.2/sing-box-telegram
sudo chmod +x ./sing-box-telegram

./sing-box-telegram
```

Now check the configuration that send inside your channel.

# Fake HTML and subscribe to Sing-box 
This part is optional and it used for fake html and give url link to members of the Telegram channel.


you can share ```http://ip/subscribe.txt``` to members of the Telegram channel.
And Also you can use ```http://ip/subscribe.html``` for fake html.



# Install bach maker for different SNI with Golang 

for build go file
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sing-box-telegram
```



# Install Obfs4proxy plugin (Optional)
If you don't need this server or you don't want renew the VPS, you can install this plugin to help tor project.

```
ch /root
sudo chmod +x /root/obfs4proxy.sh.sh
bash ./obfs4proxy.sh.sh
```
