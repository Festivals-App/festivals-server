#!/bin/bash
#
# update.sh 1.0.0
#
# Updates the festivals-server and restarts it.
#
# (c)2020-2022 Simon Gaus
#

# Move to working dir
#
cd /usr/local || exit

# Stop the festivals-server
#
systemctl stop festivals-server
echo "Stopped festivals-server"
sleep 1

# Install go if needed.
# Binaries linked to /usr/local/bin
#
if ! command -v go > /dev/null; then
  echo "Installing go..."
  apt-get install golang -y > /dev/null;
fi

# Install git if needed.
#
if ! command -v git > /dev/null; then
  echo "Installing git..."
  apt-get install git -y > /dev/null;
fi

# Updating festivals-server to the newest version
#
echo "Downloading current festivals-server..."
yes | sudo git clone https://github.com/Festivals-App/festivals-server.git /usr/local/festivals-server > /dev/null;
cd /usr/local/festivals-server || { echo "Failed to access working directory. Exiting." ; exit 1; }
go build main.go
mv main /usr/local/bin/festivals-server || { echo "Failed to install festivals-server binary. Exiting." ; exit 1; }
echo "Updated festivals-server binary."
sleep 1

# Removing unused files
#
echo "Cleanup..."
cd /usr/local || exit
rm -r /usr/local/festivals-server
sleep 1

echo "Done!"
sleep 1

echo "Please start the server manually by running 'systemctl start festivals-server' after you updated the configuration file at '/etc/festivals-server.conf'"
sleep 1
