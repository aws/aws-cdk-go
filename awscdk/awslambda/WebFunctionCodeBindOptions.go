package awslambda


// Options passed by `WebFunction` when binding code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionCodeBindOptions := &WebFunctionCodeBindOptions{
//   	WebFunctionArn: jsii.String("webFunctionArn"),
//   }
//
type WebFunctionCodeBindOptions struct {
	// The ARN of the WebFunction the code is bound to.
	//
	// Used to scope the code-read policy to this function's revisions.
	WebFunctionArn *string `field:"required" json:"webFunctionArn" yaml:"webFunctionArn"`
}

