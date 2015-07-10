# -*- mode: ruby -*-
# vi: set ft=ruby :

# Specify Vagrant version and Vagrant API version
VAGRANTFILE_API_VERSION = "2"

# This script is executed as "root" when the builder box is provisioned
$script = <<SCRIPT

# install docker: https://docs.docker.com/installation/ubuntulinux/#ubuntu-trusty-1404-lts-64-bit
apt-get update
apt-get install wget -y
wget -qO- https://get.docker.com/ | sh

# use Docker as a non-root user
usermod -aG docker vagrant

# add docker-registry as a local name (fake dns)
echo "127.0.1.1 docker-registry" >> /etc/hosts
echo 'DOCKER_OPTS="--insecure-registry docker-registry:5000"' >> /etc/default/docker
service docker restart

# install and run docker registry: https://github.com/docker/docker-registry
mkdir /var/registry
chmod a+w /var/registry
docker run --name registry --restart=always -v "/var/registry:/registry" -e STORAGE_PATH=/registry -p 5000:5000 -d registry

# pre-fetch and cache the go-lang image
docker pull golang:1.4.2

SCRIPT

# Create and configure the VM(s)
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network "private_network", ip: "192.168.50.5"
  # http://serverfault.com/questions/453185/vagrant-virtualbox-dns-10-0-2-3-not-working
  config.vm.provider :virtualbox do |vb|
  	vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end
  config.vm.provision "shell", inline: $script
  # Disable synced folders (prevents an NFS error on "vagrant up"): http://blog.scottlowe.org/2015/02/10/using-docker-with-vagrant/
  #config.vm.synced_folder ".", "/vagrant", disabled: true
end
