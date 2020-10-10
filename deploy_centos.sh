#
#
#

# setup go
curl -o go.tar.gz "https://dl.google.com/go/$(curl "https://golang.org/VERSION?m=text").linux-amd64.tar.gz"
tar -C /usr/local -xf go.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> .bash_profile
source ~/.bash_profile

# install repositiory
dnf install git --assumeyes
go get github.com/festivals-app/festivalsals-server
