package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsfsx/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/constructs-go/constructs/v3"
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnFileSystem := fsx.NewCfnFileSystem(this, jsii.String("MyCfnFileSystem"), &cfnFileSystemProps{
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
//   })
//
type CfnFileSystem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the FSx for Windows file system's DNSName.
	//
	// Example: `amznfsxp1honlek.corp.example.com`
	AttrDnsName() *string
	// Returns the file system's LustreMountName.
	//
	// Example for SCRATCH_1 deployment types: This value is always `fsx` .
	//
	// Example for SCRATCH_2 and PERSISTENT deployment types: `2p3fhbmv`.
	AttrLustreMountName() *string
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Experimental.
	Ref() *string
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This value is required if `FileSystemType` is set to `WINDOWS` .
	WindowsConfiguration() interface{}
	SetWindowsConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnFileSystem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnFileSystem(scope awscdk.Construct, id *string, props *CfnFileSystemProps) CfnFileSystem {
	_init_.Initialize()

	j := jsiiProxy_CfnFileSystem{}

	_jsii_.Create(
		"monocdk.aws_fsx.CfnFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::FileSystem`.
func NewCfnFileSystem_Override(c CfnFileSystem, scope awscdk.Construct, id *string, props *CfnFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.CfnFileSystem",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetFileSystemType(val *string) {
	_jsii_.Set(
		j,
		"fileSystemType",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetFileSystemTypeVersion(val *string) {
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetLustreConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"lustreConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetOntapConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"ontapConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetOpenZfsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"openZfsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem) SetWindowsConfiguration(val interface{}) {
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
// Experimental.
func CfnFileSystem_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnFileSystem",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFileSystem_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnFileSystem",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnFileSystem",
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
		"monocdk.aws_fsx.CfnFileSystem",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFileSystem) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFileSystem) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFileSystem) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFileSystem) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFileSystem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFileSystem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFileSystem) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFileSystem) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnFileSystem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnFileSystem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileSystem) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   auditLogConfigurationProperty := &auditLogConfigurationProperty{
//   	fileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   	fileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//
//   	// the properties below are optional
//   	auditLogDestination: jsii.String("auditLogDestination"),
//   }
//
type CfnFileSystem_AuditLogConfigurationProperty struct {
	// Sets which attempt type is logged by Amazon FSx for file and folder accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access files or folders are logged.
	// - `FAILURE_ONLY` - only failed attempts to access files or folders are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access files or folders are logged.
	// - `DISABLED` - access auditing of files and folders is turned off.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel" yaml:"fileAccessAuditLogLevel"`
	// Sets which attempt type is logged by Amazon FSx for file share accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access file shares are logged.
	// - `FAILURE_ONLY` - only failed attempts to access file shares are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access file shares are logged.
	// - `DISABLED` - access auditing of file shares is turned off.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel" yaml:"fileShareAccessAuditLogLevel"`
	// The Amazon Resource Name (ARN) for the destination of the audit logs.
	//
	// The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN.
	//
	// The name of the Amazon CloudWatch Logs log group must begin with the `/aws/fsx` prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the `aws-fsx` prefix.
	//
	// The destination ARN (either CloudWatch Logs log group or Kinesis Data Firehose delivery stream) must be in the same AWS partition, AWS Region , and AWS account as your Amazon FSx file system.
	AuditLogDestination *string `json:"auditLogDestination" yaml:"auditLogDestination"`
}

// Specifies who can mount an OpenZFS file system and the options available while mounting the file system.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   clientConfigurationsProperty := &clientConfigurationsProperty{
//   	clients: jsii.String("clients"),
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   }
//
type CfnFileSystem_ClientConfigurationsProperty struct {
	// A value that specifies who can mount the file system.
	//
	// You can provide a wildcard character ( `*` ), an IP address ( `0.0.0.0` ), or a CIDR address ( `192.0.2.0/24` ). By default, Amazon FSx uses the wildcard character when specifying the client.
	Clients *string `json:"clients" yaml:"clients"`
	// The options to use when mounting the file system.
	//
	// For a list of options that you can use with Network File System (NFS), see the [exports(5) - Linux man page](https://docs.aws.amazon.com/https://linux.die.net/man/5/exports) . When choosing your options, consider the following:
	//
	// - `crossmnt` is used by default. If you don't specify `crossmnt` when changing the client configuration, you won't be able to see or access snapshots in your file system's snapshot directory.
	// - `sync` is used by default. If you instead specify `async` , the system acknowledges writes before writing to disk. If the system crashes before the writes are finished, you lose the unwritten data.
	Options *[]*string `json:"options" yaml:"options"`
}

// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
//
// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   diskIopsConfigurationProperty := &diskIopsConfigurationProperty{
//   	iops: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   }
//
type CfnFileSystem_DiskIopsConfigurationProperty struct {
	// The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Specifies whether the number of IOPS for the file system is using the system default ( `AUTOMATIC` ) or was provisioned by the customer ( `USER_PROVISIONED` ).
	Mode *string `json:"mode" yaml:"mode"`
}

// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
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
	AutoImportPolicy *string `json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean flag indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value. Only valid for use with `PERSISTENT_1` deployment types.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Sets the data compression configuration for the file system. `DataCompressionType` can have the following values:.
	//
	// - `NONE` - (Default) Data compression is turned off when the file system is created.
	// - `LZ4` - Data compression is turned on with the LZ4 algorithm.
	//
	// For more information, see [Lustre data compression](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html) in the *Amazon FSx for Lustre User Guide* .
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
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
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// The type of drive cache used by `PERSISTENT_1` file systems that are provisioned with HDD storage devices.
	//
	// This parameter is required when storage type is HDD. Set this property to `READ` to improve the performance for frequently accessed files by caching up to 20% of the total storage capacity of the file system.
	//
	// This parameter is required when `StorageType` is set to `HDD` and `DeploymentType` is `PERSISTENT_1` .
	DriveCacheType *string `json:"driveCacheType" yaml:"driveCacheType"`
	// (Optional) Available with `Scratch` and `Persistent_1` deployment types.
	//
	// Specifies the path in the Amazon S3 bucket where the root of your Amazon FSx file system is exported. The path must use the same Amazon S3 bucket as specified in ImportPath. You can provide an optional prefix to which new and changed data is to be exported from your Amazon FSx for Lustre file system. If an `ExportPath` value is not provided, Amazon FSx sets a default export path, `s3://import-bucket/FSxLustre[creation-timestamp]` . The timestamp is in UTC format, for example `s3://import-bucket/FSxLustre20181105T222312Z` .
	//
	// The Amazon S3 export bucket must be the same as the import bucket specified by `ImportPath` . If you specify only a bucket name, such as `s3://import-bucket` , you get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is overwritten on export. If you provide a custom prefix in the export path, such as `s3://import-bucket/[custom-optional-prefix]` , Amazon FSx exports the contents of your file system to that export prefix in the Amazon S3 bucket.
	//
	// > This parameter is not supported for file systems using the `Persistent_2` deployment type.
	ExportPath *string `json:"exportPath" yaml:"exportPath"`
	// (Optional) For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// (Optional) The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// The root of your FSx for Lustre file system will be mapped to the root of the Amazon S3 bucket you select. An example is `s3://import-bucket/optional-prefix` . If you specify a prefix after the Amazon S3 bucket name, only object keys with that prefix are loaded into the file system.
	//
	// > This parameter is not supported for Lustre file systems using the `Persistent_2` deployment type.
	ImportPath *string `json:"importPath" yaml:"importPath"`
	// Required with `PERSISTENT_1` and `PERSISTENT_2` deployment types, provisions the amount of read and write throughput for each 1 tebibyte (TiB) of file system storage capacity, in MB/s/TiB.
	//
	// File system throughput capacity is calculated by multiplying ﬁle system storage capacity (TiB) by the `PerUnitStorageThroughput` (MB/s/TiB). For a 2.4-TiB ﬁle system, provisioning 50 MB/s/TiB of `PerUnitStorageThroughput` yields 120 MB/s of ﬁle system throughput. You pay for the amount of throughput that you provision.
	//
	// Valid values:
	//
	// - For `PERSISTENT_1` SSD storage: 50, 100, 200 MB/s/TiB.
	// - For `PERSISTENT_1` HDD storage: 12, 40 MB/s/TiB.
	// - For `PERSISTENT_2` SSD storage: 125, 250, 500, 1000 MB/s/TiB.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

// The configuration object for mounting a file system.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   nfsExportsProperty := &nfsExportsProperty{
//   	clientConfigurations: []interface{}{
//   		&clientConfigurationsProperty{
//   			clients: jsii.String("clients"),
//   			options: []*string{
//   				jsii.String("options"),
//   			},
//   		},
//   	},
//   }
//
type CfnFileSystem_NfsExportsProperty struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system.
	ClientConfigurations interface{} `json:"clientConfigurations" yaml:"clientConfigurations"`
}

// The configuration for this Amazon FSx for NetApp ONTAP file system.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   ontapConfigurationProperty := &ontapConfigurationProperty{
//   	deploymentType: jsii.String("deploymentType"),
//
//   	// the properties below are optional
//   	automaticBackupRetentionDays: jsii.Number(123),
//   	dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	diskIopsConfiguration: &diskIopsConfigurationProperty{
//   		iops: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   	},
//   	endpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   	fsxAdminPassword: jsii.String("fsxAdminPassword"),
//   	preferredSubnetId: jsii.String("preferredSubnetId"),
//   	routeTableIds: []*string{
//   		jsii.String("routeTableIds"),
//   	},
//   	throughputCapacity: jsii.Number(123),
//   	weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_OntapConfigurationProperty struct {
	// Specifies the FSx for ONTAP file system deployment type to use in creating the file system.
	//
	// `MULTI_AZ_1` is the supported ONTAP deployment type.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS configuration for the FSx for ONTAP file system.
	DiskIopsConfiguration interface{} `json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Specifies the IP address range in which the endpoints to access your file system will be created.
	//
	// By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	//
	// > The Endpoint IP address range you select for your file system must exist outside the VPC's CIDR range and must be at least /30 or larger.
	EndpointIpAddressRange *string `json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// The ONTAP administrative password for the `fsxadmin` user with which you administer your file system using the NetApp ONTAP CLI and REST API.
	FsxAdminPassword *string `json:"fsxAdminPassword" yaml:"fsxAdminPassword"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located.
	PreferredSubnetId *string `json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Specifies the virtual private cloud (VPC) route tables in which your file system's endpoints will be created.
	//
	// You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds *[]*string `json:"routeTableIds" yaml:"routeTableIds"`
	// Sets the throughput capacity for the file system that you're creating.
	//
	// Valid values are 128, 256, 512, 1024, and 2048 MBps.
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

