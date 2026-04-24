package interfacesawsneptunegraph


// A reference to a GraphSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphSnapshotReference := &GraphSnapshotReference{
//   	GraphSnapshotArn: jsii.String("graphSnapshotArn"),
//   }
//
type GraphSnapshotReference struct {
	// The Arn of the GraphSnapshot resource.
	GraphSnapshotArn *string `field:"required" json:"graphSnapshotArn" yaml:"graphSnapshotArn"`
}

