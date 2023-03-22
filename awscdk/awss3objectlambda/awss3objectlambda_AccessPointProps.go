package awss3objectlambda

import (
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// The S3 object lambda access point configuration.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import s3objectlambda "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stack := cdk.NewStack()
//   bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
//   handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(jsii.String("lambda.zip")),
//   })
//   s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &accessPointProps{
//   	bucket: bucket,
//   	handler: handler,
//   	accessPointName: jsii.String("my-access-point"),
//   	payload: map[string]interface{}{
//   		"prop": jsii.String("value"),
//   	},
//   })
//
// Experimental.
type AccessPointProps struct {
	// The bucket to which this access point belongs.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The Lambda function used to transform objects.
	// Experimental.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// The name of the S3 object lambda access point.
	// Experimental.
	AccessPointName *string `field:"optional" json:"accessPointName" yaml:"accessPointName"`
	// Whether CloudWatch metrics are enabled for the access point.
	// Experimental.
	CloudWatchMetricsEnabled *bool `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// Additional JSON that provides supplemental data passed to the Lambda function on every request.
	// Experimental.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Whether the Lambda function can process `GetObject-PartNumber` requests.
	// Experimental.
	SupportsGetObjectPartNumber *bool `field:"optional" json:"supportsGetObjectPartNumber" yaml:"supportsGetObjectPartNumber"`
	// Whether the Lambda function can process `GetObject-Range` requests.
	// Experimental.
	SupportsGetObjectRange *bool `field:"optional" json:"supportsGetObjectRange" yaml:"supportsGetObjectRange"`
}

