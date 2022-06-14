# AWS::S3ObjectLambda Construct Library

This construct library allows you to define S3 object lambda access points.

```go
// Example automatically generated from non-compiling source. May contain errors.
import lambda "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"


stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &accessPointProps{
	bucket: bucket,
	handler: handler,
	accessPointName: jsii.String("my-access-point"),
	payload: map[string]interface{}{
		"prop": jsii.String("value"),
	},
})
```

## Handling range and part number requests

Lambdas are currently limited to only transforming `GetObject` requests. However, they can additionally support `GetObject-Range` and `GetObject-PartNumber` requests, which needs to be specified in the access point configuration:

```go
// Example automatically generated from non-compiling source. May contain errors.
import lambda "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"


stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &accessPointProps{
	bucket: bucket,
	handler: handler,
	accessPointName: jsii.String("my-access-point"),
	supportsGetObjectRange: jsii.Boolean(true),
	supportsGetObjectPartNumber: jsii.Boolean(true),
})
```

## Pass additional data to Lambda function

You can specify an additional object that provides supplemental data to the Lambda function used to transform objects. The data is delivered as a JSON payload to the Lambda:

```go
// Example automatically generated from non-compiling source. May contain errors.
import lambda "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
import s3objectlambda "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"


stack := cdk.NewStack()
bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(jsii.String("lambda.zip")),
})
s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &accessPointProps{
	bucket: bucket,
	handler: handler,
	accessPointName: jsii.String("my-access-point"),
	payload: map[string]interface{}{
		"prop": jsii.String("value"),
	},
})
```
