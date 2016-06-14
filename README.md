# GO WATSON:
-----------

![pablo](https://cloud.githubusercontent.com/assets/8342133/16011734/430826da-31a4-11e6-9c6d-9c316637edf8.png)


## About:
---------

Simple golang project to build up your developer 
enviorment of Watson on your favorite OS(Operating System).

## Features:
-----------

* Can install [Watson](http://github.com/IBM-Watson) on Linux
* Can install a variety of tools related to [Watson](http://github.com/IBM-Watson)
* Can be built up on [Docker](http://github.com/docker/docker)
* Built using [Golang](http://github.com/golang/go)

## Behind the Curtains:
-----------------------

Go-watson works by testing your enviorment if it meets the following prerequisites:
* Docker
* Npm
* Yeoman Generator

After these prerequisites are met it installs the watson npm package to your system.
Using the `yo watson` you can start watson on your enviorment.Thereafter, you can 
install various utilities of watson,presently:
* Kale [CLI tool]
* Runner [Task Runner]
* A11y.js [JS helpers]

## How does it work:
-------------------

Currently the watson enviorment can be installed and run using 2 methods:

### Using Github repo:

* Download the repository `https://github.com/ramitsurana/go-watson`
* `cd go-watson`
* `go run *.go`

### Using Docker:

* You can use `docker run -i -t ramitsurana/go-watson`
* Run `yo watson`

## License:
----------

The work done above is Licensed under [Apache License](LICENSE).

