package awsfsx


// The configuration object that specifies the snapshot to use as the origin of the data for the volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originSnapshotProperty := &originSnapshotProperty{
//   	copyStrategy: jsii.String("copyStrategy"),
//   	snapshotArn: jsii.String("snapshotArn"),
//   }
//
type CfnVolume_OriginSnapshotProperty struct {
	// The strategy used when copying data from the snapshot to the new volume.
	//
	// - `CLONE` - The new volume references the data in the origin snapshot. Cloning a snapshot is faster than copying data from the snapshot to a new volume and doesn't consume disk throughput. However, the origin snapshot can't be deleted if there is a volume using its copied data.
	// - `FULL_COPY` - Copies all data from the snapshot to the new volume.
	CopyStrategy *string `field:"required" json:"copyStrategy" yaml:"copyStrategy"`
	// Specifies the snapshot to use when creating an OpenZFS volume from a snapshot.
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

