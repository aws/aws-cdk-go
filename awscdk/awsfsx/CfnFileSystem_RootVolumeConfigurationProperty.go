package awsfsx


// The configuration of an Amazon FSx for OpenZFS root volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rootVolumeConfigurationProperty := &RootVolumeConfigurationProperty{
//   	CopyTagsToSnapshots: jsii.Boolean(false),
//   	DataCompressionType: jsii.String("dataCompressionType"),
//   	NfsExports: []interface{}{
//   		&NfsExportsProperty{
//   			ClientConfigurations: []interface{}{
//   				&ClientConfigurationsProperty{
//   					Clients: jsii.String("clients"),
//   					Options: []*string{
//   						jsii.String("options"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ReadOnly: jsii.Boolean(false),
//   	RecordSizeKiB: jsii.Number(123),
//   	UserAndGroupQuotas: []interface{}{
//   		&UserAndGroupQuotasProperty{
//   			Id: jsii.Number(123),
//   			StorageCapacityQuotaGiB: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html
//
type CfnFileSystem_RootVolumeConfigurationProperty struct {
	// A Boolean value indicating whether tags for the volume should be copied to snapshots of the volume.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-copytagstosnapshots
	//
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Specifies the method used to compress the data on the volume. The compression type is `NONE` by default.
	//
	// - `NONE` - Doesn't compress the data on the volume. `NONE` is the default.
	// - `ZSTD` - Compresses the data in the volume using the Zstandard (ZSTD) compression algorithm. Compared to LZ4, Z-Standard provides a better compression ratio to minimize on-disk storage utilization.
	// - `LZ4` - Compresses the data in the volume using the LZ4 compression algorithm. Compared to Z-Standard, LZ4 is less compute-intensive and delivers higher write throughput speeds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-datacompressiontype
	//
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// The configuration object for mounting a file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-nfsexports
	//
	NfsExports interface{} `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// A Boolean value indicating whether the volume is read-only.
	//
	// Setting this value to `true` can be useful after you have completed changes to a volume and no longer want changes to occur.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Specifies the record size of an OpenZFS root volume, in kibibytes (KiB).
	//
	// Valid values are 4, 8, 16, 32, 64, 128, 256, 512, or 1024 KiB. The default is 128 KiB. Most workloads should use the default record size. Database workflows can benefit from a smaller record size, while streaming workflows can benefit from a larger record size. For additional guidance on setting a custom record size, see [Tips for maximizing performance](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance.html#performance-tips-zfs) in the *Amazon FSx for OpenZFS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-recordsizekib
	//
	RecordSizeKiB *float64 `field:"optional" json:"recordSizeKiB" yaml:"recordSizeKiB"`
	// An object specifying how much storage users or groups can use on the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-rootvolumeconfiguration.html#cfn-fsx-filesystem-rootvolumeconfiguration-userandgroupquotas
	//
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

