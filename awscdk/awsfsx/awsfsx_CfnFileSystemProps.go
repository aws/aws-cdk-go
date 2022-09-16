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
//   cfnFileSystemProps := &cfnFileSystemProps{
//   	fileSystemType: jsii.String("fileSystemType"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	backupId: jsii.String("backupId"),
//   	fileSystemTypeVersion: jsii.String("fileSystemTypeVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	lustreConfiguration: &lustreConfigurationProperty{
//   		autoImportPolicy: jsii.String("autoImportPolicy"),
//   		automaticBackupRetentionDays: jsii.Number(123),
//   		copyTagsToBackups: jsii.Boolean(false),
//   		dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		dataCompressionType: jsii.String("dataCompressionType"),
//   		deploymentType: jsii.String("deploymentType"),
//   		driveCacheType: jsii.String("driveCacheType"),
//   		exportPath: jsii.String("exportPath"),
//   		importedFileChunkSize: jsii.Number(123),
//   		importPath: jsii.String("importPath"),
//   		perUnitStorageThroughput: jsii.Number(123),
//   		weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	ontapConfiguration: &ontapConfigurationProperty{
//   		deploymentType: jsii.String("deploymentType"),
//
//   		// the properties below are optional
//   		automaticBackupRetentionDays: jsii.Number(123),
//   		dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		diskIopsConfiguration: &diskIopsConfigurationProperty{
//   			iops: jsii.Number(123),
//   			mode: jsii.String("mode"),
//   		},
//   		endpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   		fsxAdminPassword: jsii.String("fsxAdminPassword"),
//   		preferredSubnetId: jsii.String("preferredSubnetId"),
//   		routeTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		throughputCapacity: jsii.Number(123),
//   		weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	openZfsConfiguration: &openZFSConfigurationProperty{
//   		deploymentType: jsii.String("deploymentType"),
//
//   		// the properties below are optional
//   		automaticBackupRetentionDays: jsii.Number(123),
//   		copyTagsToBackups: jsii.Boolean(false),
//   		copyTagsToVolumes: jsii.Boolean(false),
//   		dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		diskIopsConfiguration: &diskIopsConfigurationProperty{
//   			iops: jsii.Number(123),
//   			mode: jsii.String("mode"),
//   		},
//   		options: []*string{
//   			jsii.String("options"),
//   		},
//   		rootVolumeConfiguration: &rootVolumeConfigurationProperty{
//   			copyTagsToSnapshots: jsii.Boolean(false),
//   			dataCompressionType: jsii.String("dataCompressionType"),
//   			nfsExports: []interface{}{
//   				&nfsExportsProperty{
//   					clientConfigurations: []interface{}{
//   						&clientConfigurationsProperty{
//   							clients: jsii.String("clients"),
//   							options: []*string{
//   								jsii.String("options"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			readOnly: jsii.Boolean(false),
//   			recordSizeKiB: jsii.Number(123),
//   			userAndGroupQuotas: []interface{}{
//   				&userAndGroupQuotasProperty{
//   					id: jsii.Number(123),
//   					storageCapacityQuotaGiB: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		throughputCapacity: jsii.Number(123),
//   		weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	storageCapacity: jsii.Number(123),
//   	storageType: jsii.String("storageType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	windowsConfiguration: &windowsConfigurationProperty{
//   		throughputCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		activeDirectoryId: jsii.String("activeDirectoryId"),
//   		aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		auditLogConfiguration: &auditLogConfigurationProperty{
//   			fileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   			fileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//
//   			// the properties below are optional
//   			auditLogDestination: jsii.String("auditLogDestination"),
//   		},
//   		automaticBackupRetentionDays: jsii.Number(123),
//   		copyTagsToBackups: jsii.Boolean(false),
//   		dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		deploymentType: jsii.String("deploymentType"),
//   		preferredSubnetId: jsii.String("preferredSubnetId"),
//   		selfManagedActiveDirectoryConfiguration: &selfManagedActiveDirectoryConfigurationProperty{
//   			dnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			domainName: jsii.String("domainName"),
//   			fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   			organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   			password: jsii.String("password"),
//   			userName: jsii.String("userName"),
//   		},
//   		weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
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
	// > The following parameters are not supported for file systems with the Lustre `Persistent_2` deployment type.
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
	// `StorageCapacity` is required if you are creating a new file system. Do not include `StorageCapacity` if you are creating a file system from a backup.
	//
	// *FSx for Lustre file systems* - The amount of storage capacity that you can configure depends on the value that you set for `StorageType` and the Lustre `DeploymentType` , as follows:
	//
	// - For `SCRATCH_2` , `PERSISTENT_2` and `PERSISTENT_1` deployment types using SSD storage type, the valid values are 1200 GiB, 2400 GiB, and increments of 2400 GiB.
	// - For `PERSISTENT_1` HDD file systems, valid values are increments of 6000 GiB for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB file systems.
	// - For `SCRATCH_1` deployment type, valid values are 1200 GiB, 2400 GiB, and increments of 3600 GiB.
	//
	// *FSx for ONTAP file systems* - The amount of storage capacity that you can configure is from 1024 GiB up to 196,608 GiB (192 TiB).
	//
	// *FSx for OpenZFS file systems* - The amount of storage capacity that you can configure is from 64 GiB up to 524,288 GiB (512 TiB).
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

