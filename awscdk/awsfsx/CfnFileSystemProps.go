package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFileSystem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFileSystemProps := &CfnFileSystemProps{
//   	FileSystemType: jsii.String("fileSystemType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	BackupId: jsii.String("backupId"),
//   	FileSystemTypeVersion: jsii.String("fileSystemTypeVersion"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LustreConfiguration: &LustreConfigurationProperty{
//   		AutoImportPolicy: jsii.String("autoImportPolicy"),
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DataCompressionType: jsii.String("dataCompressionType"),
//   		DeploymentType: jsii.String("deploymentType"),
//   		DriveCacheType: jsii.String("driveCacheType"),
//   		ExportPath: jsii.String("exportPath"),
//   		ImportedFileChunkSize: jsii.Number(123),
//   		ImportPath: jsii.String("importPath"),
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	OntapConfiguration: &OntapConfigurationProperty{
//   		DeploymentType: jsii.String("deploymentType"),
//
//   		// the properties below are optional
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		EndpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   		FsxAdminPassword: jsii.String("fsxAdminPassword"),
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		RouteTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	OpenZfsConfiguration: &OpenZFSConfigurationProperty{
//   		DeploymentType: jsii.String("deploymentType"),
//
//   		// the properties below are optional
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		CopyTagsToVolumes: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		Options: []*string{
//   			jsii.String("options"),
//   		},
//   		RootVolumeConfiguration: &RootVolumeConfigurationProperty{
//   			CopyTagsToSnapshots: jsii.Boolean(false),
//   			DataCompressionType: jsii.String("dataCompressionType"),
//   			NfsExports: []interface{}{
//   				&NfsExportsProperty{
//   					ClientConfigurations: []interface{}{
//   						&ClientConfigurationsProperty{
//   							Clients: jsii.String("clients"),
//   							Options: []*string{
//   								jsii.String("options"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			ReadOnly: jsii.Boolean(false),
//   			RecordSizeKiB: jsii.Number(123),
//   			UserAndGroupQuotas: []interface{}{
//   				&UserAndGroupQuotasProperty{
//   					Id: jsii.Number(123),
//   					StorageCapacityQuotaGiB: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WindowsConfiguration: &WindowsConfigurationProperty{
//   		ThroughputCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		ActiveDirectoryId: jsii.String("activeDirectoryId"),
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		AuditLogConfiguration: &AuditLogConfigurationProperty{
//   			FileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   			FileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//
//   			// the properties below are optional
//   			AuditLogDestination: jsii.String("auditLogDestination"),
//   		},
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DeploymentType: jsii.String("deploymentType"),
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		SelfManagedActiveDirectoryConfiguration: &SelfManagedActiveDirectoryConfigurationProperty{
//   			DnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			DomainName: jsii.String("domainName"),
//   			FileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   			Password: jsii.String("password"),
//   			UserName: jsii.String("userName"),
//   		},
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   }
//
type CfnFileSystemProps struct {
	// The type of Amazon FSx file system, which can be `LUSTRE` , `WINDOWS` , `ONTAP` , or `OPENZFS` .
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
	// Specifies the IDs of the subnets that the file system will be accessible from.
	//
	// For Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet IDs, one for the preferred file server and one for the standby file server. You specify one of these subnets as the preferred subnet using the `WindowsConfiguration > PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID` properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the *Amazon FSx for Windows User Guide* and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the *Amazon FSx for ONTAP User Guide* .
	//
	// For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre deployment types, provide exactly one subnet ID. The file server is launched in that subnet's Availability Zone.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the file system backup that you are using to create a file system.
	//
	// For more information, see [CreateFileSystemFromBackup](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemFromBackup.html) .
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// (Optional) For FSx for Lustre file systems, sets the Lustre version for the file system that you're creating.
	//
	// Valid values are `2.10` and `2.12` :
	//
	// - 2.10 is supported by the Scratch and Persistent_1 Lustre deployment types.
	// - 2.12 is supported by all Lustre deployment types. `2.12` is required when setting FSx for Lustre `DeploymentType` to `PERSISTENT_2` .
	//
	// Default value = `2.10` , except when `DeploymentType` is set to `PERSISTENT_2` , then the default is `2.12` .
	//
	// > If you set `FileSystemTypeVersion` to `2.10` for a `PERSISTENT_2` Lustre deployment type, the `CreateFileSystem` operation fails.
	FileSystemTypeVersion *string `field:"optional" json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
	// The ID of the AWS Key Management Service ( AWS KMS ) key used to encrypt Amazon FSx file system data.
	//
	// Used as follows with Amazon FSx file system types:
	//
	// - Amazon FSx for Lustre `PERSISTENT_1` and `PERSISTENT_2` deployment types only.
	//
	// `SCRATCH_1` and `SCRATCH_2` types are encrypted using the Amazon FSx service AWS KMS key for your account.
	// - Amazon FSx for NetApp ONTAP
	// - Amazon FSx for OpenZFS
	// - Amazon FSx for Windows File Server.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Lustre configuration for the file system being created.
	//
	// > The following parameters are not supported for file systems with a data repository association.
	// >
	// > - `AutoImportPolicy`
	// > - `ExportPath`
	// > - `ImportedChunkSize`
	// > - `ImportPath`.
	LustreConfiguration interface{} `field:"optional" json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The ONTAP configuration properties of the FSx for ONTAP file system that you are creating.
	OntapConfiguration interface{} `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The Amazon FSx for OpenZFS configuration properties for the file system that you are creating.
	OpenZfsConfiguration interface{} `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Sets the storage capacity of the file system that you're creating.
	//
	// `StorageCapacity` is required if you are creating a new file system.
	//
	// *FSx for Lustre file systems* - The amount of storage capacity that you can configure depends on the value that you set for `StorageType` and the Lustre `DeploymentType` , as follows:
	//
	// - For `SCRATCH_2` , `PERSISTENT_2` and `PERSISTENT_1` deployment types using SSD storage type, the valid values are 1200 GiB, 2400 GiB, and increments of 2400 GiB.
	// - For `PERSISTENT_1` HDD file systems, valid values are increments of 6000 GiB for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB file systems.
	// - For `SCRATCH_1` deployment type, valid values are 1200 GiB, 2400 GiB, and increments of 3600 GiB.
	//
	// *FSx for ONTAP file systems* - The amount of storage capacity that you can configure is from 1024 GiB up to 196,608 GiB (192 TiB).
	//
	// *FSx for OpenZFS file systems* - The amount of storage capacity that you can configure is from 64 GiB up to 524,288 GiB (512 TiB). If you are creating a file system from a backup, you can specify a storage capacity equal to or greater than the original file system's storage capacity.
	//
	// *FSx for Windows File Server file systems* - The amount of storage capacity that you can configure depends on the value that you set for `StorageType` as follows:
	//
	// - For SSD storage, valid values are 32 GiB-65,536 GiB (64 TiB).
	// - For HDD storage, valid values are 2000 GiB-65,536 GiB (64 TiB).
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Sets the storage type for the file system that you're creating. Valid values are `SSD` and `HDD` .
	//
	// - Set to `SSD` to use solid state drive storage. SSD is supported on all Windows, Lustre, ONTAP, and OpenZFS deployment types.
	// - Set to `HDD` to use hard disk drive storage. HDD is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types, and on `PERSISTENT_1` Lustre file system deployment types.
	//
	// Default value is `SSD` . For more information, see [Storage type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the *FSx for Windows File Server User Guide* and [Multiple storage options](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html#storage-options) in the *FSx for Lustre User Guide* .
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This value is required if `FileSystemType` is set to `WINDOWS` .
	WindowsConfiguration interface{} `field:"optional" json:"windowsConfiguration" yaml:"windowsConfiguration"`
}

