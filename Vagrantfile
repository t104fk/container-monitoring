# -*- mode: ruby -*-
# # vi: set ft=ruby :

Vagrant.require_version ">= 1.6.0"
VAGRANTFILE_API_VERSION = "2"

LOCAL_DB_IP="172.17.8.101"
LOCAL_APP_IP="172.17.8.102"
LOCAL_MONITOR_IP="172.17.8.103"

SHARED_DIR="/home/vagrant/share"

$update_channel = "stable"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.ssh.insert_key = false
  config.vm.define "db" do |cfg|
    cfg.vm.box = "ubuntu/trusty64"
    cfg.vm.box_url = "https://atlas.hashicorp.com/ubuntu/boxes/trusty64"
    cfg.vm.host_name = "database.vm"
    cfg.vm.network :forwarded_port, guest:5432, host:5432
    cfg.vm.network :private_network, ip: LOCAL_DB_IP
    cfg.vm.synced_folder ".", SHARED_DIR, id: "core", :nfs => true, :mount_options => ['nolock,vers=3,udp']
  end
  config.vm.define "app" do |cfg|
    cfg.vm.box = "coreos-%s" % $update_channel
    cfg.vm.box_url = "http://%s.release.core-os.net/amd64-usr/current/coreos_production_vagrant.json" % $update_channel
    cfg.vm.host_name = "localapp.vm"
    cfg.vm.network :private_network, ip: LOCAL_APP_IP
    cfg.vm.synced_folder ".", SHARED_DIR, id: "core", :nfs => true, :mount_options => ['nolock,vers=3,udp']
  end
  config.vm.define "monitor" do |cfg|
    cfg.vm.box = "ubuntu/trusty64"
    cfg.vm.box_url = "https://atlas.hashicorp.com/ubuntu/boxes/trusty64"
    cfg.vm.host_name = "monitor.vm"
    cfg.vm.network :private_network, ip: LOCAL_MONITOR_IP
  end
end

