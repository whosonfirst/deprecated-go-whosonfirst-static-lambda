# go-whosonfirst-static-lambda

Run the `go-whosonfirst-static` code in an AWS Lambda function.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.7 so let's just assume you need [Go 1.9](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

This is _super_ early work and does not yet support all the features of the `go-whosonfirst-static` package. Here is a short list of things this package lacks:

* Documentation, specifically about how to set things up as a Lambda
  function. Have a look at the
  [go-rasterzen-lambda](https://github.com/whosonfirst/go-rasterzen-lambda)
  package. It's basically the same thing.

* Support for reading data from anything but an S3 source.

* Support for HTML output.

* Support for toggling output. Currently all `graphics` and `data` endpoints are
  enabled by default.

## See also

### Package specific

* https://github.com/whosonfirst/go-whosonfirst-static
* https://github.com/whosonfirst/go-whosonfirst-readwrite
* https://github.com/whosonfirst/go-whosonfirst-readwrite-s3
* https://github.com/whosonfirst/algnhsa
* https://artem.krylysov.com/blog/2018/01/18/porting-go-web-applications-to-aws-lambda/

### AWS

* https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/
* https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html
* https://docs.aws.amazon.com/lambda/latest/dg/env_variables.html
* https://docs.aws.amazon.com/cli/latest/reference/sts/get-session-token.html

### General

* https://apimeister.com/2017/05/09/hosting-a-cloudfront-site-with-s3-and-api-gateway.html
* https://github.com/tilezen/tapalcatl-py#lambda-gotchas
