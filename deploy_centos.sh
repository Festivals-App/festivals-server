#
#
#

systemctl enable firewalld >/dev/null
systemctl start firewalld >/dev/null
echo "1. Enabled firewalld"

firewall-cmd --permanent --new-service=festivals-server
firewall-cmd --permanent --service=festivals-server --set-description="A live and lightweight go server app providing the FestivalsAPI."
firewall-cmd --permanent --service=festivals-server --add-port=10439/tcp
firewall-cmd --permanent --add-service=festivals-server
firewall-cmd --reload
echo "2. Add festivals-server.service to firewalld"

curl -o go.tar.gz "https://dl.google.com/go/$(curl "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf go.tar.gz
rm go.tar.gz
ln -s /usr/local/go/bin/* /usr/local/bin
echo "3. Installed go"

# install repository
dnf install git --assumeyes >/dev/null
echo "4. Installed git"

go get github.com/Festivals-App/festivals-server >/dev/null
echo "5. Downloaded festivals-server"

cd ~/go/src/github.com/Festivals-App/festivals-server || exit
go build main.go
echo "6. Build festivals-server"

mv main /usr/local/bin/festivals-server
restorecon -v /usr/local/bin/festivals-server
mv config_template.toml /etc/festivals-server.conf
echo "7. Installed festivals-server"

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
echo "8. Created systemd service"

systemctl enable festivals-server
systemctl start festivals-server
echo "9. Enabled systemd service"

# cleanup after installation
rm -R ~/go


# remove this script
rm -- "$0"