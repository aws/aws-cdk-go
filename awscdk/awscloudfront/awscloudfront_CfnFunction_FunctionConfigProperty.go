package awscloudfront


// Contains configuration information about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionConfigProperty := &functionConfigProperty{
//   	comment: jsii.String("comment"),
//   	runtime: jsii.String("runtime"),
//   }
//
type CfnFunction_FunctionConfigProperty struct {
	// A comment to describe the function.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// The function’s runtime environment.
	//
	// The only valid value is `cloudfront-js-1.0` .
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
}

