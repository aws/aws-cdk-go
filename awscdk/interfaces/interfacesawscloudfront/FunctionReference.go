package interfacesawscloudfront


// A reference to a Function resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionReference := &FunctionReference{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
type FunctionReference struct {
	// The FunctionARN of the Function resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

