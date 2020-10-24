#
#
#

cd /usr/local

systemctl enable firewalld >/dev/null
systemctl start firewalld >/dev/null
echo "1. Enabled firewalld"

firewall-cmd --permanent --new-service=festivals-server >/dev/null
firewall-cmd --permanent --service=festivals-server --set-description="A live and lightweight go server app providing the FestivalsAPI." >/dev/null
firewall-cmd --permanent --service=festivals-server --add-port=10439/tcp >/dev/null
firewall-cmd --permanent --add-service=festivals-server >/dev/null
firewall-cmd --reload >/dev/null
echo "2. Added festivals-server.service to firewalld"
sleep 1

echo "2. Downloading current go version..."
curl -o /usr/local/go.tar.gz "https://dl.google.com/go/$(curl --silent "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf /usr/local/go.tar.gz
rm /usr/local/go.tar.gz
ln -sf /usr/local/go/bin/* /usr/local/bin
echo "3. Updated go"
sleep 1

echo "4. Downloading unzip..."
dnf install unzip --assumeyes >/dev/null
echo "5. Installed unzip"
sleep 1

echo "6. Downloading current festivals-server..."
curl -L -o /usr/local/festivals-server.zip https://github.com/Festivals-App/festivals-server/archive/master.zip
unzip /usr/local/festivals-server.zip -d /usr/local >/dev/null
rm /usr/local/festivals-server.zip
cd /usr/local/festivals-server-master || exit
/usr/local/bin/go build main.go
echo "7. Build festivals-server"
sleep 1

mv main /usr/local/bin/festivals-server
restorecon -v /usr/local/bin/festivals-server >/dev/null
mv config_template.toml /etc/festivals-server.conf
echo "8. Installed festivals-server"
sleep 1

# create systemctl service
sudo tee -a /etc/systemd/system/festivals-server.service > /dev/null <<EOT
[Unit]
Description=FestivalsAPI server, a live and lightweight go server app.
ConditionPathExists=/usr/local/bin/festivals-server

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStartPre=/bin/mkdir -p /var/log/festivals-server
ExecStart=/usr/local/bin/festivals-server

[Install]
WantedBy=multi-user.target
EOT
echo "9. Created systemd service"
sleep 1

systemctl enable festivals-server
systemctl start festivals-server
echo "10. Enabled systemd service"
sleep 1

rm -R /usr/local/festivals-server-master
echo "11. Done"