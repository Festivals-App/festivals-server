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
mkdir /usr/local/festivals-server || { echo "Failed to create working directory. Exiting." ; exit 1; }
cd /usr/local/festivals-server || { echo "Failed to access working directory. Exiting." ; exit 1; }

# Stop the festivals-server
#
systemctl stop festivals-server
echo "Stopped festivals-server"
sleep 1

# Get system os
#
if [ "$(uname -s)" = "Darwin" ]; then
  os="darwin"
elif [ "$(uname -s)" = "Linux" ]; then
  os="linux"
else
  echo "System is not Darwin or Linux. Exiting."
  exit 1
fi

# Get systems cpu architecture
#
if [ "$(uname -m)" = "x86_64" ]; then
  arch="amd64"
elif [ "$(uname -m)" = "arm64" ]; then
  arch="arm64"
else
  echo "System is not x86_64 or arm64. Exiting."
  exit 1
fi

# Build url to latest binary for the given system
#
file_url="https://github.com/Festivals-App/festivals-server/releases/latest/download/festivals-server-$os-$arch.tar.gz"
echo "The system is $os on $arch."
sleep 1

# Updating festivals-server to the newest binary release
#
echo "Downloading newest festivals-server binary release..."
curl -L "$file_url" -o festivals-server.tar.gz
tar -xf festivals-server.tar.gz
mv festivals-server /usr/local/bin/festivals-server || { echo "Failed to install festivals-server binary. Exiting." ; exit 1; }
echo "Updated festivals-server binary."
sleep 1

# Removing unused files
#
echo "Cleanup..."
cd /usr/local
rm -r /usr/local/festivals-server
sleep 1

# Stop the festivals-server
#
systemctl start festivals-server
echo "Started festivals-server"
sleep 1

echo "Done!"
sleep 1