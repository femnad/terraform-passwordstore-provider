# terraform-provider-passwordstore

Terraform provider for reading contents of [pass](https://www.passwordstore.org/) secrets, only implements [data source](https://www.terraform.io/docs/configuration/data-sources.html) functionality

## Better Alternatives

https://github.com/camptocamp/terraform-provider-pass

## Why This Might be A Bad Idea

Because everything provided by a data source ends up in the state file and that means if you are sharing state via a backend then the credentials end up in the backend storage, defeating the purpose of using fine grained permissions for resource provisioning and state storage.
