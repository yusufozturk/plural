rm -rf $WORKSPACE/*
export GOPATH=$WORKSPACE/
export PATH=$PATH:/usr/local/bin

go get github.com/spf13/viper
go get github.com/shirou/gopsutil
go get github.com/dustin/go-humanize
go get github.com/fsouza/go-dockerclient
go get github.com/drael/GOnetstat

cd src
git clone https://github.com/marshyski/plural.git
cd plural/plural && go build

rm -rf /opt/plural

mkdir -p /opt/plural/{bin,conf}

mv plural /opt/plural/bin

sudo -n mv -f ../build/plural.init.el6 /etc/init.d/plural

fpm -s dir -t rpm -n "plural" -v $BUILD_NUMBER --rpm-user root --rpm-group root --rpm-compression bzip2 /opt/plural /etc/init.d/plural

fpm -s dir -t deb -n "plural" -v $BUILD_NUMBER --deb-user root --deb-group root --deb-compression bzip2 /opt/plural