// The OpenZFS configuration for the file system that's being created.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
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
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A Boolean value indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToVolumes interface{} `json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
	//
	// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
	DiskIopsConfiguration interface{} `json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// To delete a file system if there are child volumes present below the root volume, use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS` .
	//
	// If your file system has child volumes and you don't use this option, the delete request will fail.
	Options *[]*string `json:"options" yaml:"options"`
	// The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS file system.
	//
	// All volumes are children of the root volume.
	RootVolumeConfiguration interface{} `json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MB/s).
	//
	// Valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MB/s. You pay for additional throughput capacity that you provision.
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

// The configuration of an Amazon FSx for OpenZFS root volume.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   rootVolumeConfigurationProperty := &rootVolumeConfigurationProperty{
//   	copyTagsToSnapshots: jsii.Boolean(false),
//   	dataCompressionType: jsii.String("dataCompressionType"),
//   	nfsExports: []interface{}{
//   		&nfsExportsProperty{
//   			clientConfigurations: []interface{}{
//   				&clientConfigurationsProperty{
//   					clients: jsii.String("clients"),
//   					options: []*string{
//   						jsii.String("options"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	readOnly: jsii.Boolean(false),
//   	recordSizeKiB: jsii.Number(123),
//   	userAndGroupQuotas: []interface{}{
//   		&userAndGroupQuotasProperty{
//   			id: jsii.Number(123),
//   			storageCapacityQuotaGiB: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnFileSystem_RootVolumeConfigurationProperty struct {
	// A Boolean value indicating whether tags for the volume should be copied to snapshots of the volume.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToSnapshots interface{} `json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Specifies the method used to compress the data on the volume. The compression type is `NONE` by default.
	//
	// - `NONE` - Doesn't compress the data on the volume. `NONE` is the default.
	// - `ZSTD` - Compresses the data in the volume using the Zstandard (ZSTD) compression algorithm. Compared to LZ4, Z-Standard provides a better compression ratio to minimize on-disk storage utilization.
	// - `LZ4` - Compresses the data in the volume using the LZ4 compression algorithm. Compared to Z-Standard, LZ4 is less compute-intensive and delivers higher write throughput speeds.
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
	// The configuration object for mounting a file system.
	NfsExports interface{} `json:"nfsExports" yaml:"nfsExports"`
	// A Boolean value indicating whether the volume is read-only.
	//
	// Setting this value to `true` can be useful after you have completed changes to a volume and no longer want changes to occur.
	ReadOnly interface{} `json:"readOnly" yaml:"readOnly"`
	// Specifies the record size of an OpenZFS root volume, in kibibytes (KiB).
	//
	// Valid values are 4, 8, 16, 32, 64, 128, 256, 512, or 1024 KiB. The default is 128 KiB. Most workloads should use the default record size. Database workflows can benefit from a smaller record size, while streaming workflows can benefit from a larger record size. For additional guidance on setting a custom record size, see [Tips for maximizing performance](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance.html#performance-tips-zfs) in the *Amazon FSx for OpenZFS User Guide* .
	RecordSizeKiB *float64 `json:"recordSizeKiB" yaml:"recordSizeKiB"`
	// An object specifying how much storage users or groups can use on the volume.
	UserAndGroupQuotas interface{} `json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

// The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an ONTAP storage virtual machine (SVM) to a self-managed (including on-premises) Microsoft Active Directory (AD) directory.
//
// For more information, see [Using Amazon FSx with your self-managed Microsoft Active Directory](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html) or [Managing SVMs](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   selfManagedActiveDirectoryConfigurationProperty := &selfManagedActiveDirectoryConfigurationProperty{
//   	dnsIps: []*string{
//   		jsii.String("dnsIps"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	password: jsii.String("password"),
//   	userName: jsii.String("userName"),
//   }
//
type CfnFileSystem_SelfManagedActiveDirectoryConfigurationProperty struct {
	// A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	DnsIps *[]*string `json:"dnsIps" yaml:"dnsIps"`
	// The fully qualified domain name of the self-managed AD directory, such as `corp.example.com` .
	DomainName *string `json:"domainName" yaml:"domainName"`
	// (Optional) The name of the domain group whose members are granted administrative privileges for the file system.
	//
	// Administrative privileges include taking ownership of files and folders, setting audit controls (audit ACLs) on files and folders, and administering the file system remotely by using the FSx Remote PowerShell. The group that you specify must already exist in your domain. If you don't provide one, your AD domain's Domain Admins group is used.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// (Optional) The fully qualified distinguished name of the organizational unit within your self-managed AD directory.
	//
	// Amazon FSx only accepts OU as the direct parent of the file system. An example is `OU=FSx,DC=yourdomain,DC=corp,DC=com` . To learn more, see [RFC 2253](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2253) . If none is provided, the FSx file system is created in the default location of your self-managed AD directory.
	//
	// > Only Organizational Unit (OU) objects can be the direct parent of the file system that you're creating.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Password *string `json:"password" yaml:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// This account must have the permission to join computers to the domain in the organizational unit provided in `OrganizationalUnitDistinguishedName` , or in the default location of your AD domain.
	UserName *string `json:"userName" yaml:"userName"`
}

// The configuration for how much storage a user or group can use on the volume.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   userAndGroupQuotasProperty := &userAndGroupQuotasProperty{
//   	id: jsii.Number(123),
//   	storageCapacityQuotaGiB: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnFileSystem_UserAndGroupQuotasProperty struct {
	// The ID of the user or group.
	Id *float64 `json:"id" yaml:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	StorageCapacityQuotaGiB *float64 `json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	Type *string `json:"type" yaml:"type"`
}

