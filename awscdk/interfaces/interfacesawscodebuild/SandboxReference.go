package interfacesawscodebuild


// A reference to a Sandbox resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sandboxReference := &SandboxReference{
//   	SandboxArn: jsii.String("sandboxArn"),
//   }
//
type SandboxReference struct {
	// The Arn of the Sandbox resource.
	SandboxArn *string `field:"required" json:"sandboxArn" yaml:"sandboxArn"`
}

