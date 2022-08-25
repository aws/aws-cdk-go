package awscloudfront


// Contains metadata about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionMetadataProperty := &functionMetadataProperty{
//   	functionArn: jsii.String("functionArn"),
//   }
//
type CfnFunction_FunctionMetadataProperty struct {
	// The Amazon Resource Name (ARN) of the function.
	//
	// The ARN uniquely identifies the function.
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

