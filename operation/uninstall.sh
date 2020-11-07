#!/bin/bash
#!/bin/bash
#
# uninstall.sh 1.0.0
#
# Removes the firewall configuration, uninstalls go, git and the festivals-server and stops and removes it as a service.
#
# (c)2020 Simon Gaus
#

systemctl stop festivals-server >/dev/null
echo "1. Stopped festivals-server"
sleep 1

systemctl disable festivals-server >/dev/null
rm /etc/systemd/system/festivals-server.service
echo "2. Removed systemd service"
sleep 1

# Disable and un-configure the firewall.
# Supported firewalls: ufw and firewalld
# This step is skipped under macOS.
#
if command -v firewalld > /dev/null; then

  firewall-cmd --permanent --remove-service=festivals-server >/dev/null
  rm -f /etc/firewalld/services/festivals-server.xml
  rm -f /etc/firewalld/services/festivals-server.xml.old
  firewall-cmd --reload >/dev/null

  echo "3. Removed firewalld configuration"
  sleep 1

elif command -v ufw > /dev/null; then

  ufw default deny incoming
  ufw default allow outgoing
  ufw allow OpenSSH
  yes | sudo ufw enable >/dev/null
  echo "Enabled ufw"
  sleep 1

  ufw allow 10439/tcp
  echo "Added festivals-server to ufw"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

















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
