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
//   	EndpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   	Options: []*string{
//   		jsii.String("options"),
//   	},
//   	PreferredSubnetId: jsii.String("preferredSubnetId"),
//   	ReadCacheConfiguration: &ReadCacheConfigurationProperty{
//   		SizeGiB: jsii.Number(123),
//   		SizingMode: jsii.String("sizingMode"),
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
//   	RouteTableIds: []*string{
//   		jsii.String("routeTableIds"),
//   	},
//   	ThroughputCapacity: jsii.Number(123),
//   	WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html
//
type CfnFileSystem_OpenZFSConfigurationProperty struct {
	// Specifies the file system deployment type. Valid values are the following:.
	//
	// - `MULTI_AZ_1` - Creates file systems with high availability and durability by replicating your data and supporting failover across multiple Availability Zones in the same AWS Region .
	// - `SINGLE_AZ_HA_2` - Creates file systems with high availability and throughput capacities of 160 - 10,240 MB/s using an NVMe L2ARC cache by deploying a primary and standby file system within the same Availability Zone.
	// - `SINGLE_AZ_HA_1` - Creates file systems with high availability and throughput capacities of 64 - 4,096 MB/s by deploying a primary and standby file system within the same Availability Zone.
	// - `SINGLE_AZ_2` - Creates file systems with throughput capacities of 160 - 10,240 MB/s using an NVMe L2ARC cache that automatically recover within a single Availability Zone.
	// - `SINGLE_AZ_1` - Creates file systems with throughput capacities of 64 - 4,096 MBs that automatically recover within a single Availability Zone.
	//
	// For a list of which AWS Regions each deployment type is available in, see [Deployment type availability](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/availability-durability.html#available-aws-regions) . For more information on the differences in performance between deployment types, see [File system performance](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance.html#zfs-fs-performance) in the *Amazon FSx for OpenZFS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-deploymenttype
	//
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `30` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-automaticbackupretentiondays
	//
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean value indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-copytagstobackups
	//
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A Boolean value indicating whether tags for the file system should be copied to volumes.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to volumes where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to volumes. If you specify one or more tags when creating the volume, no tags are copied from the file system, regardless of this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-copytagstovolumes
	//
	CopyTagsToVolumes interface{} `field:"optional" json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-dailyautomaticbackupstarttime
	//
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP, Amazon FSx for Windows File Server, or FSx for OpenZFS file system.
	//
	// By default, Amazon FSx automatically provisions 3 IOPS per GB of storage capacity. You can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how it is was provisioned, or the mode (by the customer or by Amazon FSx).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-diskiopsconfiguration
	//
	DiskIopsConfiguration interface{} `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
	//
	// By default in the Amazon FSx API and Amazon FSx console, Amazon FSx selects an available /28 IP address range for you from one of the VPC's CIDR ranges. You can have overlapping endpoint IP addresses for file systems deployed in the same VPC/route tables, as long as they don't overlap with any subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-endpointipaddressrange
	//
	EndpointIpAddressRange *string `field:"optional" json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// To delete a file system if there are child volumes present below the root volume, use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS` .
	//
	// If your file system has child volumes and you don't use this option, the delete request will fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-options
	//
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-preferredsubnetid
	//
	PreferredSubnetId *string `field:"optional" json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Specifies the optional provisioned SSD read cache on file systems that use the Intelligent-Tiering storage class.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-readcacheconfiguration
	//
	ReadCacheConfiguration interface{} `field:"optional" json:"readCacheConfiguration" yaml:"readCacheConfiguration"`
	// The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS file system.
	//
	// All volumes are children of the root volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-rootvolumeconfiguration
	//
	RootVolumeConfiguration interface{} `field:"optional" json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// (Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules for routing traffic to the correct file server.
	//
	// You should specify all virtual private cloud (VPC) route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-routetableids
	//
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
	// Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MBps).
	//
	// Valid values depend on the `DeploymentType` and `StorageType` that you choose, as follows:
	//
	// - For `INTELIGENT_TIERING` , valid values are 1280, 2560, 3840, 5120, 7680, or 10240 MBps.
	// - For `MULTI_AZ_1` and `SINGLE_AZ_2` , valid values are 160, 320, 640, 1280, 2560, 3840, 5120, 7680, or 10240 MBps.
	// - For `SINGLE_AZ_1` , valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MBps.
	//
	// You pay for additional throughput capacity that you provision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-throughputcapacity
	//
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-weeklymaintenancestarttime
	//
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

