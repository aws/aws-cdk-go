package awsfsx


// The OpenZFS configuration for the file system that's being created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openZFSConfigurationProperty := &openZFSConfigurationProperty{
//   	deploymentType: jsii.String("deploymentType"),
//
//   	// the properties below are optional
//   	automaticBackupRetentionDays: jsii.Number(123),
//   	copyTagsToBackups: jsii.Boolean(false),
//   	copyTagsToVolumes: jsii.Boolean(false),
//   	dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	diskIopsConfiguration: &diskIopsConfigurationProperty{
//   		iops: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   	},
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   	rootVolumeConfiguration: &rootVolumeConfigurationProperty{
//   		copyTagsToSnapshots: jsii.Boolean(false),
//   		dataCompressionType: jsii.String("dataCompressionType"),
//   		nfsExports: []interface{}{
//   			&nfsExportsProperty{
//   				clientConfigurations: []interface{}{
//   					&clientConfigurationsProperty{
//   						clients: jsii.String("clients"),
//   						options: []*string{
//   							jsii.String("options"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		readOnly: jsii.Boolean(false),
//   		recordSizeKiB: jsii.Number(123),
//   		userAndGroupQuotas: []interface{}{
//   			&userAndGroupQuotasProperty{
//   				id: jsii.Number(123),
//   				storageCapacityQuotaGiB: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	throughputCapacity: jsii.Number(123),
//   	weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_OpenZFSConfigurationProperty struct {
	// Specifies the file system deployment type.
	//
	// Amazon FSx for OpenZFS supports `SINGLE_AZ_1` . `SINGLE_AZ_1` deployment type is configured for redundancy within a single Availability Zone.
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean value indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToVolumes interface{} `field:"optional" json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
	//
	// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
	DiskIopsConfiguration interface{} `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// To delete a file system if there are child volumes present below the root volume, use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS` .
	//
	// If your file system has child volumes and you don't use this option, the delete request will fail.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS file system.
	//
	// All volumes are children of the root volume.
	RootVolumeConfiguration interface{} `field:"optional" json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MB/s).
	//
	// Valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MB/s. You pay for additional throughput capacity that you provision.
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

