package awscloudfront


// Contains metadata about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionMetadataProperty := &FunctionMetadataProperty{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionmetadata.html
//
type CfnFunction_FunctionMetadataProperty struct {
	// The Amazon Resource Name (ARN) of the function.
	//
	// The ARN uniquely identifies the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionmetadata.html#cfn-cloudfront-function-functionmetadata-functionarn
	//
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

