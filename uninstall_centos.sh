systemctl stop festivals-server

firewall-cmd --permanent --remove-service=festivals-server
rm -f /etc/firewalld/services/festivals-server.xml
rm -f /etc/firewalld/services/festivals-server.xml.old
firewall-cmd --reload
firewall-cmd --info-service=festivals-server

rm -R /usr/local/go
rm /usr/local/bin/festivals-server
rm /etc/festivals-server.conf
rm -R /var/log/festivals-server
rm /usr/local/bin/go
rm /usr/local/bin/gofmt

dnf remove unzip --assumeyes

systemctl disable festivals-server
rm /etc/systemd/system/festivals-server.service