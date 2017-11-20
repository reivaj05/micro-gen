# Microservices generator [![Build Status](https://travis-ci.org/reivaj05/micro-gen.svg?branch=master)](https://travis-ci.org/reivaj05/micro-gen) [![codecov](https://codecov.io/gh/reivaj05/micro-gen/branch/master/graph/badge.svg)](https://codecov.io/gh/reivaj05/micro-gen)


## Description

Command line tool to create microservices for different languages (Golang, Javascript, Python, Ruby, Rust).

### Current support:
 - Golang
 - Python
 - Rust

## Usage

 - Install [Golang](https://golang.org/doc/install)

```
$ go get github.com/reivaj05/micro-gen
$ micro-gen create-service [serviceName] --lang [serviceLanguage (default go)]
```
## Example:

```
$ micro-gen create-service micro-example --lang go
```

 - The last command will create a directory called micro-example using go as primary language (check the available languages list above if you want to use a different one).

 - It will also create docker files to easily create an image of the service. Just run the following command and your service will be up and running so you can start working on it

```
$ docker-compose up micro-example
```

 - Go to [localhost](http://localhost:8000), and you'll hit a simple rest endpoint which you can modify according to your needs.
 
 - Travis files will also be created inside the new directory, you just need to do is (check auto created travis file for more information)

	 - Setup a new repo locally and in Github (or any other provider).
	 - Create a travis account if you don't have one.
	 - Enable your new Github repo in Travis.
	 - Commit your changes and push your first commit, this will trigger a new build inside travis.

