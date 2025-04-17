package awscdks3objectlambdaalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The S3 object lambda access point configuration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import s3objectlambda "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   stack := cdk.NewStack()
//   bucket := s3.NewBucket(stack, jsii.String("MyBucket"))
//   handler := lambda.NewFunction(stack, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(jsii.String("lambda.zip")),
//   })
//   s3objectlambda.NewAccessPoint(stack, jsii.String("MyObjectLambda"), &AccessPointProps{
//   	Bucket: Bucket,
//   	Handler: Handler,
//   	AccessPointName: jsii.String("my-access-point"),
//   	Payload: map[string]interface{}{
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
	// Default: a unique name will be generated.
	//
	// Experimental.
	AccessPointName *string `field:"optional" json:"accessPointName" yaml:"accessPointName"`
	// Whether CloudWatch metrics are enabled for the access point.
	// Default: false.
	//
	// Experimental.
	CloudWatchMetricsEnabled *bool `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// Additional JSON that provides supplemental data passed to the Lambda function on every request.
	// Default: - No data.
	//
	// Experimental.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Whether the Lambda function can process `GetObject-PartNumber` requests.
	// Default: false.
	//
	// Experimental.
	SupportsGetObjectPartNumber *bool `field:"optional" json:"supportsGetObjectPartNumber" yaml:"supportsGetObjectPartNumber"`
	// Whether the Lambda function can process `GetObject-Range` requests.
	// Default: false.
	//
	// Experimental.
	SupportsGetObjectRange *bool `field:"optional" json:"supportsGetObjectRange" yaml:"supportsGetObjectRange"`
}

