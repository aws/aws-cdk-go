package awslightsail


// A reference to a InstanceSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceSnapshotReference := &InstanceSnapshotReference{
//   	InstanceSnapshotArn: jsii.String("instanceSnapshotArn"),
//   	InstanceSnapshotName: jsii.String("instanceSnapshotName"),
//   }
//
type InstanceSnapshotReference struct {
	// The ARN of the InstanceSnapshot resource.
	InstanceSnapshotArn *string `field:"required" json:"instanceSnapshotArn" yaml:"instanceSnapshotArn"`
	// The InstanceSnapshotName of the InstanceSnapshot resource.
	InstanceSnapshotName *string `field:"required" json:"instanceSnapshotName" yaml:"instanceSnapshotName"`
}

