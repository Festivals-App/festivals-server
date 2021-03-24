#!/bin/bash
#
#

# Move to working dir
#
cd /usr/local || exit

# Stop the festivals-server
#
systemctl stop festivals-server
echo "Stopped festivals-server"
sleep 1

# Updating go to the newest version
#
echo "Updating go..."
goVersion="$(curl --silent "https://golang.org/VERSION?m=text")"
currentGo="$goVersion.linux-amd64.tar.gz"
goURL="https://dl.google.com/go/$currentGo"
goOut=/var/cache/festivals-server/$currentGo
if ! [ -f "$goOut" ]; then
    mkdir -p /var/cache/festivals-server >/dev/null || { echo "Failed to create cache directory. Exiting." ; exit 1; }
    curl --progress-bar -o "$goOut" "$goURL" || { echo "Failed to download go. Exiting." ; exit 1; }
else
    echo "Using cached go package at $goOut"
    sleep 1
fi
tar -C /usr/local -xf "$goOut"
ln -sf /usr/local/go/bin/* /usr/local/bin
echo "Updated go to ($currentGo)"
sleep 1

# Install unzip if needed.
#
if ! command -v unzip > /dev/null; then
  echo "Installing unzip..."
  apt-get install unzip -y > /dev/null;
fi

# Updating festivals-server to the newest version
#
echo "Downloading the current festivals-server"
curl --progress-bar -L -o /usr/local/festivals-server.zip https://github.com/Festivals-App/festivals-server/archive/master.zip
unzip /usr/local/festivals-server.zip -d /usr/local >/dev/null
cd /usr/local/festivals-server-master || exit
echo "Building the festivals-server"
/usr/local/bin/go build main.go
sleep 1
echo "Installing the festivals-server"
mv main /usr/local/bin/festivals-server
sleep 1

# Updating go to the newest version
#
systemctl start festivals-server
echo "Started festivals-server"
sleep 1

# Remving used files
#
echo "Cleanup..."
cd /usr/local || exit
rm -R /usr/local/festivals-server-master
rm /usr/local/festivals-server.zip
sleep 1

echo "Done!"
sleep 1