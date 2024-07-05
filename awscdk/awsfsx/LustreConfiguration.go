package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var bucket bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	Vpc: vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageCapacityGiB: jsii.Number(1200),
//   	LustreConfiguration: LustreConfiguration,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
type LustreConfiguration struct {
	// The type of backing file system deployment used by FSx.
	DeploymentType LustreDeploymentType `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// Available with `Scratch` and `Persistent_1` deployment types.
	//
	// When you create your file system, your existing S3 objects appear as file and directory listings. Use this property to choose how Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. `AutoImportPolicy` can have the following values:
	//
	// For more information, see [Automatically import updates from your S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/autoimport-data-repo.html) .
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	// Default: - no import policy.
	//
	AutoImportPolicy LustreAutoImportPolicy `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to 0 disables automatic backups.
	// You can retain automatic backups for a maximum of 90 days.
	//
	// Automatic Backups is not supported on scratch file systems.
	// Default: Duration.days(0)
	//
	AutomaticBackupRetention awscdk.Duration `field:"optional" json:"automaticBackupRetention" yaml:"automaticBackupRetention"`
	// A boolean flag indicating whether tags for the file system should be copied to backups.
	// Default: - false.
	//
	CopyTagsToBackups *bool `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Start time for 30-minute daily automatic backup window in Coordinated Universal Time (UTC).
	// Default: - no backup window.
	//
	DailyAutomaticBackupStartTime DailyAutomaticBackupStartTime `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Sets the data compression configuration for the file system.
	//
	// For more information, see [Lustre data compression](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html) in the *Amazon FSx for Lustre User Guide* .
	// Default: - no compression.
	//
	DataCompressionType LustreDataCompressionType `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// The path in Amazon S3 where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same
	// Amazon S3 bucket as specified in ImportPath. If you only specify a bucket name, such as s3://import-bucket, you
	// get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is
	// overwritten on export. If you provide a custom prefix in the export path, such as
	// s3://import-bucket/[custom-optional-prefix], Amazon FSx exports the contents of your file system to that export
	// prefix in the Amazon S3 bucket.
	// Default: s3://import-bucket/FSxLustre[creation-timestamp].
	//
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// Allowed values are between 1 and 512,000.
	// Default: 1024.
	//
	ImportedFileChunkSizeMiB *float64 `field:"optional" json:"importedFileChunkSizeMiB" yaml:"importedFileChunkSizeMiB"`
	// The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// Must be of the format "s3://{bucketName}/optional-prefix" and cannot
	// exceed 900 characters.
	// Default: - no bucket is imported.
	//
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// Required for the PERSISTENT_1 deployment type, describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB.
	//
	// Valid values are 50, 100, 200.
	// Default: - no default, conditionally required for PERSISTENT_1 deployment type.
	//
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// The preferred day and time to perform weekly maintenance.
	//
	// The first digit is the day of the week, starting at 1
	// for Monday, then the following are hours and minutes in the UTC time zone, 24 hour clock. For example: '2:20:30'
	// is Tuesdays at 20:30.
	// Default: - no preference.
	//
	WeeklyMaintenanceStartTime LustreMaintenanceTime `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

