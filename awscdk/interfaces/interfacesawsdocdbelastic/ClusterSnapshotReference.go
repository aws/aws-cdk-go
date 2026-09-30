package interfacesawsdocdbelastic


// A reference to a ClusterSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSnapshotReference := &ClusterSnapshotReference{
//   	SnapshotArn: jsii.String("snapshotArn"),
//   }
//
type ClusterSnapshotReference struct {
	// The SnapshotArn of the ClusterSnapshot resource.
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

