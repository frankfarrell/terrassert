# terrassert
Terraform provider for making assertions on the output of modules

This is a terraform provider for making assertions about outputs from modules and resources. 

## Motivation
The ultimate goal is to have mock terraform providers so that HCL terraform configuration can be unit tested. 
For example, something like this: 
```
provider "mock_aws" {
  region     = "us-east-1"
}

resource "mock_aws_when" "kinesis" {
    resource = "aws_kinesis_stream"
    then_return = {
        arns = ["firstArn", "secondArn]
        id = ["firstId", "secondId"]
    }
}

locals {
  foo = 12
  bar = "xyz"
}

resource "aws_kinesis_stream" "test_stream" {
  name             = "${local.bar}"
  shard_count      = "${local.foo}"
  retention_period = 48
}

data "terrassert_equals" "kinesis_name" {
    actual = "${aws_kinesis_stream.test_stream.name}"
    expected = "xyz"
    message = "Name should be xyz"
}

data "terrassert_gt" "kinesis_counte" {
    actual = "${aws_kinesis_stream.test_stream.shard_count}"
    gt = 1
    message = "Shard count shoudl be greater than 1"
}

```

This provider will just provide the 2 last data providers. I have no idea (yet) if the mock_aws provider will work or not. 