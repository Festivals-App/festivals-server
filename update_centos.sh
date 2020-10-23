systemctl stop festivals-server
echo "1. Stopped festivals-server"

curl -o go.tar.gz "https://dl.google.com/go/$(curl "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf go.tar.gz
rm go.tar.gz
ln -sfn /usr/local/go/bin/* /usr/local/bin
echo "2. Updated go"

cd ~/go/src/github.com/Festivals-App/festivals-server || exit
go build main.go
echo "3. Build updated festivals-server"

mv main /usr/local/bin/festivals-server
restorecon -v /usr/local/bin/festivals-server
echo "4. Installed festivals-server"

systemctl start festivals-server
echo "5. Enabled systemd service"

rm -R ~/go
echo "6. Cleaning up after updating"
sleep 2

# remove this script
rm -- "$0"