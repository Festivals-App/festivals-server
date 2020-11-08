#!/bin/bash
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
# Supported firewalls: ufw and firewalld
# This step is skipped under macOS.
#
if command -v firewalld > /dev/null; then

  firewall-cmd --permanent --remove-service=festivals-server >/dev/null
  rm -f /etc/firewalld/services/festivals-server.xml
  rm -f /etc/firewalld/services/festivals-server.xml.old
  firewall-cmd --reload >/dev/null
  echo "Removed firewalld configuration"
  sleep 1

elif command -v ufw > /dev/null; then

  ufw delete allow 10439/tcp
  echo "Removed ufw configuration"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Remove go
#
rm -R /usr/local/go
rm /usr/local/bin/go
rm /usr/local/bin/gofmt
echo "Removed go"
sleep 1

# Remove festivals-server
#
rm /usr/local/bin/festivals-server
rm /etc/festivals-server.conf
rm -R /var/log/festivals-server
echo "Removed festivals-server"
sleep 1

echo "Done"
