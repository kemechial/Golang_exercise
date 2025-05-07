Vagrant.configure("2") do |config|
  # Set the VM name
  config.vm.define "gohost" do |gohost|

    # Use Ubuntu 22.04 as the base image
    gohost.vm.box = "ubuntu/jammy64"

    # Set up a synchronized folder (host to guest)
    gohost.vm.synced_folder "C:\\GoExercise\\files", "/tmp/files"
        
    # Provisioning to install Go
    gohost.vm.provision "shell", inline: <<-SHELL
      # Update package list
      sudo apt-get update

      # Install prerequisites
      sudo apt-get install -y wget tar

      # Download and install Go
      GO_VERSION=1.20.5
      wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
      sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
      rm go${GO_VERSION}.linux-amd64.tar.gz

      # Set up Go environment variables
      echo "export PATH=\\$PATH:/usr/local/go/bin" | sudo tee /etc/profile.d/go.sh
      echo "export GOPATH=\\$HOME/go" | sudo tee -a /etc/profile.d/go.sh
      echo "export PATH=\\$PATH:\\$GOPATH/bin" | sudo tee -a /etc/profile.d/go.sh
      sudo chmod +x /etc/profile.d/go.sh

      # Source the profile so Go is available immediately
      source /etc/profile.d/go.sh
    SHELL

    # Set up a private network with a static IP
    gohost.vm.network "private_network", ip: "192.168.56.10"
  end
end
