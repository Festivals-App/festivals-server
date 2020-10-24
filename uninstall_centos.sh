systemctl stop festivals-server >/dev/null
echo "1. Stop festivals-server"
sleep 1

systemctl disable festivals-server >/dev/null
rm /etc/systemd/system/festivals-server.service
echo "2. Remove systemd service"
sleep 1

firewall-cmd --permanent --remove-service=festivals-server >/dev/null
rm -f /etc/firewalld/services/festivals-server.xml
rm -f /etc/firewalld/services/festivals-server.xml.old
firewall-cmd --reload >/dev/null
echo "3. Remove firewalld configuration"
sleep 1

rm -R /usr/local/go
rm /usr/local/bin/go
rm /usr/local/bin/gofmt
echo "4. Remove go"
sleep 1

rm /usr/local/bin/festivals-server
rm /etc/festivals-server.conf
rm -R /var/log/festivals-server
echo "5. Remove festivals-server"
sleep 1

dnf remove unzip --assumeyes >/dev/null
echo "6. Uninstall unzip"
sleep 1

echo "7. Done"
