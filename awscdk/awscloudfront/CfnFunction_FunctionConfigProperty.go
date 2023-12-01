package awscloudfront


// Contains configuration information about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionConfigProperty := &FunctionConfigProperty{
//   	Comment: jsii.String("comment"),
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html
//
type CfnFunction_FunctionConfigProperty struct {
	// A comment to describe the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html#cfn-cloudfront-function-functionconfig-comment
	//
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// The function's runtime environment version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html#cfn-cloudfront-function-functionconfig-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
}

