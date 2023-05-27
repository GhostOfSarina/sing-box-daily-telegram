# sing-box-daily-telegram
sing box with send configuration in the telegram channel every day.


# Persian Articles
[Explain sb-server-configer](https://telegra.ph/Small-family-servers-05-17)<br />
[Explain Sing-box](https://telegra.ph/How-run-Reality-protocol-with-Xray-or-Sing-box-Core-with-iSegaro-04-18)

# Copywriting
This project is fork of [sing-REALITY-box](https://github.com/deathline94/sing-REALITY-Box).<br />
The main Idea is combine [sb-server-configer](https://github.com/hrostami/sb-server-configer) with bash script.
It means that implement outstanding feature [sb-server-configer] with bash script.

# How to use
Clone the Project and run the sing-REALITY-box bash script

```
git clone https://github.com/GostOfSarina/sing-box-daily-telegram.git

cp -ar ./sing-box-daily-telegram /root/
```

```
sudo chmod +x /root/sing-REALITY-box.sh
sudo chmod +x /root/renew.sh

bash /root/sing-REALITY-box
```



## Fill these files with your own information.


We have three configuration options.

```/root/public_key.txt``` <br />
I add store public key in the original project folder.


```
touch /root/bot_token.txt
echo "2XXXXXXXX1:AXXXX_9XXXXXXXXXXXXXXXN-RXXXXXs" > /root/bot_token.txt

touch /root/chat_id.txt
echo "-10000000000000" > /root/chat_id.txt

```

```/root/bot_token.txt```

```/root/chat_id.txt```


get bot token from [BotFather](https://t.me/BotFather)<br />
get chat id from [Find Channel id](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35)

public key is automatically make with sing-Realty-Box script.

# Setup the cronjob
```
bash ./cronjob.sh
```



see the cronjob list
```crontab -l```

result:

```0 10 * * * /root/renew.sh > /root/cronjob.log 2>&1```



for edit cronjob use these command
```crontab -e```

[more information for cron job](https://www.youtube.com/watch?v=v952m13p-b4) 


show log of cronjob ``` cat cronjob.log ```

you can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


for example use ```30 9 * * 6``` for the “At 09:30 on Saturday.” 


# Get New Configuration

for sending the new configuration to telegram channel

```
bash ./renew.sh
```
