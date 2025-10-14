package awslightsail


// A reference to a DiskSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskSnapshotReference := &DiskSnapshotReference{
//   	DiskSnapshotArn: jsii.String("diskSnapshotArn"),
//   	DiskSnapshotName: jsii.String("diskSnapshotName"),
//   }
//
type DiskSnapshotReference struct {
	// The ARN of the DiskSnapshot resource.
	DiskSnapshotArn *string `field:"required" json:"diskSnapshotArn" yaml:"diskSnapshotArn"`
	// The DiskSnapshotName of the DiskSnapshot resource.
	DiskSnapshotName *string `field:"required" json:"diskSnapshotName" yaml:"diskSnapshotName"`
}

