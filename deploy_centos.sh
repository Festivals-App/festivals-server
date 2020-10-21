#
#
#

# launch on startup and launch firewalld
systemctl enable firewalld
systemctl start firewalld

# configure firewall-cmd
firewall-cmd --permanent --new-service=festivals-server
firewall-cmd --permanent --service=festivals-server --set-description="A live and lightweight go server app providing the FestivalsAPI."
firewall-cmd --permanent --service=festivals-server --add-port=10439/tcp
firewall-cmd --permanent --add-service=festivals-server
firewall-cmd --reload

# setup go
curl -o go.tar.gz "https://dl.google.com/go/$(curl "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf go.tar.gz
rm go.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> .bash_profile
source ~/.bash_profile

# install repository
dnf install git --assumeyes
go get github.com/Festivals-App/festivals-server
cd ~/go/src/github.com/Festivals-App/festivals-server || exit
go build main.go
mv main /usr/local/bin/festivals-server
mv config_template.toml /etc/festivals-server.conf

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

systemctl enable festivals-server
systemctl start festivals-server

# cleanup after installation
rm -R ~/go

# remove this script
rm -- "$0"