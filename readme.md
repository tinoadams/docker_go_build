# Run the box

Prerequisites:
  - [Vagrant](http://www.vagrantup.com/downloads.html)
  - [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

Install Vagrant and VirtualBox on your machine and then execute the following:

```
$ git clone https://github.com/tinoadams/docker_go_build.git
$ cd docker_go_build
$ vagrant up
$ vagrant ssh
```

# Compile and run the Go hello world example

```
$ cd /vagrant/go_code
$ ./docker_run build
$ ./run app
```

# Compile and run the example using a custom Docker image

```
$ cd /vagrant/custom_image
$ ./docker_run build
$ ./run app
```
