# Terraform Provider for Medium

[Medium](https://medium.com) is an online publishing platform. This provider
allows a Medium user with write API access to write a story in Markdown and
upload it to Medium.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.12.x (to build the provider plugin)
- [Developer access to the Medium
  API](https://github.com/Medium/medium-api-docs#21-browser-based-authentication)
- [Integration token for your user](https://medium.com/me/settings)    

## Building The Provider

Clone repository to:
`$GOPATH/src/github.com/joatmon08/terraform-provider-medium`

```sh
$ mkdir -p $GOPATH/src/github.com/joatmon08; cd $GOPATH/src/github.com/joatmon08
$ git clone git@github.com:joatmon08/terraform-provider-medium
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/joatmon08/terraform-provider-medium
$ make build
```

## Developing the Provider

### Contributing Resources

### Development Environment

If you wish to work on the provider, you'll first need
[Go](http://www.golang.org) installed on your machine (version 1.9+ is
*required*). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding
`$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put
the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-medium
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```