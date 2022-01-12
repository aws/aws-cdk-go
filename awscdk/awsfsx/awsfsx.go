package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::FSx::FileSystem`.
//
// The `AWS::FSx::FileSystem` resource is an Amazon FSx resource type that creates either an Amazon FSx for Windows File Server file system or an Amazon FSx for Lustre file system.
//
// TODO: EXAMPLE
//
type CfnFileSystem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDnsName() *string
	AttrLustreMountName() *string
	AttrRootVolumeId() *string
	BackupId() *string
	SetBackupId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	FileSystemType() *string
	SetFileSystemType(val *string)
	FileSystemTypeVersion() *string
	SetFileSystemTypeVersion(val *string)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	LustreConfiguration() interface{}
	SetLustreConfiguration(val interface{})
	Node() constructs.Node
	OntapConfiguration() interface{}
	SetOntapConfiguration(val interface{})
	OpenZfsConfiguration() interface{}
	SetOpenZfsConfiguration(val interface{})
	Ref() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	Stack() awscdk.Stack
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageType() *string
	SetStorageType(val *string)
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	WindowsConfiguration() interface{}
	SetWindowsConfiguration(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
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
func CfnFileSystem_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFileSystem) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnFileSystem) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnFileSystem) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFileSystem) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFileSystem) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnFileSystem) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.
//
// TODO: EXAMPLE
//
type CfnFileSystem_AuditLogConfigurationProperty struct {
	// Sets which attempt type is logged by Amazon FSx for file and folder accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access files or folders are logged.
	// - `FAILURE_ONLY` - only failed attempts to access files or folders are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access files or folders are logged.
	// - `DISABLED` - access auditing of files and folders is turned off.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel"`
	// Sets which attempt type is logged by Amazon FSx for file share accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access file shares are logged.
	// - `FAILURE_ONLY` - only failed attempts to access file shares are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access file shares are logged.
	// - `DISABLED` - access auditing of file shares is turned off.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel"`
	// The Amazon Resource Name (ARN) for the destination of the audit logs.
	//
	// The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN.
	//
	// The name of the Amazon CloudWatch Logs log group must begin with the `/aws/fsx` prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the `aws-fsx` prefix.
	//
	// The destination ARN (either CloudWatch Logs log group or Kinesis Data Firehose delivery stream) must be in the same AWS partition, AWS Region , and AWS account as your Amazon FSx file system.
	AuditLogDestination *string `json:"auditLogDestination"`
}

// Specifies who can mount the file system and the options that can be used while mounting the file system.
//
// TODO: EXAMPLE
//
type CfnFileSystem_ClientConfigurationsProperty struct {
	// A value that specifies who can mount the file system.
	//
	// You can provide a wildcard character ( `*` ), an IP address ( `0.0.0.0` ), or a CIDR address ( `192.0.2.0/24` . By default, Amazon FSx uses the wildcard character when specifying the client.
	Clients *string `json:"clients"`
	// The options to use when mounting the file system.
	//
	// For a list of options that you can use with Network File System (NFS), see the [exports(5) - Linux man page](https://docs.aws.amazon.com/https://linux.die.net/man/5/exports) . When choosing your options, consider the following:
	//
	// - `crossmount` is used by default. If you don't specify `crossmount` when changing the client configuration, you won't be able to see or access snapshots in your file system's snapshot directory.
	// - `sync` is used by default. If you instead specify `async` , the system acknowledges writes before writing to disk. If the system crashes before the writes are finished, you lose the unwritten data.
	Options *[]*string `json:"options"`
}

// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
//
// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
//
// TODO: EXAMPLE
//
type CfnFileSystem_DiskIopsConfigurationProperty struct {
	// The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `json:"iops"`
	// Specifies whether the number of IOPS for the file system is using the system default ( `AUTOMATIC` ) or was provisioned by the customer ( `USER_PROVISIONED` ).
	Mode *string `json:"mode"`
}

