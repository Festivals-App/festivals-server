#!/bin/bash
#
# install.sh 1.0.0
#
# Enables the firewall, installs the newest go and the festivals-server and starts it as a service.
#
# (c)2020-2022 Simon Gaus
#

# Test for web server user
#
WEB_USER="www-data"
id -u "$WEB_USER" &>/dev/null;
if [ $? -ne 0 ]; then
  WEB_USER="www"
  if [ $? -ne 0 ]; then
    echo "Failed to find user to run web server. Exiting."
    exit 1
  fi
fi

# Move to working dir
#
mkdir /usr/local/festivals-server || { echo "Failed to create working directory. Exiting." ; exit 1; }
cd /usr/local/festivals-server || { echo "Failed to access working directory. Exiting." ; exit 1; }

echo "Installing festivals-server using port 10439."
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

# Install festivals-server to /usr/local/bin/festivals-server. TODO: Maybe just link to /usr/local/bin?
#
echo "Downloading newest festivals-server binary release..."
curl -L "$file_url" -o festivals-server.tar.gz
tar -xf festivals-server.tar.gz
mv festivals-server /usr/local/bin/festivals-server || { echo "Failed to install festivals-server binary. Exiting." ; exit 1; }
echo "Installed the festivals-server binary to '/usr/local/bin/festivals-server'."
mv config_template.toml /etc/festivals-server.conf
echo "Moved default festivals-server config to '/etc/festivals-server.conf'."
sleep 1
mkdir /var/log/festivals-server || { echo "Failed to create log directory. Exiting." ; exit 1; }
chown "$WEB_USER":"$WEB_USER" /var/log/festivals-server
echo "Create log directory at '/var/log/festivals-server'."


# Enable and configure the firewall.
#
if command -v ufw > /dev/null; then

  ufw allow 10439/tcp >/dev/null
  echo "Added festivals-server to ufw using port 10439."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Install systemd service
#
if command -v service > /dev/null; then

  if ! [ -f "/etc/systemd/system/festivals-server.service" ]; then
    mv service_template.service /etc/systemd/system/festivals-server.service
    echo "Created systemd service."
    sleep 1
  fi

  systemctl enable festivals-server > /dev/null
  echo "Enabled systemd service."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "Systemd is missing and not on macOS. Exiting."
  exit 1
fi

# Remving unused files
#
echo "Cleanup..."
cd /usr/local || exit
rm -R /usr/local/festivals-server
sleep 1

echo "Done!"
sleep 1

echo "You can start the server manually by running 'systemctl start festivals-server' after you updated the configuration file at '/etc/festivals-server.conf'"
sleep 1
