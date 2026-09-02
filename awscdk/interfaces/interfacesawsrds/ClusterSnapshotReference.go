package interfacesawsrds


// A reference to a ClusterSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSnapshotReference := &ClusterSnapshotReference{
//   	DbClusterSnapshotArn: jsii.String("dbClusterSnapshotArn"),
//   }
//
type ClusterSnapshotReference struct {
	// The DBClusterSnapshotArn of the ClusterSnapshot resource.
	DbClusterSnapshotArn *string `field:"required" json:"dbClusterSnapshotArn" yaml:"dbClusterSnapshotArn"`
}

