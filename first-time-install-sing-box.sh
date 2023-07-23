#!/bin/bash



#easy install

cd /root
rm -rf /root/reinstall-sing-box.sh*
rm -rf /root/make-subscribe.sh*

wget https://raw.githubusercontent.com/GhostOfSarina/sing-box-daily-telegram/main/reinstall-sing-box.sh
wget https://raw.githubusercontent.com/GhostOfSarina/sing-box-daily-telegram/main/make-subscribe.sh

sudo chmod +x /root/reinstall-sing-box.sh
sudo chmod +x /root/make-subscribe.sh


rm -rf /root/sing-box-telegram*
wget https://github.com/GhostOfSarina/sing-box-daily-telegram/releases/download/v.1.5.1/sing-box-telegram
sudo chmod +x ./sing-box-telegram





#instal monitoring
apt-get update
apt-get install nload
apt-get install htop
apt-get install iftop
apt-get install vnstat
apt-get install speedtest-cli
apt-get install net-tools
apt-get install -y jq

journalctl --vacuum-size=50M


# Check if reality.json, sing-box, and sing-box.service already exist
if [ -f "/root/reality.json" ] && [ -f "/root/sing-box" ] && [ -f "/etc/systemd/system/sing-box.service" ]; then

    echo "Reality files already exist."
    echo ""
    echo "Please choose an option:"
    echo ""
    echo "1. Reinstall"
    echo "2. Uninstall"
    echo ""
    read -p "Enter your choice (1-2): " choice

    case $choice in
        1)
            echo "Reinstalling..."
            # Uninstall previous installation
            systemctl stop sing-box
            systemctl disable sing-box
            rm /etc/systemd/system/sing-box.service
            rm /root/reality.json
            rm /root/sing-box
            rm /root/subscribe.txt
            rm /root/public_key.txt
            rm /root/sing-box-telegram


            # Proceed with installation
            ;;
        2)
            echo "Uninstalling..."
            # Stop and disable sing-box service
            systemctl stop sing-box
            systemctl disable sing-box

            # Remove files
            rm /etc/systemd/system/sing-box.service
            rm /root/reality.json
            rm /root/sing-box
            rm /root/subscribe.*
            rm -rf /var/www/hml/subscribe.*
            rm /root/public_key.txt
            rm /root/sing-box-telegram
            rm /root/first-time-install-sing-box.sh
            rm /root/reinstall-sing-box.sh
            rm /root/make-subscribe.sh

            echo "if you want to delete setting.json, please run rm /root/setting.json"


	    echo "DONE!"
            exit 0
            ;;
        *)
            echo "Invalid choice. Exiting."
            exit 1
            ;;
    esac
fi



# Fetch the latest (including pre-releases) release version number from GitHub API
latest_version_tag=$(curl -s "https://api.github.com/repos/SagerNet/sing-box/releases" | grep -Po '"tag_name": "\K.*?(?=")' | head -n 1)
latest_version=${latest_version_tag#v}  # Remove 'v' prefix from version number
echo "Latest version: $latest_version"
echo "Latest version tag: $latest_version_tag"
latest_version="1.3.0"

# Detect server architecture
arch=$(uname -m)

# Map architecture names
case ${arch} in
    x86_64)
        arch="amd64"
        ;;
    aarch64)
        arch="arm64"
        ;;
    armv7l)
        arch="armv7"
        ;;
esac

# Prepare package names
package_name="sing-box-${latest_version}-linux-${arch}"

# Download the latest release package (.tar.gz) from GitHub
curl -sLo "/root/${package_name}.tar.gz" "https://github.com/SagerNet/sing-box/releases/download/v${latest_version}/${package_name}.tar.gz"

# Extract the package and move the binary to /root
tar -xzf "/root/${package_name}.tar.gz" -C /root
mv "/root/${package_name}/sing-box" /root/

# Cleanup the package
rm -r "/root/${package_name}.tar.gz" "/root/${package_name}"

# Set the permissions
chown root:root /root/sing-box
chmod +x /root/sing-box


