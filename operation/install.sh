#!/bin/bash
#
# install.sh 1.0.0
#
# Enables the firewall, installs the newest go and the festivals-server and starts it as a service.
#
# (c)2020 Simon Gaus
#

# Move to working directory
#
cd /usr/local || exit

# Enable and configure the firewall.
# Supported firewalls: ufw and firewalld
# This step is skipped under macOS.
#
if command -v firewalld > /dev/null; then

  systemctl enable firewalld >/dev/null
  systemctl start firewalld >/dev/null
  echo "Enabled firewalld"
  sleep 1

  firewall-cmd --permanent --new-service=festivals-server >/dev/null
  firewall-cmd --permanent --service=festivals-server --set-description="A live and lightweight go server app providing the FestivalsAPI." >/dev/null
  firewall-cmd --permanent --service=festivals-server --add-port=10439/tcp >/dev/null
  firewall-cmd --permanent --add-service=festivals-server >/dev/null
  firewall-cmd --reload >/dev/null
  echo "Added festivals-server.service to firewalld"
  sleep 1

elif command -v ufw > /dev/null; then

  ufw default deny incoming >/dev/null
  ufw default allow outgoing >/dev/null
  ufw allow OpenSSH >/dev/null
  yes | sudo ufw enable >/dev/null
  echo "Enabled ufw"
  sleep 1

  ufw allow 10439/tcp >/dev/null
  echo "Added festivals-server to ufw"
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "No firewall detected and not on macOS. Exiting."
  exit 1
fi

# Install go to /usr/local/go if needed.
# Binaries linked to /usr/local/bin
#
if ! command -v go > /dev/null; then
  echo "Downloading current go version..."
  goVersion="$(curl --silent "https://golang.org/VERSION?m=text")"
  currentGo="$goVersion.linux-amd64.tar.gz"
  goURL="https://dl.google.com/go/$currentGo"
  goOut=/var/cache/festivals-server/$currentGo
  if ! [ -f "/etc/systemd/system/festivals-server.service" ]; then
    curl --progress-bar -o goOut goURL
  else
    echo "Using cached go package at $goOut"
    sleep 1
  fi
  tar -C /usr/local -xf goOut
  ln -sf /usr/local/go/bin/* /usr/local/bin
  echo "Installed go ($currentGo)"
  sleep 1
fi

# Install git if needed.
#
if ! command -v git > /dev/null; then
  if command -v dnf > /dev/null; then
    echo "Installing git"
    dnf install git -y > /dev/null;
  elif command -v apt > /dev/null; then
    echo "Installing git"
    apt install git -y > /dev/null;
  else
    echo "Unable to install git. Exiting."
    sleep 1
    exit 1
  fi
else
  echo "Already installed git"
fi

# Install festivals-server to /usr/local/bin/festivals-server. TODO: Maybe just link to /usr/local/bin?
#
echo "Downloading current festivals-server..."
yes | sudo git clone https://github.com/Festivals-App/festivals-server.git /usr/local/festivals-server
cd /usr/local/festivals-server || exit
/usr/local/bin/go build main.go
mv main /usr/local/bin/festivals-server
if command -v restorecon > /dev/null; then
  restorecon -v /usr/local/bin/festivals-server >/dev/null
fi
mv config_template.toml /etc/festivals-server.conf
echo "Installed festivals-server."
sleep 1

# Install systemd service
#
if command -v service > /dev/null; then

  if ! [ -f "/etc/systemd/system/festivals-server.service" ]; then
    mv operation/service_template.service /etc/systemd/system/festivals-server.service
    echo "Created systemd service."
    sleep 1
  fi

  systemctl enable festivals-server > /dev/null
  systemctl start festivals-server > /dev/null
  echo "Enabled systemd service."
  sleep 1

elif ! [ "$(uname -s)" = "Darwin" ]; then
  echo "Systemd is missing and not on macOS. Exiting."
  exit 1
fi

echo "Done."