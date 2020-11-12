Terraform Provider
==================

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) `>=0.12.7` (as this project is using [the standalone SDK](https://www.terraform.io/docs/extend/plugin-sdk.html))
- [Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/pritunl/terraform-provider-pritunl`

```sh
$ mkdir -p $GOPATH/src/github.com/pritunl/terraform-provider-pritunl; cd $GOPATH/src/github.com/pritunl/terraform-provider-pritunl
$ git clone git@github.com:ricardclau/terraform-provider-pritunl.git .
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/pritunl/terraform-provider-pritunl
$ go install
```

If you check your `$GOPATH/bin` folder you should see a freshly compiled binary of this provider

And now, in order to get terraform auto-discover mechanism to work you need to create a `~/.terraformrc` file with this content

```hcl
providers {
  pritunl = "${GOPATH}/bin/terraform-provider-pritunl"
}
```

Using The Provider
---------------------

Terraform code example

```hcl
provider "pritunl" {
  pritunl_host   = "<pritunl_host>"
  pritunl_token  = "<pritunl_api_token>"
  pritunl_secret = "<pritunl_api_secret>"
}

resource "pritunl_organization" "cevinio" {
  name = "cevinio"
}

resource "pritunl_user" "ricard" {
  organization_id = pritunl_organization.cevinio.id
  name            = "ricardclau"
  email           = "ricard.clau@cevinio.com"
  auth_type       = "local"
  pin             = "11111"
}

data "pritunl_server" "test" {
  name = "Test"
}

resource "pritunl_route" "client" {
  server  = data.pritunl_server.test.id
  network = "10.255.0.0/16"
  comment = "Some Client VPC"
}
```

In order to execute this just run:

```sh
> terraform init

> terraform plan

> terraform apply
```

Every time you build a new version of the provider you will have to run


```sh
> terraform init
```

So that the .terraform folder with your plugins state is updated with the new binary