# Setting up development environment

## On your local machine
1. Make sure you have golang 1.10 [installed](https://golang.org/doc/install#download)
1. Make sure you have `dep` [installed](https://github.com/golang/dep)
1. Make sure you have `kubebuilder` [installed](https://book.kubebuilder.io/getting_started/installation_and_setup.html) (currently only required to run some of the tests)
1. Clone this repository into `$GOPATH/src/github.com/DevFactory/smartnat`
1. Run `dep`
  ```
  dep ensure --vendor-only
  ```
5. You should be able to run unit tests now and you're good to go
  ```
  go test ./pkg/...
  ```

## In docker container
If you don't want to install all the gp dependencies on your machine, you can easily get started by setting up development environment in a docker container.

Start by building docker image:
```
docker build -f ./Dockerfile.devenv -t smartnat_devenv:latest .
```

Now, you can start a ready container that has all the dependecies, code and vim ready:
```
docker run -it --name smartnat-dev smartnat_devenv:latest
```

Everything should be in place and you should be able to run all tests:
```
go test ./...
```
