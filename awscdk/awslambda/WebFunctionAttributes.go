package awslambda


// Attributes for importing an existing WebFunction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionAttributes := &WebFunctionAttributes{
//   	WebFunctionArn: jsii.String("webFunctionArn"),
//   }
//
type WebFunctionAttributes struct {
	// The ARN of the WebFunction.
	WebFunctionArn *string `field:"required" json:"webFunctionArn" yaml:"webFunctionArn"`
}

