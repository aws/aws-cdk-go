package awsfsx


// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lustreConfigurationProperty := &lustreConfigurationProperty{
//   	autoImportPolicy: jsii.String("autoImportPolicy"),
//   	automaticBackupRetentionDays: jsii.Number(123),
//   	copyTagsToBackups: jsii.Boolean(false),
//   	dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	dataCompressionType: jsii.String("dataCompressionType"),
//   	deploymentType: jsii.String("deploymentType"),
//   	driveCacheType: jsii.String("driveCacheType"),
//   	exportPath: jsii.String("exportPath"),
//   	importedFileChunkSize: jsii.Number(123),
//   	importPath: jsii.String("importPath"),
//   	perUnitStorageThroughput: jsii.Number(123),
//   	weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_LustreConfigurationProperty struct {
	// (Optional) Available with `Scratch` and `Persistent_1` deployment types.
	//
	// When you create your file system, your existing S3 objects appear as file and directory listings. Use this property to choose how Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. `AutoImportPolicy` can have the following values:
	//
	// - `NONE` - (Default) AutoImport is off. Amazon FSx only updates file and directory listings from the linked S3 bucket when the file system is created. FSx does not update file and directory listings for any new or changed objects after choosing this option.
	// - `NEW` - AutoImport is on. Amazon FSx automatically imports directory listings of any new objects added to the linked S3 bucket that do not currently exist in the FSx file system.
	// - `NEW_CHANGED` - AutoImport is on. Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket and any existing objects that are changed in the S3 bucket after you choose this option.
	// - `NEW_CHANGED_DELETED` - AutoImport is on. Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket, any existing objects that are changed in the S3 bucket, and any objects that were deleted in the S3 bucket.
	//
	// For more information, see [Automatically import updates from your S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) .
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	AutoImportPolicy *string `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean flag indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value. Only valid for use with `PERSISTENT_1` deployment types.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Sets the data compression configuration for the file system. `DataCompressionType` can have the following values:.
	//
	// - `NONE` - (Default) Data compression is turned off when the file system is created.
	// - `LZ4` - Data compression is turned on with the LZ4 algorithm.
	//
	// For more information, see [Lustre data compression](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html) in the *Amazon FSx for Lustre User Guide* .
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// (Optional) Choose `SCRATCH_1` and `SCRATCH_2` deployment types when you need temporary storage and shorter-term processing of data.
	//
	// The `SCRATCH_2` deployment type provides in-transit encryption of data and higher burst throughput capacity than `SCRATCH_1` .
	//
	// Choose `PERSISTENT_1` for longer-term storage and for throughput-focused workloads that aren’t latency-sensitive. `PERSISTENT_1` supports encryption of data in transit, and is available in all AWS Regions in which FSx for Lustre is available.
	//
	// Choose `PERSISTENT_2` for longer-term storage and for latency-sensitive workloads that require the highest levels of IOPS/throughput. `PERSISTENT_2` supports SSD storage, and offers higher `PerUnitStorageThroughput` (up to 1000 MB/s/TiB). `PERSISTENT_2` is available in a limited number of AWS Regions . For more information, and an up-to-date list of AWS Regions in which `PERSISTENT_2` is available, see [File system deployment options for FSx for Lustre](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-fsx-lustre.html#lustre-deployment-types) in the *Amazon FSx for Lustre User Guide* .
	//
	// > If you choose `PERSISTENT_2` , and you set `FileSystemTypeVersion` to `2.10` , the `CreateFileSystem` operation fails.
	//
	// Encryption of data in transit is automatically turned on when you access `SCRATCH_2` , `PERSISTENT_1` and `PERSISTENT_2` file systems from Amazon EC2 instances that [support automatic encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/data-                 protection.html) in the AWS Regions where they are available. For more information about encryption in transit for FSx for Lustre file systems, see [Encrypting data in transit](https://docs.aws.amazon.com/fsx/latest/LustreGuide/encryption-in-transit-fsxl.html) in the *Amazon FSx for Lustre User Guide* .
	//
	// (Default = `SCRATCH_1` ).
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// The type of drive cache used by `PERSISTENT_1` file systems that are provisioned with HDD storage devices.
	//
	// This parameter is required when storage type is HDD. Set this property to `READ` to improve the performance for frequently accessed files by caching up to 20% of the total storage capacity of the file system.
	//
	// This parameter is required when `StorageType` is set to `HDD` and `DeploymentType` is `PERSISTENT_1` .
	DriveCacheType *string `field:"optional" json:"driveCacheType" yaml:"driveCacheType"`
	// (Optional) Available with `Scratch` and `Persistent_1` deployment types.
	//
	// Specifies the path in the Amazon S3 bucket where the root of your Amazon FSx file system is exported. The path must use the same Amazon S3 bucket as specified in ImportPath. You can provide an optional prefix to which new and changed data is to be exported from your Amazon FSx for Lustre file system. If an `ExportPath` value is not provided, Amazon FSx sets a default export path, `s3://import-bucket/FSxLustre[creation-timestamp]` . The timestamp is in UTC format, for example `s3://import-bucket/FSxLustre20181105T222312Z` .
	//
	// The Amazon S3 export bucket must be the same as the import bucket specified by `ImportPath` . If you specify only a bucket name, such as `s3://import-bucket` , you get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is overwritten on export. If you provide a custom prefix in the export path, such as `s3://import-bucket/[custom-optional-prefix]` , Amazon FSx exports the contents of your file system to that export prefix in the Amazon S3 bucket.
	//
	// > This parameter is not supported for file systems using the `Persistent_2` deployment type.
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// (Optional) For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// (Optional) The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// The root of your FSx for Lustre file system will be mapped to the root of the Amazon S3 bucket you select. An example is `s3://import-bucket/optional-prefix` . If you specify a prefix after the Amazon S3 bucket name, only object keys with that prefix are loaded into the file system.
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// Required with `PERSISTENT_1` and `PERSISTENT_2` deployment types, provisions the amount of read and write throughput for each 1 tebibyte (TiB) of file system storage capacity, in MB/s/TiB.
	//
	// File system throughput capacity is calculated by multiplying ﬁle system storage capacity (TiB) by the `PerUnitStorageThroughput` (MB/s/TiB). For a 2.4-TiB ﬁle system, provisioning 50 MB/s/TiB of `PerUnitStorageThroughput` yields 120 MB/s of ﬁle system throughput. You pay for the amount of throughput that you provision.
	//
	// Valid values:
	//
	// - For `PERSISTENT_1` SSD storage: 50, 100, 200 MB/s/TiB.
	// - For `PERSISTENT_1` HDD storage: 12, 40 MB/s/TiB.
	// - For `PERSISTENT_2` SSD storage: 125, 250, 500, 1000 MB/s/TiB.
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