// The Microsoft Windows configuration for the file system that's being created.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   windowsConfigurationProperty := &windowsConfigurationProperty{
//   	throughputCapacity: jsii.Number(123),
//
//   	// the properties below are optional
//   	activeDirectoryId: jsii.String("activeDirectoryId"),
//   	aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	auditLogConfiguration: &auditLogConfigurationProperty{
//   		fileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   		fileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//
//   		// the properties below are optional
//   		auditLogDestination: jsii.String("auditLogDestination"),
//   	},
//   	automaticBackupRetentionDays: jsii.Number(123),
//   	copyTagsToBackups: jsii.Boolean(false),
//   	dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	deploymentType: jsii.String("deploymentType"),
//   	preferredSubnetId: jsii.String("preferredSubnetId"),
//   	selfManagedActiveDirectoryConfiguration: &selfManagedActiveDirectoryConfigurationProperty{
//   		dnsIps: []*string{
//   			jsii.String("dnsIps"),
//   		},
//   		domainName: jsii.String("domainName"),
//   		fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		password: jsii.String("password"),
//   		userName: jsii.String("userName"),
//   	},
//   	weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_WindowsConfigurationProperty struct {
	// Sets the throughput capacity of an Amazon FSx file system, measured in megabytes per second (MB/s), in 2 to the *n* th increments, between 2^3 (8) and 2^11 (2048).
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// The ID for an existing AWS Managed Microsoft Active Directory (AD) instance that the file system should join when it's created.
	//
	// Required if you are joining the file system to an existing AWS Managed Microsoft AD.
	ActiveDirectoryId *string `json:"activeDirectoryId" yaml:"activeDirectoryId"`
	// An array of one or more DNS alias names that you want to associate with the Amazon FSx file system.
	//
	// Aliases allow you to use existing DNS names to access the data in your Amazon FSx file system. You can associate up to 50 aliases with a file system at any time.
	//
	// For more information, see [Working with DNS Aliases](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html) and [Walkthrough 5: Using DNS aliases to access your file system](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/walkthrough05-file-system-custom-CNAME.html) , including additional steps you must take to be able to access your file system using a DNS alias.
	//
	// An alias name has to meet the following requirements:
	//
	// - Formatted as a fully-qualified domain name (FQDN), `hostname.domain` , for example, `accounting.example.com` .
	// - Can contain alphanumeric characters, the underscore (_), and the hyphen (-).
	// - Cannot start or end with a hyphen.
	// - Can start with a numeric.
	//
	// For DNS alias names, Amazon FSx stores alphabetical characters as lowercase letters (a-z), regardless of how you specify them: as uppercase letters, lowercase letters, or the corresponding letters in escape codes.
	Aliases *[]*string `json:"aliases" yaml:"aliases"`
	// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.
	AuditLogConfiguration interface{} `json:"auditLogConfiguration" yaml:"auditLogConfiguration"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A boolean flag indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Specifies the file system deployment type, valid values are the following:.
	//
	// - `MULTI_AZ_1` - Deploys a high availability file system that is configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability. You can only deploy a Multi-AZ file system in AWS Regions that have a minimum of three Availability Zones. Also supports HDD storage type
	// - `SINGLE_AZ_1` - (Default) Choose to deploy a file system that is configured for single AZ redundancy.
	// - `SINGLE_AZ_2` - The latest generation Single AZ file system. Specifies a file system that is configured for single AZ redundancy and supports HDD storage type.
	//
	// For more information, see [Availability and Durability: Single-AZ and Multi-AZ File Systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) .
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located. For in- AWS applications, we recommend that you launch your clients in the same availability zone as your preferred file server to reduce cross-availability zone data transfer costs and minimize latency.
	PreferredSubnetId *string `json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an ONTAP storage virtual machine (SVM) to a self-managed (including on-premises) Microsoft Active Directory (AD) directory.
	//
	// For more information, see [Using Amazon FSx with your self-managed Microsoft Active Directory](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html) or [Managing SVMs](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) .
	SelfManagedActiveDirectoryConfiguration interface{} `json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

// Properties for defining a `CfnFileSystem`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
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
	FileSystemType *string `json:"fileSystemType" yaml:"fileSystemType"`
	// Specifies the IDs of the subnets that the file system will be accessible from.
	//
	// For Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet IDs, one for the preferred file server and one for the standby file server. You specify one of these subnets as the preferred subnet using the `WindowsConfiguration > PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID` properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the *Amazon FSx for Windows User Guide* and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the *Amazon FSx for ONTAP User Guide* .
	//
	// For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre deployment types, provide exactly one subnet ID. The file server is launched in that subnet's Availability Zone.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The ID of the file system backup that you are using to create a file system.
	//
	// For more information, see [CreateFileSystemFromBackup](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemFromBackup.html) .
	BackupId *string `json:"backupId" yaml:"backupId"`
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
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
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
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Lustre configuration for the file system being created.
	//
	// > The following parameters are not supported for file systems with the Lustre `Persistent_2` deployment type.
	// >
	// > - `AutoImportPolicy`
	// > - `ExportPath`
	// > - `ImportedChunkSize`
	// > - `ImportPath`.
	LustreConfiguration interface{} `json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The ONTAP configuration properties of the FSx for ONTAP file system that you are creating.
	OntapConfiguration interface{} `json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The Amazon FSx for OpenZFS configuration properties for the file system that you are creating.
	OpenZfsConfiguration interface{} `json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
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
	StorageCapacity *float64 `json:"storageCapacity" yaml:"storageCapacity"`
	// Sets the storage type for the file system that you're creating. Valid values are `SSD` and `HDD` .
	//
	// - Set to `SSD` to use solid state drive storage. SSD is supported on all Windows, Lustre, ONTAP, and OpenZFS deployment types.
	// - Set to `HDD` to use hard disk drive storage. HDD is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types, and on `PERSISTENT_1` Lustre file system deployment types.
	//
	// Default value is `SSD` . For more information, see [Storage type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the *FSx for Windows File Server User Guide* and [Multiple storage options](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html#storage-options) in the *FSx for Lustre User Guide* .
	StorageType *string `json:"storageType" yaml:"storageType"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This value is required if `FileSystemType` is set to `WINDOWS` .
	WindowsConfiguration interface{} `json:"windowsConfiguration" yaml:"windowsConfiguration"`
}

// A CloudFormation `AWS::FSx::Snapshot`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnSnapshot := fsx.NewCfnSnapshot(this, jsii.String("MyCfnSnapshot"), &cfnSnapshotProps{
//   	name: jsii.String("name"),
//   	volumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnSnapshot interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrResourceArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// `AWS::FSx::Snapshot.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::FSx::Snapshot.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::FSx::Snapshot.VolumeId`.
	VolumeId() *string
	SetVolumeId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSnapshot
type jsiiProxy_CfnSnapshot struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSnapshot) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshot) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}


// Create a new `AWS::FSx::Snapshot`.
func NewCfnSnapshot(scope awscdk.Construct, id *string, props *CfnSnapshotProps) CfnSnapshot {
	_init_.Initialize()

	j := jsiiProxy_CfnSnapshot{}

	_jsii_.Create(
		"monocdk.aws_fsx.CfnSnapshot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::Snapshot`.
