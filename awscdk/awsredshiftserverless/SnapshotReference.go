package awsredshiftserverless


// A reference to a Snapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotReference := &SnapshotReference{
//   	SnapshotName: jsii.String("snapshotName"),
//   }
//
type SnapshotReference struct {
	// The SnapshotName of the Snapshot resource.
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
}