// The configuration for the Amazon FSx for Lustre file system.
//
// TODO: EXAMPLE
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
	// > This parameter is not supported for file systems with the `Persistent_2` deployment type. Instead, use `CreateDataRepositoryAssociation` to create a data repository association to link your Lustre file system to a data repository.
	AutoImportPolicy *string `json:"autoImportPolicy"`
	// The number of days to retain automatic backups.
	//
	// Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is 0. Only valid for use with `PERSISTENT_1` deployment types. For more information, see [Working with backups](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html) in the *Amazon FSx for Lustre User Guide* . (Default = 0)
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// A Boolean flag indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value. Only valid for use with `PERSISTENT_1` deployment types.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily. Only valid for use with `PERSISTENT_1` deployment types.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// Sets the data compression configuration for the file system. `DataCompressionType` can have the following values:.
	//
	// - `NONE` - (Default) Data compression is turned off when the file system is created.
	// - `LZ4` - Data compression is turned on with the LZ4 algorithm.
	//
	// For more information, see [Lustre data compression](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html) in the *Amazon FSx for Lustre User Guide* .
	DataCompressionType *string `json:"dataCompressionType"`
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
	// (Default = `SCRATCH_1` )
	DeploymentType *string `json:"deploymentType"`
	// The type of drive cache used by `PERSISTENT_1` file systems that are provisioned with HDD storage devices.
	//
	// This parameter is required when storage type is HDD. Set this property to `READ` to improve the performance for frequently accessed files by caching up to 20% of the total storage capacity of the file system.
	//
	// This parameter is required when `StorageType` is set to `HDD` .
	DriveCacheType *string `json:"driveCacheType"`
	// (Optional) Available with `Scratch` and `Persistent_1` deployment types.
	//
	// Specifies the path in the Amazon S3 bucket where the root of your Amazon FSx file system is exported. The path must use the same Amazon S3 bucket as specified in ImportPath. You can provide an optional prefix to which new and changed data is to be exported from your Amazon FSx for Lustre file system. If an `ExportPath` value is not provided, Amazon FSx sets a default export path, `s3://import-bucket/FSxLustre[creation-timestamp]` . The timestamp is in UTC format, for example `s3://import-bucket/FSxLustre20181105T222312Z` .
	//
	// The Amazon S3 export bucket must be the same as the import bucket specified by `ImportPath` . If you specify only a bucket name, such as `s3://import-bucket` , you get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is overwritten on export. If you provide a custom prefix in the export path, such as `s3://import-bucket/[custom-optional-prefix]` , Amazon FSx exports the contents of your file system to that export prefix in the Amazon S3 bucket.
	//
	// > This parameter is not supported for file systems with the `Persistent_2` deployment type. Instead, use `CreateDataRepositoryAssociation` to create a data repository association to link your Lustre file system to a data repository.
	ExportPath *string `json:"exportPath"`
	// (Optional) For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	//
	// This parameter is not supported for file systems with the `Persistent_2` deployment type. Instead, use `CreateDataRepositoryAssociation` to create a data repository association to link your Lustre file system to a data repository.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize"`
	// (Optional) The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// The root of your FSx for Lustre file system will be mapped to the root of the Amazon S3 bucket you select. An example is `s3://import-bucket/optional-prefix` . If you specify a prefix after the Amazon S3 bucket name, only object keys with that prefix are loaded into the file system.
	//
	// This parameter is not supported for file systems with the `Persistent_2` deployment type. Instead, use `CreateDataRepositoryAssociation` to create a data repository association to link your Lustre file system to a data repository.
	ImportPath *string `json:"importPath"`
	// Required with `PERSISTENT_1` and `PERSISTENT_2` deployment types, provisions the amount of read and write throughput for each 1 tebibyte (TiB) of file system storage capacity, in MB/s/TiB.
	//
	// File system throughput capacity is calculated by multiplying ﬁle system storage capacity (TiB) by the `PerUnitStorageThroughput` (MB/s/TiB). For a 2.4-TiB ﬁle system, provisioning 50 MB/s/TiB of `PerUnitStorageThroughput` yields 120 MB/s of ﬁle system throughput. You pay for the amount of throughput that you provision.
	//
	// Valid values:
	//
	// - For `PERSISTENT_1` SSD storage: 50, 100, 200 MB/s/TiB.
	// - For `PERSISTENT_1` HDD storage: 12, 40 MB/s/TiB.
	// - For `PERSISTENT_2` SSD storage: 125, 250, 500, 1000 MB/s/TiB.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput"`
	// (Optional) The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone, where d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// The configuration object for mounting a file system.
