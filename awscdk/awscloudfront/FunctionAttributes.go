package awscloudfront


// Attributes of an existing CloudFront Function to import it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionAttributes := &FunctionAttributes{
//   	FunctionArn: jsii.String("functionArn"),
//   	FunctionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	FunctionRuntime: jsii.String("functionRuntime"),
//   }
//
type FunctionAttributes struct {
	// The ARN of the function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The name of the function.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The Runtime of the function.
	// Default: FunctionRuntime.JS_1_0
	//
	FunctionRuntime *string `field:"optional" json:"functionRuntime" yaml:"functionRuntime"`
}

