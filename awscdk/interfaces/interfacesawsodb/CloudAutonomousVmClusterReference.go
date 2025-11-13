package interfacesawsodb


// A reference to a CloudAutonomousVmCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudAutonomousVmClusterReference := &CloudAutonomousVmClusterReference{
//   	CloudAutonomousVmClusterArn: jsii.String("cloudAutonomousVmClusterArn"),
//   }
//
type CloudAutonomousVmClusterReference struct {
	// The CloudAutonomousVmClusterArn of the CloudAutonomousVmCluster resource.
	CloudAutonomousVmClusterArn *string `field:"required" json:"cloudAutonomousVmClusterArn" yaml:"cloudAutonomousVmClusterArn"`
}

