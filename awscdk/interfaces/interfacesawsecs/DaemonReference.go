package interfacesawsecs


// A reference to a Daemon resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   daemonReference := &DaemonReference{
//   	DaemonArn: jsii.String("daemonArn"),
//   }
//
type DaemonReference struct {
	// The DaemonArn of the Daemon resource.
	DaemonArn *string `field:"required" json:"daemonArn" yaml:"daemonArn"`
}

