aws-inventory
================

**Ansible dynamic inventory plugin for use with AWS EC2.**

The Ansible repository already contains an [EC2 dynamic inventory plugin](https://github.com/ansible/ansible/tree/devel/plugins/inventory), which may be useful to you.

This plugin differs from the 'official' one in the following ways:
 
 * For each host, sets the `ansible_ssh_host` variable using the public ip. This eliminates the need to reference hosts by their ip, or maintain your `/etc/hosts` file. You can then create another inventory file in the same directory, and reference the hosts by their EC2 Tag `Name`.
 
 * Returns host variables in the `_meta` top level element, reducing the number of api calls to AWS and speeding up the provisioning process. This eliminates the need to call the executable with `--host` for each host.
 
 * Only makes 2 requests to the AWS API when called with `--list`.
 
 * No external dependencies.
 
 * Creates less variables per host, but adding more would be trival. Open a pull-request if you need one defined. 

See [Developing Dynamic Inventory Sources](http://docs.ansible.com/developing_inventory.html) for more information.

## Download

**Grab the latest release from [Releases](https://github.com/awilliams/aws-inventory/releases)**

## Usage

 * Download the appropriate package from releases.
 
 * Place the executable inside your ansible directory, alongside other inventory files in a directory or wherever you like.

 * Create a `aws-inventory.ini` file with your AWS credentials, in the same directory as the executable. See the inlcluded example ini file `aws-inventory.example.ini`.

 * Test the output

 `./aws-inventory --list`

## Building

A hacked together `Makefile` is included. Try:

    make
    make package
