#
#
#

# configure firewall-cmd
firewall-cmd --permanent --new-service=festivals-server
firewall-cmd --permanent --service=festivals-server --set-description="A live and lightweight go server app providing the FestivalsAPI."
firewall-cmd --permanent --service=festivals-server --add-port=10439/tcp
firewall-cmd --permanent --add-service=festivals-server
firewall-cmd --reload

# setup go
curl -o go.tar.gz "https://dl.google.com/go/$(curl "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf go.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> .bash_profile
source ~/.bash_profile

# install repositiory
dnf install git --assumeyes
go get github.com/Festivals-App/festivals-server
cd ~/go/src/github.com/Festivals-App/festivals-server
go build main.go
mv main /usr/local/bin/festivals-server
mv config_template.toml /etc/festivals-server.conf

# create sytemctl service
echo 'test text here' | sudo tee -a /etc/systemd/system/festivals-server.service
echo 'test text here' | sudo tee -a /etc/systemd/system/festivals-server.service
