package interfacesawslambda


// A reference to a Version resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versionReference := &VersionReference{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
type VersionReference struct {
	// The FunctionArn of the Version resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

