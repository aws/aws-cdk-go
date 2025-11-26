package previewawsfsxmixins


// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lustreConfigurationProperty := &LustreConfigurationProperty{
//   	AutoImportPolicy: jsii.String("autoImportPolicy"),
//   	AutomaticBackupRetentionDays: jsii.Number(123),
//   	CopyTagsToBackups: jsii.Boolean(false),
//   	DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	DataCompressionType: jsii.String("dataCompressionType"),
//   	DataReadCacheConfiguration: &DataReadCacheConfigurationProperty{
//   		SizeGiB: jsii.Number(123),
//   		SizingMode: jsii.String("sizingMode"),
//   	},
//   	DeploymentType: jsii.String("deploymentType"),
//   	DriveCacheType: jsii.String("driveCacheType"),
//   	EfaEnabled: jsii.Boolean(false),
//   	ExportPath: jsii.String("exportPath"),
//   	ImportedFileChunkSize: jsii.Number(123),
//   	ImportPath: jsii.String("importPath"),
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		Iops: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   	},
//   	PerUnitStorageThroughput: jsii.Number(123),
//   	ThroughputCapacity: jsii.Number(123),
//   	WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
type CfnFileSystemPropsMixin_LustreConfigurationProperty struct {
	// (Optional) When you create your file system, your existing S3 objects appear as file and directory listings.
	//
	// Use this property to choose how Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. `AutoImportPolicy` can have the following values:
	//
	// - `NONE` - (Default) AutoImport is off. Amazon FSx only updates file and directory listings from the linked S3 bucket when the file system is created. FSx does not update file and directory listings for any new or changed objects after choosing this option.
	// - `NEW` - AutoImport is on. Amazon FSx automatically imports directory listings of any new objects added to the linked S3 bucket that do not currently exist in the FSx file system.
	// - `NEW_CHANGED` - AutoImport is on. Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket and any existing objects that are changed in the S3 bucket after you choose this option.
	// - `NEW_CHANGED_DELETED` - AutoImport is on. Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket, any existing objects that are changed in the S3 bucket, and any objects that were deleted in the S3 bucket.
	//
	// For more information, see [Automatically import updates from your S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) .
	//
	// > This parameter is not supported for Lustre file systems with a data repository association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-autoimportpolicy
	//
	AutoImportPolicy *string `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-automaticbackupretentiondays
	//
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// (Optional) Not available for use with file systems that are linked to a data repository.
	//
	// A boolean flag indicating whether tags for the file system should be copied to backups. The default value is false. If `CopyTagsToBackups` is set to true, all file system tags are copied to all automatic and user-initiated backups when the user doesn't specify any backup-specific tags. If `CopyTagsToBackups` is set to true and you specify one or more backup tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	//
	// (Default = `false` )
	//
	// For more information, see [Working with backups](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html) in the *Amazon FSx for Lustre User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-copytagstobackups
	//
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-dailyautomaticbackupstarttime
	//
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Sets the data compression configuration for the file system. `DataCompressionType` can have the following values:.
	//
	// - `NONE` - (Default) Data compression is turned off when the file system is created.
	// - `LZ4` - Data compression is turned on with the LZ4 algorithm.
	//
	// For more information, see [Lustre data compression](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html) in the *Amazon FSx for Lustre User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-datacompressiontype
	//
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// Specifies the optional provisioned SSD read cache on FSx for Lustre file systems that use the Intelligent-Tiering storage class.
	//
	// Required when `StorageType` is set to `INTELLIGENT_TIERING` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-datareadcacheconfiguration
	//
	DataReadCacheConfiguration interface{} `field:"optional" json:"dataReadCacheConfiguration" yaml:"dataReadCacheConfiguration"`
	// (Optional) Choose `SCRATCH_1` and `SCRATCH_2` deployment types when you need temporary storage and shorter-term processing of data.
	//
	// The `SCRATCH_2` deployment type provides in-transit encryption of data and higher burst throughput capacity than `SCRATCH_1` .
	//
	// Choose `PERSISTENT_1` for longer-term storage and for throughput-focused workloads that aren’t latency-sensitive. `PERSISTENT_1` supports encryption of data in transit, and is available in all AWS Regions in which FSx for Lustre is available.
	//
	// Choose `PERSISTENT_2` for longer-term storage and for latency-sensitive workloads that require the highest levels of IOPS/throughput. `PERSISTENT_2` supports the SSD and Intelligent-Tiering storage classes. You can optionally specify a metadata configuration mode for `PERSISTENT_2` which supports increasing metadata performance. `PERSISTENT_2` is available in a limited number of AWS Regions . For more information, and an up-to-date list of AWS Regions in which `PERSISTENT_2` is available, see [Deployment and storage class options for FSx for Lustre file systems](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-fsx-lustre.html) in the *Amazon FSx for Lustre User Guide* .
	//
	// > If you choose `PERSISTENT_2` , and you set `FileSystemTypeVersion` to `2.10` , the `CreateFileSystem` operation fails.
	//
	// Encryption of data in transit is automatically turned on when you access `SCRATCH_2` , `PERSISTENT_1` , and `PERSISTENT_2` file systems from Amazon EC2 instances that support automatic encryption in the AWS Regions where they are available. For more information about encryption in transit for FSx for Lustre file systems, see [Encrypting data in transit](https://docs.aws.amazon.com/fsx/latest/LustreGuide/encryption-in-transit-fsxl.html) in the *Amazon FSx for Lustre User Guide* .
	//
	// (Default = `SCRATCH_1` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-deploymenttype
	//
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// The type of drive cache used by `PERSISTENT_1` file systems that are provisioned with HDD storage devices.
	//
	// This parameter is required when storage type is HDD. Set this property to `READ` to improve the performance for frequently accessed files by caching up to 20% of the total storage capacity of the file system.
	//
	// This parameter is required when `StorageType` is set to `HDD` and `DeploymentType` is `PERSISTENT_1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-drivecachetype
	//
	DriveCacheType *string `field:"optional" json:"driveCacheType" yaml:"driveCacheType"`
	// (Optional) Specifies whether Elastic Fabric Adapter (EFA) and GPUDirect Storage (GDS) support is enabled for the Amazon FSx for Lustre file system.
	//
	// (Default = `false` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-efaenabled
	//
	EfaEnabled interface{} `field:"optional" json:"efaEnabled" yaml:"efaEnabled"`
	// (Optional) Specifies the path in the Amazon S3 bucket where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same Amazon S3 bucket as specified in ImportPath. You can provide an optional prefix to which new and changed data is to be exported from your Amazon FSx for Lustre file system. If an `ExportPath` value is not provided, Amazon FSx sets a default export path, `s3://import-bucket/FSxLustre[creation-timestamp]` . The timestamp is in UTC format, for example `s3://import-bucket/FSxLustre20181105T222312Z` .
	//
	// The Amazon S3 export bucket must be the same as the import bucket specified by `ImportPath` . If you specify only a bucket name, such as `s3://import-bucket` , you get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is overwritten on export. If you provide a custom prefix in the export path, such as `s3://import-bucket/[custom-optional-prefix]` , Amazon FSx exports the contents of your file system to that export prefix in the Amazon S3 bucket.
	//
	// > This parameter is not supported for file systems with a data repository association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-exportpath
	//
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// (Optional) For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	//
	// > This parameter is not supported for Lustre file systems with a data repository association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-importedfilechunksize
	//
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// (Optional) The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// The root of your FSx for Lustre file system will be mapped to the root of the Amazon S3 bucket you select. An example is `s3://import-bucket/optional-prefix` . If you specify a prefix after the Amazon S3 bucket name, only object keys with that prefix are loaded into the file system.
	//
	// > This parameter is not supported for Lustre file systems with a data repository association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-importpath
	//
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// The Lustre metadata performance configuration for the creation of an FSx for Lustre file system using a `PERSISTENT_2` deployment type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// Required with `PERSISTENT_1` and `PERSISTENT_2` deployment types, provisions the amount of read and write throughput for each 1 tebibyte (TiB) of file system storage capacity, in MB/s/TiB.
	//
	// File system throughput capacity is calculated by multiplying ﬁle system storage capacity (TiB) by the `PerUnitStorageThroughput` (MB/s/TiB). For a 2.4-TiB ﬁle system, provisioning 50 MB/s/TiB of `PerUnitStorageThroughput` yields 120 MB/s of ﬁle system throughput. You pay for the amount of throughput that you provision.
	//
	// Valid values:
	//
	// - For `PERSISTENT_1` SSD storage: 50, 100, 200 MB/s/TiB.
	// - For `PERSISTENT_1` HDD storage: 12, 40 MB/s/TiB.
	// - For `PERSISTENT_2` SSD storage: 125, 250, 500, 1000 MB/s/TiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-perunitstoragethroughput
	//
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// Specifies the throughput of an FSx for Lustre file system using the Intelligent-Tiering storage class, measured in megabytes per second (MBps).
	//
	// Valid values are 4000 MBps or multiples of 4000 MBps. You pay for the amount of throughput that you provision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-throughputcapacity
	//
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone, where d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html#cfn-fsx-filesystem-lustreconfiguration-weeklymaintenancestarttime
	//
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

