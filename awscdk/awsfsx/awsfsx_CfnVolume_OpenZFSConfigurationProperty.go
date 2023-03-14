package awsfsx


// Specifies the configuration of the Amazon FSx for OpenZFS volume that you are creating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openZFSConfigurationProperty := &openZFSConfigurationProperty{
//   	parentVolumeId: jsii.String("parentVolumeId"),
//
//   	// the properties below are optional
//   	copyTagsToSnapshots: jsii.Boolean(false),
//   	dataCompressionType: jsii.String("dataCompressionType"),
//   	nfsExports: []interface{}{
//   		&nfsExportsProperty{
//   			clientConfigurations: []interface{}{
//   				&clientConfigurationsProperty{
//   					clients: jsii.String("clients"),
//   					options: []*string{
//   						jsii.String("options"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   	originSnapshot: &originSnapshotProperty{
//   		copyStrategy: jsii.String("copyStrategy"),
//   		snapshotArn: jsii.String("snapshotArn"),
//   	},
//   	readOnly: jsii.Boolean(false),
//   	recordSizeKiB: jsii.Number(123),
//   	storageCapacityQuotaGiB: jsii.Number(123),
//   	storageCapacityReservationGiB: jsii.Number(123),
//   	userAndGroupQuotas: []interface{}{
//   		&userAndGroupQuotasProperty{
//   			id: jsii.Number(123),
//   			storageCapacityQuotaGiB: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnVolume_OpenZFSConfigurationProperty struct {
	// The ID of the volume to use as the parent volume of the volume that you are creating.
	ParentVolumeId *string `field:"required" json:"parentVolumeId" yaml:"parentVolumeId"`
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Specifies the method used to compress the data on the volume. The compression type is `NONE` by default.
	//
	// - `NONE` - Doesn't compress the data on the volume. `NONE` is the default.
	// - `ZSTD` - Compresses the data in the volume using the Zstandard (ZSTD) compression algorithm. Compared to LZ4, Z-Standard provides a better compression ratio to minimize on-disk storage utilization.
	// - `LZ4` - Compresses the data in the volume using the LZ4 compression algorithm. Compared to Z-Standard, LZ4 is less compute-intensive and delivers higher write throughput speeds.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// The configuration object for mounting a Network File System (NFS) file system.
	NfsExports interface{} `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// To delete the volume's child volumes, snapshots, and clones, use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS` .
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// The configuration object that specifies the snapshot to use as the origin of the data for the volume.
	OriginSnapshot interface{} `field:"optional" json:"originSnapshot" yaml:"originSnapshot"`
	// A Boolean value indicating whether the volume is read-only.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Specifies the suggested block size for a volume in a ZFS dataset, in kibibytes (KiB).
	//
	// Valid values are 4, 8, 16, 32, 64, 128, 256, 512, or 1024 KiB. The default is 128 KiB. We recommend using the default setting for the majority of use cases. Generally, workloads that write in fixed small or large record sizes may benefit from setting a custom record size, like database workloads (small record size) or media streaming workloads (large record size). For additional guidance on when to set a custom record size, see [ZFS Record size](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance.html#record-size-performance) in the *Amazon FSx for OpenZFS User Guide* .
	RecordSizeKiB *float64 `field:"optional" json:"recordSizeKiB" yaml:"recordSizeKiB"`
	// Sets the maximum storage size in gibibytes (GiB) for the volume.
	//
	// You can specify a quota that is larger than the storage on the parent volume. A volume quota limits the amount of storage that the volume can consume to the configured amount, but does not guarantee the space will be available on the parent volume. To guarantee quota space, you must also set `StorageCapacityReservationGiB` . To *not* specify a storage capacity quota, set this to `-1` .
	//
	// For more information, see [Volume properties](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/managing-volumes.html#volume-properties) in the *Amazon FSx for OpenZFS User Guide* .
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// Specifies the amount of storage in gibibytes (GiB) to reserve from the parent volume.
	//
	// Setting `StorageCapacityReservationGiB` guarantees that the specified amount of storage space on the parent volume will always be available for the volume. You can't reserve more storage than the parent volume has. To *not* specify a storage capacity reservation, set this to `0` or `-1` . For more information, see [Volume properties](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/managing-volumes.html#volume-properties) in the *Amazon FSx for OpenZFS User Guide* .
	StorageCapacityReservationGiB *float64 `field:"optional" json:"storageCapacityReservationGiB" yaml:"storageCapacityReservationGiB"`
	// An object specifying how much storage users or groups can use on the volume.
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

