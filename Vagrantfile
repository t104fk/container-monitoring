# -*- mode: ruby -*-
# # vi: set ft=ruby :

Vagrant.require_version ">= 1.6.0"
VAGRANTFILE_API_VERSION = "2"

LOCAL_DB_IP="172.17.9.101"
LOCAL_APP_IP="172.17.9.102"
LOCAL_MONITOR_IP="172.17.9.103"

SHARED_DIR="/home/vagrant/container-monitoring"

$stop_container = <<SCRIPT
RUNNINGS=`docker ps -q --filter='status=running'`
for c in ${RUNNINGS[@]}
do
  if [ ! -z $c ]; then
    docker stop $c
  fi
done
SCRIPT

$clean_container = <<SCRIPT
RUNNINGS=`docker ps -q --filter='status=running'`
for c in ${RUNNINGS[@]}
do
  if [ ! -z $c ]; then
    docker stop $c
    docker rm $c
  fi
done
SCRIPT

$cadvisor_cmd = <<EOS.chomp
-storage_driver=influxdb \
-storage_driver_db=cadvisor \
-storage_driver_host=#{LOCAL_MONITOR_IP}:8086 \
-storage_driver_user=cadvisor \
-storage_driver_password=p4kU7ugpndm9V9Lw \
-storage_driver_secure=False
EOS

$cadvisor_args = <<EOS.chomp
-v /:/rootfs:ro \
-v /var/run:/var/run:rw \
-v /sys:/sys:ro \
-v /var/lib/docker/:/var/lib/docker:ro \
-p 8080:8080 
EOS

$fluentd_sender_args = <<EOS.chomp
-v /var/lib/docker/containers:/var/lib/docker/containers \
-v /var/log/docker:/var/log/docker
EOS

$fluentd_receiver_args = <<EOS.chomp
-v /var/lib/docker/containers:/var/lib/docker/containers \
-v /var/log/docker:/var/log/docker \
-p 24224:24224 \
-p 24224:24224/udp
EOS

$grafana_args = <<EOS.chomp
-p 3000:3000 \
-e INFLUXDB_HOST=localhost \
-e INFLUXDB_PORT=8086 \
-e INFLUXDB_NAME=cadvisor \
-e INFLUXDB_USER=cadvisor \
-e INFLUXDB_PASS=p4kU7ugpndm9V9Lw \
-e INFLUXDB_IS_GRAFANADB=true
EOS

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.ssh.insert_key = false
  config.vm.box = "ubuntu/trusty64"
  config.vm.box_url = "https://atlas.hashicorp.com/ubuntu/boxes/trusty64"
  config.vm.define "db" do |cfg|
    cfg.vm.host_name = "database.vm"
    cfg.vm.network :private_network, ip: LOCAL_DB_IP
    cfg.vm.synced_folder ".", SHARED_DIR, id: "shared", :nfs => true, :mount_options => ['nolock,vers=3,udp']
  end
  config.vm.define "app" do |cfg|
    cfg.vm.host_name = "localapp.vm"
    cfg.vm.network :private_network, ip: LOCAL_APP_IP
    cfg.vm.synced_folder ".", SHARED_DIR, id: "shared", :nfs => true, :mount_options => ['nolock,vers=3,udp']
    cfg.vm.provision "api-build", type: "docker" do |d|
      d.build_image "#{SHARED_DIR}/app",
        args: "-t takasing/api:0.0.1"
    end

    cfg.vm.provision "clean-container", type: "shell", inline: $clean_container

    cfg.vm.provision "api-run", type: "docker", run: "always" do |d|
      d.run "api",
        image: "takasing/api:0.0.1",
        args: "-p 3000:3000",
        auto_assign_name: true,
        daemonize: true
    end

    cfg.vm.provision "cadvisor", type: "docker", run: "always" do |d|
      d.run "cadvisor",
        image: "google/cadvisor:0.15.1",
        cmd: $cadvisor_cmd,
        args: $cadvisor_args,
        auto_assign_name: true,
        daemonize: true
    end

    cfg.vm.provision "fluentd", type: "docker", run: "always" do |d|
      d.build_image "#{SHARED_DIR}/docker/fluentd/sender", args: "-t takasing/fluentd:0.0.1"
      d.run "fluentd",
        image: "takasing/fluentd:0.0.1",
        args: $fluentd_sender_args,
        auto_assign_name: true,
        daemonize: true
    end
  end
  config.vm.define "monitor" do |cfg|
    cfg.vm.host_name = "monitor.vm"
    cfg.vm.network :private_network, ip: LOCAL_MONITOR_IP
    cfg.vm.synced_folder ".", SHARED_DIR, id: "shared", :nfs => true, :mount_options => ['nolock,vers=3,udp']
    cfg.vm.provision "influxdb", type: "ansible" do |ansible|
      ansible.playbook = "ansible/monitor.yml"
    end

    cfg.vm.provision "stop-container", type: "shell", inline: $stop_container

    cfg.vm.provision "grafana", type: "docker", run: "always" do |d|
      d.run "grafana",
        image: "grafana/grafana:latest",
        args: $grafana_args,
        auto_assign_name: true,
        daemonize: true
    end

    cfg.vm.provision "fluentd", type: "docker", run: "always" do |d|
      d.build_image "#{SHARED_DIR}/docker/fluentd/receiver", args: "-t takasing/fluentd:0.0.1"
      d.run "fluentd",
        image: "takasing/fluentd:0.0.1",
        args: $fluentd_receiver_args,
        auto_assign_name: true,
        daemonize: true
    end
  end
end

