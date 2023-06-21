package awsfsx


// The OpenZFS configuration for the file system that's being created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openZFSConfigurationProperty := &OpenZFSConfigurationProperty{
//   	DeploymentType: jsii.String("deploymentType"),
//
//   	// the properties below are optional
//   	AutomaticBackupRetentionDays: jsii.Number(123),
//   	CopyTagsToBackups: jsii.Boolean(false),
//   	CopyTagsToVolumes: jsii.Boolean(false),
//   	DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   		Iops: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   	},
//   	Options: []*string{
//   		jsii.String("options"),
//   	},
//   	RootVolumeConfiguration: &RootVolumeConfigurationProperty{
//   		CopyTagsToSnapshots: jsii.Boolean(false),
//   		DataCompressionType: jsii.String("dataCompressionType"),
//   		NfsExports: []interface{}{
//   			&NfsExportsProperty{
//   				ClientConfigurations: []interface{}{
//   					&ClientConfigurationsProperty{
//   						Clients: jsii.String("clients"),
//   						Options: []*string{
//   							jsii.String("options"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		ReadOnly: jsii.Boolean(false),
//   		RecordSizeKiB: jsii.Number(123),
//   		UserAndGroupQuotas: []interface{}{
//   			&UserAndGroupQuotasProperty{
//   				Id: jsii.Number(123),
//   				StorageCapacityQuotaGiB: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ThroughputCapacity: jsii.Number(123),
//   	WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_OpenZFSConfigurationProperty struct {
	// Specifies the file system deployment type.
	//
	// Single AZ deployment types are configured for redundancy within a single Availability Zone in an AWS Region . Valid values are the following:
	//
	// - `SINGLE_AZ_1` - (Default) Creates file systems with throughput capacities of 64 - 4,096 MBps. `Single_AZ_1` is available in all AWS Regions where Amazon FSx for OpenZFS is available.
	// - `SINGLE_AZ_2` - Creates file systems with throughput capacities of 160 - 10,240 MB/s using an NVMe L2ARC cache. `Single_AZ_2` is available only in the US East (N. Virginia), US East (Ohio), US West (Oregon), and Europe (Ireland) AWS Regions .
	//
	// For more information, see: [Deployment type availability](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/availability-durability.html#available-aws-regions) and [File system performance](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance.html#zfs-fs-performance) in the *Amazon FSx for OpenZFS User Guide* .
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `30` .
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean value indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A Boolean value indicating whether tags for the file system should be copied to volumes.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to volumes where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to volumes. If you specify one or more tags when creating the volume, no tags are copied from the file system, regardless of this value.
	CopyTagsToVolumes interface{} `field:"optional" json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or FSx for OpenZFS file system.
	//
	// By default, Amazon FSx automatically provisions 3 IOPS per GB of storage capacity. You can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how it is was provisioned, or the mode (by the customer or by Amazon FSx).
	DiskIopsConfiguration interface{} `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// To delete a file system if there are child volumes present below the root volume, use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS` .
	//
	// If your file system has child volumes and you don't use this option, the delete request will fail.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS file system.
	//
	// All volumes are children of the root volume.
	RootVolumeConfiguration interface{} `field:"optional" json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MBps).
	//
	// Valid values depend on the DeploymentType you choose, as follows:
	//
	// - For `SINGLE_AZ_1` , valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MBps.
	// - For `SINGLE_AZ_2` , valid values are 160, 320, 640, 1280, 2560, 3840, 5120, 7680, or 10240 MBps.
	//
	// You pay for additional throughput capacity that you provision.
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

