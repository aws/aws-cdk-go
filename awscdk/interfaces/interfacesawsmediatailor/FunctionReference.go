package interfacesawsmediatailor


// A reference to a Function resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionReference := &FunctionReference{
//   	FunctionArn: jsii.String("functionArn"),
//   	FunctionId: jsii.String("functionId"),
//   }
//
type FunctionReference struct {
	// The ARN of the Function resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The FunctionId of the Function resource.
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
}

