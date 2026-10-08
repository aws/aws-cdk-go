package interfacesawslambda


// A reference to a WebFunction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionReference := &WebFunctionReference{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
type WebFunctionReference struct {
	// The FunctionArn of the WebFunction resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

