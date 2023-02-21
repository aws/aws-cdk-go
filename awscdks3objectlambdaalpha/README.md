# AWS::S3ObjectLambda Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This construct library allows you to define S3 object lambda access points.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"

stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_14_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &AccessPointProps{
	Bucket: Bucket,
	Handler: Handler,
	AccessPointName: jsii.String("my-access-point"),
	Payload: map[string]interface{}{
		"prop": jsii.String("value"),
	},
})
```

## Handling range and part number requests

Lambdas are currently limited to only transforming `GetObject` requests. However, they can additionally support `GetObject-Range` and `GetObject-PartNumber` requests, which needs to be specified in the access point configuration:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"

stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_14_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &AccessPointProps{
	Bucket: Bucket,
	Handler: Handler,
	AccessPointName: jsii.String("my-access-point"),
	SupportsGetObjectRange: jsii.Boolean(true),
	SupportsGetObjectPartNumber: jsii.Boolean(true),
})
```

## Pass additional data to Lambda function

You can specify an additional object that provides supplemental data to the Lambda function used to transform objects. The data is delivered as a JSON payload to the Lambda:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"

stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_14_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &AccessPointProps{
	Bucket: Bucket,
	Handler: Handler,
	AccessPointName: jsii.String("my-access-point"),
	Payload: map[string]interface{}{
		"prop": jsii.String("value"),
	},
})
```
