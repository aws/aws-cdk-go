package awsiottwinmaker


// The Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionProperty := &lambdaFunctionProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnComponentType_LambdaFunctionProperty struct {
	// The Lambda function ARN.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

