# Operation

The `operation` directory contains all configuration templates and scripts to install and run the festvials-server.

* `install.sh` script to install festvials-server on a VM
* `service_template.service` festivals server unit file for `systemctl`
* `ufw_app_profile` firewall app profile file for `ufw`
* `update.sh` script to update the festivals-server

## Deployment

Follow the [**deployment guide**](DEPLOYMENT.md) for deploying the festivals-database inside a virtual machine or the [**local deployment guide**](./local/README.md) for running it on your macOS developer machine.
