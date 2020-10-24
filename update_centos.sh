#
#
#

cd /usr/local

systemctl stop festivals-server
echo "1. Stopped festivals-server"

echo "2. Downloading current go version"
curl -o /usr/local/go.tar.gz "https://dl.google.com/go/$(curl --silent "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf /usr/local/go.tar.gz
rm /usr/local/go.tar.gz
ln -sf /usr/local/go/bin/* /usr/local/bin
echo "3. Updated go"

echo "4. Download current festivals-server"
curl -L -o /usr/local/festivals-server.zip https://github.com/Festivals-App/festivals-server/archive/master.zip >/dev/null
unzip /usr/local/festivals-server.zip -d /usr/local >/dev/null
rm /usr/local/festivals-server.zip
echo "5. Downloaded current festivals-server"
sleep 1

cd /usr/local/festivals-server-master || exit
/usr/local/bin/go build main.go
echo "6. Build festivals-server"
sleep 1

mv main /usr/local/bin/festivals-server
restorecon -v /usr/local/bin/festivals-server
echo "7. Installed festivals-server"
sleep 1

systemctl start festivals-server
echo "8. Enabled systemd service"
sleep 1

cd /usr/local
rm -R /usr/local/festivals-server-master
echo "9. Cleaning up after updating"
sleep 1