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
echo "Building the festivals-server..."
go build main.go
sleep 1
echo "Installing the festivals-server..."
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