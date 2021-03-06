#!/bin/bash
#
# uninstall.sh 1.0.0
#
# Removes the firewall configuration, uninstalls go, git and the festivals-server and stops and removes it as a service.
#
# (c)2020 Simon Gaus
#

# Move to working directory
#
cd /usr/local || exit

# Stop the service
#
systemctl stop festivals-server >/dev/null
echo "Stopped festivals-server"
sleep 1

# Remove systemd configuration
#
systemctl disable festivals-server >/dev/null
rm /etc/systemd/system/festivals-server.service
echo "Removed systemd service"
sleep 1

# Remove the firewall configuration.
# This step is skipped under macOS.
#
if command -v ufw > /dev/null; then

  ufw delete allow 10439/tcp >/dev/null
  echo "Removed ufw configuration"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Remove go
#
apt-get --purge remove golang -y
apt autoremove -y
echo "Removed go"
sleep 1

# Remove festivals-server
#
rm /usr/local/bin/festivals-server
rm /etc/festivals-server.conf
rm -R /var/log/festivals-server
rm -R /usr/local/festivals-server
echo "Removed festivals-server"
sleep 1

echo "Done"
