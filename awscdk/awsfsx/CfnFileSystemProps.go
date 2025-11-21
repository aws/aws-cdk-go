package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   		DataReadCacheConfiguration: &DataReadCacheConfigurationProperty{
//   			SizeGiB: jsii.Number(123),
//   			SizingMode: jsii.String("sizingMode"),
//   		},
//   		DeploymentType: jsii.String("deploymentType"),
//   		DriveCacheType: jsii.String("driveCacheType"),
//   		EfaEnabled: jsii.Boolean(false),
//   		ExportPath: jsii.String("exportPath"),
//   		ImportedFileChunkSize: jsii.Number(123),
//   		ImportPath: jsii.String("importPath"),
//   		MetadataConfiguration: &MetadataConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	NetworkType: jsii.String("networkType"),
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
//   		EndpointIpv6AddressRange: jsii.String("endpointIpv6AddressRange"),
//   		FsxAdminPassword: jsii.String("fsxAdminPassword"),
//   		HaPairs: jsii.Number(123),
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		RouteTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		ThroughputCapacityPerHaPair: jsii.Number(123),
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
//   		EndpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   		EndpointIpv6AddressRange: jsii.String("endpointIpv6AddressRange"),
//   		Options: []*string{
//   			jsii.String("options"),
//   		},
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		ReadCacheConfiguration: &ReadCacheConfigurationProperty{
//   			SizeGiB: jsii.Number(123),
//   			SizingMode: jsii.String("sizingMode"),
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
//   		RouteTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
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
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		SelfManagedActiveDirectoryConfiguration: &SelfManagedActiveDirectoryConfigurationProperty{
//   			DnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			DomainJoinServiceAccountSecret: jsii.String("domainJoinServiceAccountSecret"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
type CfnFileSystemProps struct {
	// The type of Amazon FSx file system, which can be `LUSTRE` , `WINDOWS` , `ONTAP` , or `OPENZFS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-filesystemtype
	//
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
	// Specifies the IDs of the subnets that the file system will be accessible from.
	//
	// For Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet IDs, one for the preferred file server and one for the standby file server. You specify one of these subnets as the preferred subnet using the `WindowsConfiguration > PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID` properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the *Amazon FSx for Windows User Guide* and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the *Amazon FSx for ONTAP User Guide* .
	//
	// For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre deployment types, provide exactly one subnet ID. The file server is launched in that subnet's Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the file system backup that you are using to create a file system.
	//
	// For more information, see [CreateFileSystemFromBackup](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemFromBackup.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-backupid
	//
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// For FSx for Lustre file systems, sets the Lustre version for the file system that you're creating.
	//
	// Valid values are `2.10` , `2.12` , and `2.15` :
	//
	// - `2.10` is supported by the Scratch and Persistent_1 Lustre deployment types.
	// - `2.12` is supported by all Lustre deployment types, except for `PERSISTENT_2` with a metadata configuration mode.
	// - `2.15` is supported by all Lustre deployment types and is recommended for all new file systems.
	//
	// Default value is `2.10` , except for the following deployments:
	//
	// - Default value is `2.12` when `DeploymentType` is set to `PERSISTENT_2` without a metadata configuration mode.
	// - Default value is `2.15` when `DeploymentType` is set to `PERSISTENT_2` with a metadata configuration mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-filesystemtypeversion
	//
	FileSystemTypeVersion *string `field:"optional" json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
	// The ID of the AWS Key Management Service ( AWS  ) key used to encrypt Amazon FSx file system data.
	//
	// Used as follows with Amazon FSx file system types:
	//
	// - Amazon FSx for Lustre `PERSISTENT_1` and `PERSISTENT_2` deployment types only.
	//
	// `SCRATCH_1` and `SCRATCH_2` types are encrypted using the Amazon FSx service AWS  key for your account.
	// - Amazon FSx for NetApp ONTAP
	// - Amazon FSx for OpenZFS
	// - Amazon FSx for Windows File Server
	//
	// If this ID isn't specified, the Amazon FSx-managed key for your account is used. For more information, see [Encrypt](https://docs.aws.amazon.com//kms/latest/APIReference/API_Encrypt.html) in the *AWS Key Management Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Lustre configuration for the file system being created.
	//
	// This configuration is required if the `FileSystemType` is set to `LUSTRE` .
	//
	// > The following parameters are not supported when creating Lustre file systems with a data repository association.
	// >
	// > - `AutoImportPolicy`
	// > - `ExportPath`
	// > - `ImportedChunkSize`
	// > - `ImportPath`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-lustreconfiguration
	//
	LustreConfiguration interface{} `field:"optional" json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The network type of the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The ONTAP configuration properties of the FSx for ONTAP file system that you are creating.
	//
	// This configuration is required if the `FileSystemType` is set to `ONTAP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-ontapconfiguration
	//
	OntapConfiguration interface{} `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The Amazon FSx for OpenZFS configuration properties for the file system that you are creating.
	//
	// This configuration is required if the `FileSystemType` is set to `OPENZFS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-openzfsconfiguration
	//
	OpenZfsConfiguration interface{} `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	//
	// > You must specify a security group if you are creating a Multi-AZ FSx for ONTAP file system in a VPC subnet that has been shared with you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Sets the storage capacity of the file system that you're creating.
	//
	// `StorageCapacity` is required if you are creating a new file system. It is not required if you are creating a file system by restoring a backup.
	//
	// *FSx for Lustre file systems* - The amount of storage capacity that you can configure depends on the value that you set for `StorageType` and the Lustre `DeploymentType` , as follows:
	//
	// - For `SCRATCH_2` , `PERSISTENT_2` and `PERSISTENT_1` deployment types using SSD storage type, the valid values are 1200 GiB, 2400 GiB, and increments of 2400 GiB.
	// - For `PERSISTENT_1` HDD file systems, valid values are increments of 6000 GiB for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB file systems.
	// - For `SCRATCH_1` deployment type, valid values are 1200 GiB, 2400 GiB, and increments of 3600 GiB.
	//
	// *FSx for ONTAP file systems* - The amount of SSD storage capacity that you can configure depends on the value of the `HAPairs` property. The minimum value is calculated as 1,024 GiB * HAPairs and the maximum is calculated as 524,288 GiB * HAPairs, up to a maximum amount of SSD storage capacity of 1,048,576 GiB (1 pebibyte).
	//
	// *FSx for OpenZFS file systems* - The amount of storage capacity that you can configure is from 64 GiB up to 524,288 GiB (512 TiB). If you are creating a file system from a backup, you can specify a storage capacity equal to or greater than the original file system's storage capacity.
	//
	// *FSx for Windows File Server file systems* - The amount of storage capacity that you can configure depends on the value that you set for `StorageType` as follows:
	//
	// - For SSD storage, valid values are 32 GiB-65,536 GiB (64 TiB).
	// - For HDD storage, valid values are 2000 GiB-65,536 GiB (64 TiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-storagecapacity
	//
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Sets the storage class for the file system that you're creating.
	//
	// Valid values are `SSD` , `HDD` , and `INTELLIGENT_TIERING` .
	//
	// - Set to `SSD` to use solid state drive storage. SSD is supported on all Windows, Lustre, ONTAP, and OpenZFS deployment types.
	// - Set to `HDD` to use hard disk drive storage, which is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types, and on `PERSISTENT_1` Lustre file system deployment types.
	// - Set to `INTELLIGENT_TIERING` to use fully elastic, intelligently-tiered storage. Intelligent-Tiering is only available for OpenZFS file systems with the Multi-AZ deployment type and for Lustre file systems with the Persistent_2 deployment type.
	//
	// Default value is `SSD` . For more information, see [Storage type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the *FSx for Windows File Server User Guide* , [FSx for Lustre storage classes](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-fsx-lustre.html#lustre-storage-classes) in the *FSx for Lustre User Guide* , and [Working with Intelligent-Tiering](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance-intelligent-tiering) in the *Amazon FSx for OpenZFS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The tags to associate with the file system.
	//
	// For more information, see [Tagging your Amazon FSx resources](https://docs.aws.amazon.com/fsx/latest/LustreGuide/tag-resources.html) in the *Amazon FSx for Lustre User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This configuration is required if `FileSystemType` is set to `WINDOWS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-windowsconfiguration
	//
	WindowsConfiguration interface{} `field:"optional" json:"windowsConfiguration" yaml:"windowsConfiguration"`
}

