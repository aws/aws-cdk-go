package awsiottwinmaker


// The Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionProperty := &LambdaFunctionProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-lambdafunction.html
//
type CfnComponentType_LambdaFunctionProperty struct {
	// The Lambda function ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-lambdafunction.html#cfn-iottwinmaker-componenttype-lambdafunction-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

