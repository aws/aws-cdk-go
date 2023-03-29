package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::FSx::FileSystem`.
//
// The `AWS::FSx::FileSystem` resource is an Amazon FSx resource type that specifies an Amazon FSx file system. You can create any of the following supported file system types:
//
// - Amazon FSx for Lustre
// - Amazon FSx for NetApp ONTAP
// - Amazon FSx for OpenZFS
// - Amazon FSx for Windows File Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFileSystem := awscdk.Aws_fsx.NewCfnFileSystem(this, jsii.String("MyCfnFileSystem"), &CfnFileSystemProps{
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
//   })
//
type CfnFileSystem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the FSx for Windows file system's DNSName.
	//
	// Example: `amznfsxp1honlek.corp.example.com`
	AttrDnsName() *string
	// Returns the Lustre file system's `LustreMountName` .
	//
	// Example for `SCRATCH_1` deployment types: This value is always `fsx` .
	//
	// Example for `SCRATCH_2` and `PERSISTENT` deployment types: `2p3fhbmv`.
	AttrLustreMountName() *string
	// Returns the Amazon Resource Name (ARN) for the Amazon FSx file system.
	//
	// Example: `arn:aws:fsx:us-east-2:111122223333:file-system/fs-0123abcd56789ef0a`.
	AttrResourceArn() *string
	// Returns the root volume ID of the FSx for OpenZFS file system.
	//
	// Example: `fsvol-0123456789abcdefa`.
	AttrRootVolumeId() *string
	// The ID of the file system backup that you are using to create a file system.
	//
	// For more information, see [CreateFileSystemFromBackup](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemFromBackup.html) .
	BackupId() *string
	SetBackupId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The type of Amazon FSx file system, which can be `LUSTRE` , `WINDOWS` , `ONTAP` , or `OPENZFS` .
	FileSystemType() *string
	SetFileSystemType(val *string)
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
	FileSystemTypeVersion() *string
	SetFileSystemTypeVersion(val *string)
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
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The Lustre configuration for the file system being created.
	//
	// > The following parameters are not supported for file systems with the Lustre `Persistent_2` deployment type.
	// >
	// > - `AutoImportPolicy`
	// > - `ExportPath`
	// > - `ImportedChunkSize`
	// > - `ImportPath`.
	LustreConfiguration() interface{}
	SetLustreConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// The ONTAP configuration properties of the FSx for ONTAP file system that you are creating.
	OntapConfiguration() interface{}
	SetOntapConfiguration(val interface{})
	// The Amazon FSx for OpenZFS configuration properties for the file system that you are creating.
	OpenZfsConfiguration() interface{}
	SetOpenZfsConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	// Sets the storage type for the file system that you're creating. Valid values are `SSD` and `HDD` .
	//
	// - Set to `SSD` to use solid state drive storage. SSD is supported on all Windows, Lustre, ONTAP, and OpenZFS deployment types.
	// - Set to `HDD` to use hard disk drive storage. HDD is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types, and on `PERSISTENT_1` Lustre file system deployment types.
	//
	// Default value is `SSD` . For more information, see [Storage type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the *FSx for Windows File Server User Guide* and [Multiple storage options](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html#storage-options) in the *FSx for Lustre User Guide* .
	StorageType() *string
	SetStorageType(val *string)
	// Specifies the IDs of the subnets that the file system will be accessible from.
	//
	// For Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet IDs, one for the preferred file server and one for the standby file server. You specify one of these subnets as the preferred subnet using the `WindowsConfiguration > PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID` properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the *Amazon FSx for Windows User Guide* and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the *Amazon FSx for ONTAP User Guide* .
	//
	// For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre deployment types, provide exactly one subnet ID. The file server is launched in that subnet's Availability Zone.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This value is required if `FileSystemType` is set to `WINDOWS` .
	WindowsConfiguration() interface{}
	SetWindowsConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFileSystem
type jsiiProxy_CfnFileSystem struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFileSystem) AttrDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) AttrLustreMountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLustreMountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) AttrRootVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRootVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) FileSystemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) LustreConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lustreConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) OntapConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ontapConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) OpenZfsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openZfsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) WindowsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::FSx::FileSystem`.
func NewCfnFileSystem(scope constructs.Construct, id *string, props *CfnFileSystemProps) CfnFileSystem {
	_init_.Initialize()

	if err := validateNewCfnFileSystemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::FileSystem`.
func NewCfnFileSystem_Override(c CfnFileSystem, scope constructs.Construct, id *string, props *CfnFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetFileSystemType(val *string) {
	if err := j.validateSetFileSystemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemType",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetFileSystemTypeVersion(val *string) {
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetLustreConfiguration(val interface{}) {
	if err := j.validateSetLustreConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lustreConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetOntapConfiguration(val interface{}) {
	if err := j.validateSetOntapConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ontapConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetOpenZfsConfiguration(val interface{}) {
	if err := j.validateSetOpenZfsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openZfsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetWindowsConfiguration(val interface{}) {
	if err := j.validateSetWindowsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFileSystem_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileSystem_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFileSystem_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFileSystem_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFileSystem_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_fsx.CfnFileSystem",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFileSystem) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFileSystem) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFileSystem) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFileSystem) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileSystem) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFileSystem) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

