package interfacesawselasticache


// A reference to a Snapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotReference := &SnapshotReference{
//   	SnapshotArn: jsii.String("snapshotArn"),
//   }
//
type SnapshotReference struct {
	// The Arn of the Snapshot resource.
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

