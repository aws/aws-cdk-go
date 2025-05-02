package awsfsx


// The configuration object that specifies the snapshot to use as the origin of the data for the volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originSnapshotProperty := &OriginSnapshotProperty{
//   	CopyStrategy: jsii.String("copyStrategy"),
//   	SnapshotArn: jsii.String("snapshotArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-originsnapshot.html
//
type CfnVolume_OriginSnapshotProperty struct {
	// Specifies the strategy used when copying data from the snapshot to the new volume.
	//
	// - `CLONE` - The new volume references the data in the origin snapshot. Cloning a snapshot is faster than copying data from the snapshot to a new volume and doesn't consume disk throughput. However, the origin snapshot can't be deleted if there is a volume using its copied data.
	// - `FULL_COPY` - Copies all data from the snapshot to the new volume.
	//
	// Specify this option to create the volume from a snapshot on another FSx for OpenZFS file system.
	//
	// > The `INCREMENTAL_COPY` option is only for updating an existing volume by using a snapshot from another FSx for OpenZFS file system. For more information, see [CopySnapshotAndUpdateVolume](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CopySnapshotAndUpdateVolume.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-originsnapshot.html#cfn-fsx-volume-originsnapshot-copystrategy
	//
	CopyStrategy *string `field:"required" json:"copyStrategy" yaml:"copyStrategy"`
	// Specifies the snapshot to use when creating an OpenZFS volume from a snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-originsnapshot.html#cfn-fsx-volume-originsnapshot-snapshotarn
	//
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