//
// TODO: EXAMPLE
//
type CfnFileSystem_NfsExportsProperty struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system.
	ClientConfigurations interface{} `json:"clientConfigurations"`
}

// The configuration for this FSx for ONTAP file system.
//
// TODO: EXAMPLE
//
type CfnFileSystem_OntapConfigurationProperty struct {
	// Specifies the FSx for ONTAP file system deployment type to use in creating the file system.
	//
	// `MULTI_AZ_1` is the supported ONTAP deployment type.
	DeploymentType *string `json:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS configuration for the FSx for ONTAP file system.
	DiskIopsConfiguration interface{} `json:"diskIopsConfiguration"`
	// Specifies the IP address range in which the endpoints to access your file system will be created.
	//
	// By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange *string `json:"endpointIpAddressRange"`
	// The ONTAP administrative password for the `fsxadmin` user with which you administer your file system using the NetApp ONTAP CLI and REST API.
	FsxAdminPassword *string `json:"fsxAdminPassword"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// Specifies the virtual private cloud (VPC) route tables in which your file system's endpoints will be created.
	//
	// You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds *[]*string `json:"routeTableIds"`
	// Sets the throughput capacity for the file system that you're creating.
	//
	// Valid values are 128, 256, 512, 1024, and 2048 MBps.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// The OpenZFS configuration for the file system that's being created.
//
// TODO: EXAMPLE
//
type CfnFileSystem_OpenZFSConfigurationProperty struct {
	// Specifies the file system deployment type.
	//
	// Amazon FSx for OpenZFS supports `SINGLE_AZ_1` . `SINGLE_AZ_1` is a file system configured for a single Availability Zone (AZ) of redundancy.
	DeploymentType *string `json:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// A Boolean value indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` , and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToVolumes interface{} `json:"copyTagsToVolumes"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
	//
	// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
	DiskIopsConfiguration interface{} `json:"diskIopsConfiguration"`
	// The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS file system.
	//
	// All volumes are children of the root volume.
	RootVolumeConfiguration interface{} `json:"rootVolumeConfiguration"`
	// Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MB/s).
	//
	// Valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MB/s. You pay for additional throughput capacity that you provision.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// The configuration of an Amazon FSx for OpenZFS root volume.
//
// TODO: EXAMPLE
//
type CfnFileSystem_RootVolumeConfigurationProperty struct {
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to `false` . If it's set to `true` , all tags for the volume are copied to snapshots where the user doesn't specify tags. If this value is `true` and you specify one or more tags, only the specified tags are copied to snapshots. If you specify one or more tags when creating the snapshot, no tags are copied from the volume, regardless of this value.
	CopyTagsToSnapshots interface{} `json:"copyTagsToSnapshots"`
	// Specifies the method used to compress the data on the volume.
	//
	// Unless the compression type is specified, volumes inherit the `DataCompressionType` value of their parent volume.
	//
	// - `NONE` - Doesn't compress the data on the volume.
	// - `ZSTD` - Compresses the data in the volume using the ZStandard (ZSTD) compression algorithm. This algorithm reduces the amount of space used on your volume and has very little impact on compute resources.
	DataCompressionType *string `json:"dataCompressionType"`
	// The configuration object for mounting a file system.
	NfsExports interface{} `json:"nfsExports"`
	// A Boolean value indicating whether the volume is read-only.
	//
	// Setting this value to `true` can be useful after you have completed changes to a volume and no longer want changes to occur.
	ReadOnly interface{} `json:"readOnly"`
	// An object specifying how much storage users or groups can use on the volume.
	UserAndGroupQuotas interface{} `json:"userAndGroupQuotas"`
}

// The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an ONTAP storage virtual machine (SVM) to a self-managed (including on-premises) Microsoft Active Directory (AD) directory.
//
// For more information, see [Using Amazon FSx with your self-managed Microsoft Active Directory](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html) or [Managing SVMs](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) .
//
// TODO: EXAMPLE
//
type CfnFileSystem_SelfManagedActiveDirectoryConfigurationProperty struct {
	// A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	DnsIps *[]*string `json:"dnsIps"`
	// The fully qualified domain name of the self-managed AD directory, such as `corp.example.com` .
	DomainName *string `json:"domainName"`
	// (Optional) The name of the domain group whose members are granted administrative privileges for the file system.
	//
	// Administrative privileges include taking ownership of files and folders, setting audit controls (audit ACLs) on files and folders, and administering the file system remotely by using the FSx Remote PowerShell. The group that you specify must already exist in your domain. If you don't provide one, your AD domain's Domain Admins group is used.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup"`
	// (Optional) The fully qualified distinguished name of the organizational unit within your self-managed AD directory.
	//
	// Amazon FSx only accepts OU as the direct parent of the file system. An example is `OU=FSx,DC=yourdomain,DC=corp,DC=com` . To learn more, see [RFC 2253](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2253) . If none is provided, the FSx file system is created in the default location of your self-managed AD directory.
	//
	// > Only Organizational Unit (OU) objects can be the direct parent of the file system that you're creating.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// We strongly suggest that you follow best practices and *do not* embed passwords in your CFN templates.
	//
	// The recommended approach is to use AWS Secrets Manager to store your passwords. You can retrieve them for use in your templates using the `secretsmanager` dynamic reference. There are additional costs associated with using AWS Secrets Manager . To learn more, see [Secrets Manager secrets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) in the *AWS CloudFormation User Guide* .
	//
	// Alternatively, you can use the `NoEcho` property to obfuscate the password parameter value. For more information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) in the *AWS CloudFormation User Guide* .
	Password *string `json:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// This account must have the permission to join computers to the domain in the organizational unit provided in `OrganizationalUnitDistinguishedName` , or in the default location of your AD domain.
	UserName *string `json:"userName"`
}

// The configuration for how much storage a user or group can use on the volume.
//
// TODO: EXAMPLE
//
type CfnFileSystem_UserAndGroupQuotasProperty struct {
	// The ID of the user or group.
	Id *float64 `json:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	StorageCapacityQuotaGiB *float64 `json:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	Type *string `json:"type"`
}

// The Microsoft Windows configuration for the file system that's being created.
//
// TODO: EXAMPLE
//
type CfnFileSystem_WindowsConfigurationProperty struct {
	// Sets the throughput capacity of an Amazon FSx file system, measured in megabytes per second (MB/s), in 2 to the *n* th increments, between 2^3 (8) and 2^11 (2048).
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// The ID for an existing AWS Managed Microsoft Active Directory (AD) instance that the file system should join when it's created.
	ActiveDirectoryId *string `json:"activeDirectoryId"`
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
	Aliases *[]*string `json:"aliases"`
	// `CfnFileSystem.WindowsConfigurationProperty.AuditLogConfiguration`.
	AuditLogConfiguration interface{} `json:"auditLogConfiguration"`
	// The number of days to retain automatic backups.
	//
	// The default is to retain backups for 7 days. Setting this value to 0 disables the creation of automatic backups. The maximum retention period for backups is 90 days.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// A Boolean flag indicating whether tags for the file system should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the file system are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the file system, regardless of this value.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// The preferred time to take daily automatic backups, formatted HH:MM in the UTC time zone.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// Specifies the file system deployment type, valid values are the following:.
	//
	// - `MULTI_AZ_1` - Deploys a high availability file system that is configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability. You can only deploy a Multi-AZ file system in AWS Regions that have a minimum of three Availability Zones. Also supports HDD storage type
	// - `SINGLE_AZ_1` - (Default) Choose to deploy a file system that is configured for single AZ redundancy.
	// - `SINGLE_AZ_2` - The latest generation Single AZ file system. Specifies a file system that is configured for single AZ redundancy and supports HDD storage type.
	//
	// For more information, see [Availability and Durability: Single-AZ and Multi-AZ File Systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) .
	DeploymentType *string `json:"deploymentType"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located. For in- AWS applications, we recommend that you launch your clients in the same Availability Zone (AZ) as your preferred file server to reduce cross-AZ data transfer costs and minimize latency.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an ONTAP storage virtual machine (SVM) to a self-managed (including on-premises) Microsoft Active Directory (AD) directory.
	//
	// For more information, see [Using Amazon FSx with your self-managed Microsoft Active Directory](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html) or [Managing SVMs](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) .
	SelfManagedActiveDirectoryConfiguration interface{} `json:"selfManagedActiveDirectoryConfiguration"`
	// The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone, where d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// Properties for defining a `CfnFileSystem`.
//
// TODO: EXAMPLE
//
type CfnFileSystemProps struct {
	// The type of Amazon FSx file system, which can be `LUSTRE` , `WINDOWS` , `ONTAP` , or `OPENZFS` .
	FileSystemType *string `json:"fileSystemType"`
	// Specifies the IDs of the subnets that the file system will be accessible from.
	//
	// For Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet IDs, one for the preferred file server and one for the standby file server. You specify one of these subnets as the preferred subnet using the `WindowsConfiguration > PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID` properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the *Amazon FSx for Windows User Guide* and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the *Amazon FSx for ONTAP User Guide* .
	//
	// For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre deployment types, provide exactly one subnet ID. The file server is launched in that subnet's Availability Zone.
	SubnetIds *[]*string `json:"subnetIds"`
	// The ID of the source backup.
	//
	// Specifies the backup that you are copying.
	BackupId *string `json:"backupId"`
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
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion"`
	// The ID of the AWS Key Management Service ( AWS KMS ) key used to encrypt the file system's data for Amazon FSx for Windows File Server file systems, Amazon FSx for NetApp ONTAP file systems, and `PERSISTENT` Amazon FSx for Lustre file systems at rest.
	//
	// If this ID isn't specified, the Amazon FSx-managed key for your account is used. The scratch Amazon FSx for Lustre file systems are always encrypted at rest using the Amazon FSx-managed key for your account. For more information, see [Encrypt](https://docs.aws.amazon.com//kms/latest/APIReference/API_Encrypt.html) in the *AWS Key Management Service API Reference* .
	KmsKeyId *string `json:"kmsKeyId"`
	// The Lustre configuration for the file system being created.
	//
	// > The following parameters are not supported for file systems with the `Persistent_2` deployment type. Instead, use `CreateDataRepositoryAssociation` to create a data repository association to link your Lustre file system to a data repository.
	// >
	// > - `AutoImportPolicy`
	// > - `ExportPath`
	// > - `ImportedChunkSize`
	// > - `ImportPath`
	LustreConfiguration interface{} `json:"lustreConfiguration"`
	// The ONTAP configuration properties of the FSx for ONTAP file system that you are creating.
	OntapConfiguration interface{} `json:"ontapConfiguration"`
	// The OpenZFS configuration properties for the file system that you are creating.
	OpenZfsConfiguration interface{} `json:"openZfsConfiguration"`
	// A list of IDs specifying the security groups to apply to all network interfaces created for file system access.
	//
	// This list isn't returned in later requests to describe the file system.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Sets the storage capacity of the file system that you're creating.
	//
	// `StorageCapacity` is required if you are creating a new file system. Do not include `StorageCapacity` if you are creating a file system from a backup.
	//
	// For Lustre file systems:
	StorageCapacity *float64 `json:"storageCapacity"`
	// Sets the storage type for the file system that you're creating. Valid values are `SSD` and `HDD` .
	//
	// - Set to `SSD` to use solid state drive storage. SSD is supported on all Windows, Lustre, ONTAP, and OpenZFS deployment types.
	// - Set to `HDD` to use hard disk drive storage. HDD is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types, and on `PERSISTENT` Lustre file system deployment types.
	//
	// Default value is `SSD` . For more information, see [Storage type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the *FSx for Windows File Server User Guide* and [Multiple storage options](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html#storage-options) in the *FSx for Lustre User Guide* .
	StorageType *string `json:"storageType"`
	// The tags to associate with the file system.
	//
	// For more information, see [Tagging your Amazon EC2 resources](https://docs.aws.amazon.com//AWSEC2/latest/UserGuide/Using_Tags.html) in the *Amazon EC2 User Guide* .
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The configuration object for the Microsoft Windows file system you are creating.
	//
	// This value is required if `FileSystemType` is set to `WINDOWS` .
	WindowsConfiguration interface{} `json:"windowsConfiguration"`
}

// Properties that describe an existing FSx file system.
//
// TODO: EXAMPLE
//
type FileSystemAttributes struct {
	// The DNS name assigned to this file system.
	DnsName *string `json:"dnsName"`
	// The ID of the file system, assigned by Amazon FSx.
	FileSystemId *string `json:"fileSystemId"`
	// The security group of the file system.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
}

// A new or imported FSx file system.
type FileSystemBase interface {
	awscdk.Resource
	IFileSystem
	Connections() awsec2.Connections
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	FileSystemId() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_FileSystemBase) Node() constructs.Node {
	var returns constructs.Node
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


func NewFileSystemBase_Override(f FileSystemBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.FileSystemBase",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FileSystemBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.FileSystemBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FileSystemBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.FileSystemBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
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

// Returns a string representation of this construct.
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

// Properties for the FSx file system.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
type FileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	Vpc awsec2.IVpc `json:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	BackupId *string `json:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	KmsKey awskms.IKey `json:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Security Group to assign to this file system.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
}

// Interface to implement FSx File Systems.
type IFileSystem interface {
	awsec2.IConnectable
	// The ID of the file system, assigned by Amazon FSx.
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
type LustreConfiguration struct {
	// The type of backing file system deployment used by FSx.
	DeploymentType LustreDeploymentType `json:"deploymentType"`
	// The path in Amazon S3 where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same
	// Amazon S3 bucket as specified in ImportPath. If you only specify a bucket name, such as s3://import-bucket, you
	// get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is
	// overwritten on export. If you provide a custom prefix in the export path, such as
	// s3://import-bucket/[custom-optional-prefix], Amazon FSx exports the contents of your file system to that export
	// prefix in the Amazon S3 bucket.
	ExportPath *string `json:"exportPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// Allowed values are between 1 and 512,000.
	ImportedFileChunkSizeMiB *float64 `json:"importedFileChunkSizeMiB"`
	// The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// Must be of the format "s3://{bucketName}/optional-prefix" and cannot
	// exceed 900 characters.
	ImportPath *string `json:"importPath"`
	// Required for the PERSISTENT_1 deployment type, describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB.
	//
	// Valid values are 50, 100, 200.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput"`
	// The preferred day and time to perform weekly maintenance.
	//
	// The first digit is the day of the week, starting at 1
	// for Monday, then the following are hours and minutes in the UTC time zone, 24 hour clock. For example: '2:20:30'
	// is Tuesdays at 20:30.
	WeeklyMaintenanceStartTime LustreMaintenanceTime `json:"weeklyMaintenanceStartTime"`
}

// The different kinds of file system deployments used by Lustre.
//
// TODO: EXAMPLE
//
type LustreDeploymentType string

const (
	LustreDeploymentType_SCRATCH_1 LustreDeploymentType = "SCRATCH_1"
	LustreDeploymentType_SCRATCH_2 LustreDeploymentType = "SCRATCH_2"
	LustreDeploymentType_PERSISTENT_1 LustreDeploymentType = "PERSISTENT_1"
)

// The FSx for Lustre File System implementation of IFileSystem.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
type LustreFileSystem interface {
	FileSystemBase
	Connections() awsec2.Connections
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	FileSystemId() *string
	MountName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
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

func (j *jsiiProxy_LustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
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


func NewLustreFileSystem(scope constructs.Construct, id *string, props *LustreFileSystemProps) LustreFileSystem {
	_init_.Initialize()

	j := jsiiProxy_LustreFileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLustreFileSystem_Override(l LustreFileSystem, scope constructs.Construct, id *string, props *LustreFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing FSx for Lustre file system from the given properties.
func LustreFileSystem_FromLustreFileSystemAttributes(scope constructs.Construct, id *string, attrs *FileSystemAttributes) IFileSystem {
	_init_.Initialize()

	var returns IFileSystem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"fromLustreFileSystemAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LustreFileSystem_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
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

// Returns a string representation of this construct.
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

// Properties specific to the Lustre version of the FSx file system.
//
// TODO: EXAMPLE
//
type LustreFileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	Vpc awsec2.IVpc `json:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	BackupId *string `json:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	KmsKey awskms.IKey `json:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Security Group to assign to this file system.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// Additional configuration for FSx specific to Lustre.
	LustreConfiguration *LustreConfiguration `json:"lustreConfiguration"`
	// The subnet that the file system will be accessible from.
	VpcSubnet awsec2.ISubnet `json:"vpcSubnet"`
}

// Class for scheduling a weekly manitenance time.
//
// TODO: EXAMPLE
//
type LustreMaintenanceTime interface {
	ToTimestamp() *string
}

// The jsii proxy struct for LustreMaintenanceTime
type jsiiProxy_LustreMaintenanceTime struct {
	_ byte // padding
}

func NewLustreMaintenanceTime(props *LustreMaintenanceTimeProps) LustreMaintenanceTime {
	_init_.Initialize()

	j := jsiiProxy_LustreMaintenanceTime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLustreMaintenanceTime_Override(l LustreMaintenanceTime, props *LustreMaintenanceTimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		l,
	)
}

// Converts a day, hour, and minute into a timestamp as used by FSx for Lustre's weeklyMaintenanceStartTime field.
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
// TODO: EXAMPLE
//
type LustreMaintenanceTimeProps struct {
	// The day of the week for maintenance to be performed.
	Day Weekday `json:"day"`
	// The hour of the day (from 0-24) for maintenance to be performed.
	Hour *float64 `json:"hour"`
	// The minute of the hour (from 0-59) for maintenance to be performed.
	Minute *float64 `json:"minute"`
}

// Enum for representing all the days of the week.
type Weekday string

const (
	Weekday_MONDAY Weekday = "MONDAY"
	Weekday_TUESDAY Weekday = "TUESDAY"
	Weekday_WEDNESDAY Weekday = "WEDNESDAY"
	Weekday_THURSDAY Weekday = "THURSDAY"
	Weekday_FRIDAY Weekday = "FRIDAY"
	Weekday_SATURDAY Weekday = "SATURDAY"
	Weekday_SUNDAY Weekday = "SUNDAY"
)

