package interfacesawsfsx


// A reference to a Snapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotReference := &SnapshotReference{
//   	SnapshotId: jsii.String("snapshotId"),
//   }
//
type SnapshotReference struct {
	// The Id of the Snapshot resource.
	SnapshotId *string `field:"required" json:"snapshotId" yaml:"snapshotId"`
}

