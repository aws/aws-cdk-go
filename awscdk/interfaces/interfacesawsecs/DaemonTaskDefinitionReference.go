package interfacesawsecs


// A reference to a DaemonTaskDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   daemonTaskDefinitionReference := &DaemonTaskDefinitionReference{
//   	DaemonTaskDefinitionArn: jsii.String("daemonTaskDefinitionArn"),
//   }
//
type DaemonTaskDefinitionReference struct {
	// The DaemonTaskDefinitionArn of the DaemonTaskDefinition resource.
	DaemonTaskDefinitionArn *string `field:"required" json:"daemonTaskDefinitionArn" yaml:"daemonTaskDefinitionArn"`
}