func NewCfnSnapshot_Override(c CfnSnapshot, scope awscdk.Construct, id *string, props *CfnSnapshotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.CfnSnapshot",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSnapshot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshot) SetVolumeId(val *string) {
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnSnapshot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnSnapshot",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSnapshot_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnSnapshot",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSnapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnSnapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSnapshot_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_fsx.CfnSnapshot",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSnapshot) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSnapshot) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSnapshot) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSnapshot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSnapshot) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSnapshot) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSnapshot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSnapshot) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSnapshot) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSnapshot) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSnapshot) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSnapshot) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSnapshot) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSnapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshot) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSnapshot`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnSnapshotProps := &cfnSnapshotProps{
//   	name: jsii.String("name"),
//   	volumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSnapshotProps struct {
	// `AWS::FSx::Snapshot.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::FSx::Snapshot.VolumeId`.
	VolumeId *string `json:"volumeId" yaml:"volumeId"`
	// `AWS::FSx::Snapshot.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::FSx::StorageVirtualMachine`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnStorageVirtualMachine := fsx.NewCfnStorageVirtualMachine(this, jsii.String("MyCfnStorageVirtualMachine"), &cfnStorageVirtualMachineProps{
//   	fileSystemId: jsii.String("fileSystemId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   		netBiosName: jsii.String("netBiosName"),
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
//   	},
//   	rootVolumeSecurityStyle: jsii.String("rootVolumeSecurityStyle"),
//   	svmAdminPassword: jsii.String("svmAdminPassword"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnStorageVirtualMachine interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::FSx::StorageVirtualMachine.ActiveDirectoryConfiguration`.
	ActiveDirectoryConfiguration() interface{}
	SetActiveDirectoryConfiguration(val interface{})
	AttrResourceArn() *string
	AttrStorageVirtualMachineId() *string
	AttrUuid() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::FSx::StorageVirtualMachine.FileSystemId`.
	FileSystemId() *string
	SetFileSystemId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// `AWS::FSx::StorageVirtualMachine.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::FSx::StorageVirtualMachine.RootVolumeSecurityStyle`.
	RootVolumeSecurityStyle() *string
	SetRootVolumeSecurityStyle(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::FSx::StorageVirtualMachine.SvmAdminPassword`.
	SvmAdminPassword() *string
	SetSvmAdminPassword(val *string)
	// `AWS::FSx::StorageVirtualMachine.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStorageVirtualMachine
type jsiiProxy_CfnStorageVirtualMachine struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStorageVirtualMachine) ActiveDirectoryConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) AttrStorageVirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStorageVirtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) AttrUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) RootVolumeSecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeSecurityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SvmAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"svmAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageVirtualMachine) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::FSx::StorageVirtualMachine`.
func NewCfnStorageVirtualMachine(scope awscdk.Construct, id *string, props *CfnStorageVirtualMachineProps) CfnStorageVirtualMachine {
	_init_.Initialize()

	j := jsiiProxy_CfnStorageVirtualMachine{}

	_jsii_.Create(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::StorageVirtualMachine`.
func NewCfnStorageVirtualMachine_Override(c CfnStorageVirtualMachine, scope awscdk.Construct, id *string, props *CfnStorageVirtualMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SetActiveDirectoryConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"activeDirectoryConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SetFileSystemId(val *string) {
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SetRootVolumeSecurityStyle(val *string) {
	_jsii_.Set(
		j,
		"rootVolumeSecurityStyle",
		val,
	)
}

func (j *jsiiProxy_CfnStorageVirtualMachine) SetSvmAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"svmAdminPassword",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnStorageVirtualMachine_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStorageVirtualMachine_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnStorageVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStorageVirtualMachine_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_fsx.CfnStorageVirtualMachine",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStorageVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageVirtualMachine) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   activeDirectoryConfigurationProperty := &activeDirectoryConfigurationProperty{
//   	netBiosName: jsii.String("netBiosName"),
//   	selfManagedActiveDirectoryConfiguration: &selfManagedActiveDirectoryConfigurationProperty{
//   		dnsIps: []*string{
//   			jsii.String("dnsIps"),
//   		},
//   		domainName: jsii.String("domainName"),
//   		fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		password: jsii.String("password"),
//   		userName: jsii.String("userName"),
//   	},
//   }
//
type CfnStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// `CfnStorageVirtualMachine.ActiveDirectoryConfigurationProperty.NetBiosName`.
	NetBiosName *string `json:"netBiosName" yaml:"netBiosName"`
	// `CfnStorageVirtualMachine.ActiveDirectoryConfigurationProperty.SelfManagedActiveDirectoryConfiguration`.
	SelfManagedActiveDirectoryConfiguration interface{} `json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   selfManagedActiveDirectoryConfigurationProperty := &selfManagedActiveDirectoryConfigurationProperty{
//   	dnsIps: []*string{
//   		jsii.String("dnsIps"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	password: jsii.String("password"),
//   	userName: jsii.String("userName"),
//   }
//
type CfnStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty struct {
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.DnsIps`.
	DnsIps *[]*string `json:"dnsIps" yaml:"dnsIps"`
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.DomainName`.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.FileSystemAdministratorsGroup`.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.OrganizationalUnitDistinguishedName`.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.Password`.
	Password *string `json:"password" yaml:"password"`
	// `CfnStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationProperty.UserName`.
	UserName *string `json:"userName" yaml:"userName"`
}

// Properties for defining a `CfnStorageVirtualMachine`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnStorageVirtualMachineProps := &cfnStorageVirtualMachineProps{
//   	fileSystemId: jsii.String("fileSystemId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   		netBiosName: jsii.String("netBiosName"),
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
//   	},
//   	rootVolumeSecurityStyle: jsii.String("rootVolumeSecurityStyle"),
//   	svmAdminPassword: jsii.String("svmAdminPassword"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStorageVirtualMachineProps struct {
	// `AWS::FSx::StorageVirtualMachine.FileSystemId`.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// `AWS::FSx::StorageVirtualMachine.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::FSx::StorageVirtualMachine.ActiveDirectoryConfiguration`.
	ActiveDirectoryConfiguration interface{} `json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// `AWS::FSx::StorageVirtualMachine.RootVolumeSecurityStyle`.
	RootVolumeSecurityStyle *string `json:"rootVolumeSecurityStyle" yaml:"rootVolumeSecurityStyle"`
	// `AWS::FSx::StorageVirtualMachine.SvmAdminPassword`.
	SvmAdminPassword *string `json:"svmAdminPassword" yaml:"svmAdminPassword"`
	// `AWS::FSx::StorageVirtualMachine.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::FSx::Volume`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnVolume := fsx.NewCfnVolume(this, jsii.String("MyCfnVolume"), &cfnVolumeProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	backupId: jsii.String("backupId"),
//   	ontapConfiguration: &ontapConfigurationProperty{
//   		junctionPath: jsii.String("junctionPath"),
//   		sizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		storageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		storageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   		// the properties below are optional
//   		securityStyle: jsii.String("securityStyle"),
//   		tieringPolicy: &tieringPolicyProperty{
//   			coolingPeriod: jsii.Number(123),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	openZfsConfiguration: &openZFSConfigurationProperty{
//   		parentVolumeId: jsii.String("parentVolumeId"),
//
//   		// the properties below are optional
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
//   		options: []*string{
//   			jsii.String("options"),
//   		},
//   		originSnapshot: &originSnapshotProperty{
//   			copyStrategy: jsii.String("copyStrategy"),
//   			snapshotArn: jsii.String("snapshotArn"),
//   		},
//   		readOnly: jsii.Boolean(false),
//   		recordSizeKiB: jsii.Number(123),
//   		storageCapacityQuotaGiB: jsii.Number(123),
//   		storageCapacityReservationGiB: jsii.Number(123),
//   		userAndGroupQuotas: []interface{}{
//   			&userAndGroupQuotasProperty{
//   				id: jsii.Number(123),
//   				storageCapacityQuotaGiB: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeType: jsii.String("volumeType"),
//   })
//
type CfnVolume interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrResourceArn() *string
	AttrUuid() *string
	AttrVolumeId() *string
	// `AWS::FSx::Volume.BackupId`.
	BackupId() *string
	SetBackupId(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// `AWS::FSx::Volume.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::FSx::Volume.OntapConfiguration`.
	OntapConfiguration() interface{}
	SetOntapConfiguration(val interface{})
	// `AWS::FSx::Volume.OpenZFSConfiguration`.
	OpenZfsConfiguration() interface{}
	SetOpenZfsConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::FSx::Volume.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::FSx::Volume.VolumeType`.
	VolumeType() *string
	SetVolumeType(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVolume
type jsiiProxy_CfnVolume struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVolume) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) AttrUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) AttrVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) OntapConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ontapConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) OpenZfsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openZfsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}


// Create a new `AWS::FSx::Volume`.
func NewCfnVolume(scope awscdk.Construct, id *string, props *CfnVolumeProps) CfnVolume {
	_init_.Initialize()

	j := jsiiProxy_CfnVolume{}

	_jsii_.Create(
		"monocdk.aws_fsx.CfnVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::Volume`.
func NewCfnVolume_Override(c CfnVolume, scope awscdk.Construct, id *string, props *CfnVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.CfnVolume",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVolume) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetOntapConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"ontapConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetOpenZfsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"openZfsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnVolume) SetVolumeType(val *string) {
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnVolume_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnVolume",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVolume_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnVolume",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.CfnVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVolume_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_fsx.CfnVolume",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVolume) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVolume) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVolume) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVolume) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVolume) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVolume) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVolume) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVolume) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVolume) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVolume) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVolume) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVolume) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVolume) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   clientConfigurationsProperty := &clientConfigurationsProperty{
//   	clients: jsii.String("clients"),
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   }
//
type CfnVolume_ClientConfigurationsProperty struct {
	// `CfnVolume.ClientConfigurationsProperty.Clients`.
	Clients *string `json:"clients" yaml:"clients"`
	// `CfnVolume.ClientConfigurationsProperty.Options`.
	Options *[]*string `json:"options" yaml:"options"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   nfsExportsProperty := &nfsExportsProperty{
//   	clientConfigurations: []interface{}{
//   		&clientConfigurationsProperty{
//   			clients: jsii.String("clients"),
//   			options: []*string{
//   				jsii.String("options"),
//   			},
//   		},
//   	},
//   }
//
type CfnVolume_NfsExportsProperty struct {
	// `CfnVolume.NfsExportsProperty.ClientConfigurations`.
	ClientConfigurations interface{} `json:"clientConfigurations" yaml:"clientConfigurations"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   ontapConfigurationProperty := &ontapConfigurationProperty{
//   	junctionPath: jsii.String("junctionPath"),
//   	sizeInMegabytes: jsii.String("sizeInMegabytes"),
//   	storageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   	storageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   	// the properties below are optional
//   	securityStyle: jsii.String("securityStyle"),
//   	tieringPolicy: &tieringPolicyProperty{
//   		coolingPeriod: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnVolume_OntapConfigurationProperty struct {
	// `CfnVolume.OntapConfigurationProperty.JunctionPath`.
	JunctionPath *string `json:"junctionPath" yaml:"junctionPath"`
	// `CfnVolume.OntapConfigurationProperty.SizeInMegabytes`.
	SizeInMegabytes *string `json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// `CfnVolume.OntapConfigurationProperty.StorageEfficiencyEnabled`.
	StorageEfficiencyEnabled *string `json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// `CfnVolume.OntapConfigurationProperty.StorageVirtualMachineId`.
	StorageVirtualMachineId *string `json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// `CfnVolume.OntapConfigurationProperty.SecurityStyle`.
	SecurityStyle *string `json:"securityStyle" yaml:"securityStyle"`
	// `CfnVolume.OntapConfigurationProperty.TieringPolicy`.
	TieringPolicy interface{} `json:"tieringPolicy" yaml:"tieringPolicy"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   openZFSConfigurationProperty := &openZFSConfigurationProperty{
//   	parentVolumeId: jsii.String("parentVolumeId"),
//
//   	// the properties below are optional
//   	copyTagsToSnapshots: jsii.Boolean(false),
//   	dataCompressionType: jsii.String("dataCompressionType"),
//   	nfsExports: []interface{}{
//   		&nfsExportsProperty{
//   			clientConfigurations: []interface{}{
//   				&clientConfigurationsProperty{
//   					clients: jsii.String("clients"),
//   					options: []*string{
//   						jsii.String("options"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   	originSnapshot: &originSnapshotProperty{
//   		copyStrategy: jsii.String("copyStrategy"),
//   		snapshotArn: jsii.String("snapshotArn"),
//   	},
//   	readOnly: jsii.Boolean(false),
//   	recordSizeKiB: jsii.Number(123),
//   	storageCapacityQuotaGiB: jsii.Number(123),
//   	storageCapacityReservationGiB: jsii.Number(123),
//   	userAndGroupQuotas: []interface{}{
//   		&userAndGroupQuotasProperty{
//   			id: jsii.Number(123),
//   			storageCapacityQuotaGiB: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnVolume_OpenZFSConfigurationProperty struct {
	// `CfnVolume.OpenZFSConfigurationProperty.ParentVolumeId`.
	ParentVolumeId *string `json:"parentVolumeId" yaml:"parentVolumeId"`
	// `CfnVolume.OpenZFSConfigurationProperty.CopyTagsToSnapshots`.
	CopyTagsToSnapshots interface{} `json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// `CfnVolume.OpenZFSConfigurationProperty.DataCompressionType`.
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
	// `CfnVolume.OpenZFSConfigurationProperty.NfsExports`.
	NfsExports interface{} `json:"nfsExports" yaml:"nfsExports"`
	// `CfnVolume.OpenZFSConfigurationProperty.Options`.
	Options *[]*string `json:"options" yaml:"options"`
	// `CfnVolume.OpenZFSConfigurationProperty.OriginSnapshot`.
	OriginSnapshot interface{} `json:"originSnapshot" yaml:"originSnapshot"`
	// `CfnVolume.OpenZFSConfigurationProperty.ReadOnly`.
	ReadOnly interface{} `json:"readOnly" yaml:"readOnly"`
	// `CfnVolume.OpenZFSConfigurationProperty.RecordSizeKiB`.
	RecordSizeKiB *float64 `json:"recordSizeKiB" yaml:"recordSizeKiB"`
	// `CfnVolume.OpenZFSConfigurationProperty.StorageCapacityQuotaGiB`.
	StorageCapacityQuotaGiB *float64 `json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// `CfnVolume.OpenZFSConfigurationProperty.StorageCapacityReservationGiB`.
	StorageCapacityReservationGiB *float64 `json:"storageCapacityReservationGiB" yaml:"storageCapacityReservationGiB"`
	// `CfnVolume.OpenZFSConfigurationProperty.UserAndGroupQuotas`.
	UserAndGroupQuotas interface{} `json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   originSnapshotProperty := &originSnapshotProperty{
//   	copyStrategy: jsii.String("copyStrategy"),
//   	snapshotArn: jsii.String("snapshotArn"),
//   }
//
type CfnVolume_OriginSnapshotProperty struct {
	// `CfnVolume.OriginSnapshotProperty.CopyStrategy`.
	CopyStrategy *string `json:"copyStrategy" yaml:"copyStrategy"`
	// `CfnVolume.OriginSnapshotProperty.SnapshotARN`.
	SnapshotArn *string `json:"snapshotArn" yaml:"snapshotArn"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   tieringPolicyProperty := &tieringPolicyProperty{
//   	coolingPeriod: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnVolume_TieringPolicyProperty struct {
	// `CfnVolume.TieringPolicyProperty.CoolingPeriod`.
	CoolingPeriod *float64 `json:"coolingPeriod" yaml:"coolingPeriod"`
	// `CfnVolume.TieringPolicyProperty.Name`.
	Name *string `json:"name" yaml:"name"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   userAndGroupQuotasProperty := &userAndGroupQuotasProperty{
//   	id: jsii.Number(123),
//   	storageCapacityQuotaGiB: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnVolume_UserAndGroupQuotasProperty struct {
	// `CfnVolume.UserAndGroupQuotasProperty.Id`.
	Id *float64 `json:"id" yaml:"id"`
	// `CfnVolume.UserAndGroupQuotasProperty.StorageCapacityQuotaGiB`.
	StorageCapacityQuotaGiB *float64 `json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// `CfnVolume.UserAndGroupQuotasProperty.Type`.
	Type *string `json:"type" yaml:"type"`
}

// Properties for defining a `CfnVolume`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   cfnVolumeProps := &cfnVolumeProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	backupId: jsii.String("backupId"),
//   	ontapConfiguration: &ontapConfigurationProperty{
//   		junctionPath: jsii.String("junctionPath"),
//   		sizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		storageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		storageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   		// the properties below are optional
//   		securityStyle: jsii.String("securityStyle"),
//   		tieringPolicy: &tieringPolicyProperty{
//   			coolingPeriod: jsii.Number(123),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	openZfsConfiguration: &openZFSConfigurationProperty{
//   		parentVolumeId: jsii.String("parentVolumeId"),
//
//   		// the properties below are optional
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
//   		options: []*string{
//   			jsii.String("options"),
//   		},
//   		originSnapshot: &originSnapshotProperty{
//   			copyStrategy: jsii.String("copyStrategy"),
//   			snapshotArn: jsii.String("snapshotArn"),
//   		},
//   		readOnly: jsii.Boolean(false),
//   		recordSizeKiB: jsii.Number(123),
//   		storageCapacityQuotaGiB: jsii.Number(123),
//   		storageCapacityReservationGiB: jsii.Number(123),
//   		userAndGroupQuotas: []interface{}{
//   			&userAndGroupQuotasProperty{
//   				id: jsii.Number(123),
//   				storageCapacityQuotaGiB: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnVolumeProps struct {
	// `AWS::FSx::Volume.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::FSx::Volume.BackupId`.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// `AWS::FSx::Volume.OntapConfiguration`.
	OntapConfiguration interface{} `json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// `AWS::FSx::Volume.OpenZFSConfiguration`.
	OpenZfsConfiguration interface{} `json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// `AWS::FSx::Volume.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// `AWS::FSx::Volume.VolumeType`.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

// Properties that describe an existing FSx file system.
//
// Example:
//   sg := ec2.securityGroup.fromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.lustreFileSystem.fromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &fileSystemAttributes{
//   	dnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	fileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	securityGroup: sg,
//   })
//
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("Vpc"), &vpcAttributes{
//   	availabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	vpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &instanceProps{
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_T2, ec2.instanceSize_LARGE),
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   fs.connections.allowDefaultPortFrom(inst)
//
// Experimental.
type FileSystemAttributes struct {
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName *string `json:"dnsName" yaml:"dnsName"`
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// The security group of the file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// A new or imported FSx file system.
// Experimental.
type FileSystemBase interface {
	awscdk.Resource
	IFileSystem
	// The security groups/rules used to allow network connections to the file system.
	// Experimental.
	Connections() awsec2.Connections
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for FileSystemBase
type jsiiProxy_FileSystemBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IFileSystem
}

func (j *jsiiProxy_FileSystemBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFileSystemBase_Override(f FileSystemBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.FileSystemBase",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FileSystemBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.FileSystemBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FileSystemBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.FileSystemBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FileSystemBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystemBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystemBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystemBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystemBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for the FSx file system.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"import awscdk "github.com/aws/aws-cdk-go/awscdk"import kms "github.com/aws/aws-cdk-go/awscdk/aws_kms"
//
//   var key key
//   var securityGroup securityGroup
//   var vpc vpc
//   fileSystemProps := &fileSystemProps{
//   	storageCapacityGiB: jsii.Number(123),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	backupId: jsii.String("backupId"),
//   	kmsKey: key,
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   	securityGroup: securityGroup,
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
// Experimental.
type FileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	// Experimental.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB" yaml:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Experimental.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// Interface to implement FSx File Systems.
// Experimental.
type IFileSystem interface {
	awsec2.IConnectable
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId() *string
}

// The jsii proxy for IFileSystem
type jsiiProxy_IFileSystem struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_IFileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   var vpc vpc
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	lustreConfiguration: &lustreConfiguration{
//   		deploymentType: fsx.lustreDeploymentType_SCRATCH_2,
//   	},
//   	storageCapacityGiB: jsii.Number(1200),
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
// Experimental.
type LustreConfiguration struct {
	// The type of backing file system deployment used by FSx.
	// Experimental.
	DeploymentType LustreDeploymentType `json:"deploymentType" yaml:"deploymentType"`
	// The path in Amazon S3 where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same
	// Amazon S3 bucket as specified in ImportPath. If you only specify a bucket name, such as s3://import-bucket, you
	// get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is
	// overwritten on export. If you provide a custom prefix in the export path, such as
	// s3://import-bucket/[custom-optional-prefix], Amazon FSx exports the contents of your file system to that export
	// prefix in the Amazon S3 bucket.
	// Experimental.
	ExportPath *string `json:"exportPath" yaml:"exportPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// Allowed values are between 1 and 512,000.
	// Experimental.
	ImportedFileChunkSizeMiB *float64 `json:"importedFileChunkSizeMiB" yaml:"importedFileChunkSizeMiB"`
	// The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// Must be of the format "s3://{bucketName}/optional-prefix" and cannot
	// exceed 900 characters.
	// Experimental.
	ImportPath *string `json:"importPath" yaml:"importPath"`
	// Required for the PERSISTENT_1 deployment type, describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB.
	//
	// Valid values are 50, 100, 200.
	// Experimental.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// The preferred day and time to perform weekly maintenance.
	//
	// The first digit is the day of the week, starting at 1
	// for Monday, then the following are hours and minutes in the UTC time zone, 24 hour clock. For example: '2:20:30'
	// is Tuesdays at 20:30.
	// Experimental.
	WeeklyMaintenanceStartTime LustreMaintenanceTime `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

// The different kinds of file system deployments used by Lustre.
//
// Example:
//   var vpc vpc
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	lustreConfiguration: &lustreConfiguration{
//   		deploymentType: fsx.lustreDeploymentType_SCRATCH_2,
//   	},
//   	storageCapacityGiB: jsii.Number(1200),
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   })
//
// Experimental.
type LustreDeploymentType string

const (
	// Original type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	// Experimental.
	LustreDeploymentType_SCRATCH_1 LustreDeploymentType = "SCRATCH_1"
	// Newer type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	// Provides better support for spiky workloads.
	// Experimental.
	LustreDeploymentType_SCRATCH_2 LustreDeploymentType = "SCRATCH_2"
	// Long term storage.
	//
	// Data is replicated and file servers are replaced if they fail.
	// Experimental.
	LustreDeploymentType_PERSISTENT_1 LustreDeploymentType = "PERSISTENT_1"
	// Newer type of long term storage with higher throughput tiers.
	//
	// Data is replicated and file servers are replaced if they fail.
	// Experimental.
	LustreDeploymentType_PERSISTENT_2 LustreDeploymentType = "PERSISTENT_2"
)

// The FSx for Lustre File System implementation of IFileSystem.
//
// Example:
//   sg := ec2.securityGroup.fromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.lustreFileSystem.fromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &fileSystemAttributes{
//   	dnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	fileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	securityGroup: sg,
//   })
//
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("Vpc"), &vpcAttributes{
//   	availabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	vpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &instanceProps{
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_T2, ec2.instanceSize_LARGE),
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   fs.connections.allowDefaultPortFrom(inst)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
// Experimental.
type LustreFileSystem interface {
	FileSystemBase
	// The security groups/rules used to allow network connections to the file system.
	// Experimental.
	Connections() awsec2.Connections
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ID that AWS assigns to the file system.
	// Experimental.
	FileSystemId() *string
	// The mount name of the file system, generated by FSx.
	// Experimental.
	MountName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for LustreFileSystem
type jsiiProxy_LustreFileSystem struct {
	jsiiProxy_FileSystemBase
}

func (j *jsiiProxy_LustreFileSystem) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLustreFileSystem(scope constructs.Construct, id *string, props *LustreFileSystemProps) LustreFileSystem {
	_init_.Initialize()

	j := jsiiProxy_LustreFileSystem{}

	_jsii_.Create(
		"monocdk.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLustreFileSystem_Override(l LustreFileSystem, scope constructs.Construct, id *string, props *LustreFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing FSx for Lustre file system from the given properties.
// Experimental.
func LustreFileSystem_FromLustreFileSystemAttributes(scope constructs.Construct, id *string, attrs *FileSystemAttributes) IFileSystem {
	_init_.Initialize()

	var returns IFileSystem

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.LustreFileSystem",
		"fromLustreFileSystemAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.LustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LustreFileSystem_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.LustreFileSystem",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LustreFileSystem) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreFileSystem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LustreFileSystem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreFileSystem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties specific to the Lustre version of the FSx file system.
//
// Example:
//   var vpc vpc
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	lustreConfiguration: &lustreConfiguration{
//   		deploymentType: fsx.lustreDeploymentType_SCRATCH_2,
//   	},
//   	storageCapacityGiB: jsii.Number(1200),
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   })
//
// Experimental.
type LustreFileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	// Experimental.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB" yaml:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Experimental.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// Additional configuration for FSx specific to Lustre.
	// Experimental.
	LustreConfiguration *LustreConfiguration `json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The subnet that the file system will be accessible from.
	// Experimental.
	VpcSubnet awsec2.ISubnet `json:"vpcSubnet" yaml:"vpcSubnet"`
}

// Class for scheduling a weekly manitenance time.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   lustreMaintenanceTime := fsx.NewLustreMaintenanceTime(&lustreMaintenanceTimeProps{
//   	day: fsx.weekday_MONDAY,
//   	hour: jsii.Number(123),
//   	minute: jsii.Number(123),
//   })
//
// Experimental.
type LustreMaintenanceTime interface {
	// Converts a day, hour, and minute into a timestamp as used by FSx for Lustre's weeklyMaintenanceStartTime field.
	// Experimental.
	ToTimestamp() *string
}

// The jsii proxy struct for LustreMaintenanceTime
type jsiiProxy_LustreMaintenanceTime struct {
	_ byte // padding
}

// Experimental.
func NewLustreMaintenanceTime(props *LustreMaintenanceTimeProps) LustreMaintenanceTime {
	_init_.Initialize()

	j := jsiiProxy_LustreMaintenanceTime{}

	_jsii_.Create(
		"monocdk.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLustreMaintenanceTime_Override(l LustreMaintenanceTime, props *LustreMaintenanceTimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LustreMaintenanceTime) ToTimestamp() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toTimestamp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties required for setting up a weekly maintenance time.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fsx "github.com/aws/aws-cdk-go/awscdk/aws_fsx"
//   lustreMaintenanceTimeProps := &lustreMaintenanceTimeProps{
//   	day: fsx.weekday_MONDAY,
//   	hour: jsii.Number(123),
//   	minute: jsii.Number(123),
//   }
//
// Experimental.
type LustreMaintenanceTimeProps struct {
	// The day of the week for maintenance to be performed.
	// Experimental.
	Day Weekday `json:"day" yaml:"day"`
	// The hour of the day (from 0-24) for maintenance to be performed.
	// Experimental.
	Hour *float64 `json:"hour" yaml:"hour"`
	// The minute of the hour (from 0-59) for maintenance to be performed.
	// Experimental.
	Minute *float64 `json:"minute" yaml:"minute"`
}

// Enum for representing all the days of the week.
// Experimental.
type Weekday string

const (
	// Monday.
	// Experimental.
	Weekday_MONDAY Weekday = "MONDAY"
	// Tuesday.
	// Experimental.
	Weekday_TUESDAY Weekday = "TUESDAY"
	// Wednesday.
	// Experimental.
	Weekday_WEDNESDAY Weekday = "WEDNESDAY"
	// Thursday.
	// Experimental.
	Weekday_THURSDAY Weekday = "THURSDAY"
	// Friday.
	// Experimental.
	Weekday_FRIDAY Weekday = "FRIDAY"
	// Saturday.
	// Experimental.
	Weekday_SATURDAY Weekday = "SATURDAY"
	// Sunday.
	// Experimental.
	Weekday_SUNDAY Weekday = "SUNDAY"
)

