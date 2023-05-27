#!/bin/bash

# Check if jq is installed, and install it if not
if ! command -v jq &> /dev/null; then
    echo "jq is not installed. Installing..."
    if [ -n "$(command -v apt)" ]; then
        apt update
        apt install -y jq
    elif [ -n "$(command -v yum)" ]; then
        yum install -y epel-release
        yum install -y jq
    elif [ -n "$(command -v dnf)" ]; then
        dnf install -y jq
    else
        echo "Cannot install jq. Please install jq manually and rerun the script."
        exit 1
    fi
fi

# Check if reality.json, sing-box, and sing-box.service already exist
if [ -f "/root/reality.json" ] && [ -f "/root/sing-box" ] && [ -f "/etc/systemd/system/sing-box.service" ]; then


    echo "Renew..."

    # Generate new uuid
    uuid=$(/root/sing-box generate uuid)
    # Generate necessary values
    short_id=$(jq -r '.inbounds[0].tls.reality.short_id[0]' /root/reality.json)


    # Modify reality.json with new settings
    jq --arg uuid "$uuid"  '.inbounds[0].users[0].uuid = $uuid' /root/reality.json > /root/reality_modified.json
    mv /root/reality_modified.json /root/reality.json


    listen_port=$(jq -r '.inbounds[0].listen_port' /root/reality.json)
    server_name=$(jq -r '.inbounds[0].tls.server_name' /root/reality.json)
    server_ip=$(curl -s https://api.ipify.org)


    public_key=`cat /root/public_key.txt`

    current_date_time=$(date +"%A-%T")

    server_link="vless://$uuid@$server_ip:$listen_port?encryption=none&flow=xtls-rprx-vision&security=reality&sni=$server_name&fp=chrome&pbk=$public_key&sid=$short_id&type=tcp&headerType=none#SING-BOX-$date"

    server_link_base64=$(printf $server_link | base64 -w0);


    echo "Send server link to telegram"


    BOT_TOKEN=`cat /root/bot_token.txt`
    CHAT_ID=`cat /root/chat_id.txt` 

    telegram_url="https://api.telegram.org/bot$BOT_TOKEN/sendMessage?chat_id=$CHAT_ID&text=$server_link_base64"

    curl $telegram_url



    # Restart sing-box service
    systemctl restart sing-box
    echo "DONE!"
    exit 0
    
    
fi


 exit 1