# Generate key pair
echo "Generating key pair..."
key_pair=$(/root/sing-box generate reality-keypair)
echo "Key pair generation complete."
echo

# Extract private key and public key
private_key=$(echo "$key_pair" | awk '/PrivateKey/ {print $2}' | tr -d '"')
public_key=$(echo "$key_pair" | awk '/PublicKey/ {print $2}' | tr -d '"')

# Generate necessary values
uuid=$(/root/sing-box generate uuid)
short_id=$(/root/sing-box generate rand --hex 8)


listen_port=443


server_name=www.datadoghq.com

# Retrieve the server IP address
server_ip=$(curl -s https://api.ipify.org)

# Create reality.json using jq
jq -n --arg listen_port "$listen_port" --arg server_name "$server_name" --arg private_key "$private_key" --arg short_id "$short_id" --arg uuid "$uuid" --arg server_ip "$server_ip" '{
  "inbounds": [
    {
      "type": "vless",
      "tag": "vless-in",
      "listen": "::",
      "listen_port": ($listen_port | tonumber),
      "sniff": true,
      "sniff_override_destination": true,
      "domain_strategy": "ipv4_only",
      "users": [
        {
          "uuid": $uuid,
          "flow": "xtls-rprx-vision"
        }
      ],
      "tls": {
        "enabled": true,
        "server_name": $server_name,
          "reality": {
          "enabled": true,
          "handshake": {
            "server": $server_name,
            "server_port": 443
          },
          "private_key": $private_key,
          "short_id": [$short_id]
        }
      }
    }
  ],
 "outbounds": [
      {
        "type": "direct"
      },
      {
        "type": "direct",
        "tag": "dns"
      },
      {
        "type": "block",
        "tag": "block"
      }
 ],
}' > /root/reality.json

# Create sing-box.service
cat > /etc/systemd/system/sing-box.service <<EOF
[Unit]
After=network.target nss-lookup.target

[Service]
User=root
WorkingDirectory=/root
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE CAP_NET_RAW
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE CAP_NET_RAW
ExecStart=/root/sing-box run -c /root/reality.json
ExecReload=/bin/kill -HUP \$MAINPID
Restart=on-failure
RestartSec=10
LimitNOFILE=infinity

[Install]
WantedBy=multi-user.target
EOF

#store public key in a file
touch /root/public_key.txt
echo $public_key > /root/public_key.txt


# Check configuration and start the service
if /root/sing-box check -c /root/reality.json; then
    echo "Configuration checked successfully. Starting sing-box service..."
    systemctl daemon-reload
    systemctl enable sing-box
    systemctl start sing-box
    systemctl restart sing-box

# Generate the link

    server_link="vless://$uuid@$server_ip:$listen_port?encryption=none&flow=xtls-rprx-vision&security=reality&sni=$server_name&fp=chrome&pbk=$public_key&sid=$short_id&type=tcp&headerType=none#SING-BOX-TCP"

    # Print the server details
    echo
    echo "Server IP: $server_ip"
    echo "Listen Port: $listen_port"
    echo "Server Name: $server_name"
    echo "Public Key: $public_key"
    echo "Short ID: $short_id"
    echo "UUID: $uuid"
    echo ""
    echo ""
    echo ""
    echo "Here is the link for v2rayN and v2rayNG :"
    echo ""
    echo "$server_link"

    touch /root/subscribe.txt
    echo $server_link > /root/subscribe.txt





    # Install apache2 and clone the website
    apt-get install apache2

    cd /var/www/html/
    git clone https://github.com/codingstella/vCard-personal-portfolio.git
    cp -ar ./vCard-personal-portfolio/*  /var/www/html/
    rm -rf ./vCard-personal-portfolio/
    cp /root/subscribe.txt /var/www/html/subscribe.txt 


    # Install cron job 
    croncmd="/root/sing-box-telegram > /root/cronjob.log 2>&1"
    cronjob="0 13 * * * $croncmd"
    ( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -


else
    echo "Error in configuration. Aborting."
fi

