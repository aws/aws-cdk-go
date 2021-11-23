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
// TODO: EXAMPLE
//
type CfnFileSystem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDnsName() *string
	AttrLustreMountName() *string
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
	Node() awscdk.ConstructNode
	OntapConfiguration() interface{}
	SetOntapConfiguration(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFileSystem) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFileSystem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFileSystem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFileSystem) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnFileSystem_AuditLogConfigurationProperty struct {
	// `CfnFileSystem.AuditLogConfigurationProperty.AuditLogDestination`.
	AuditLogDestination *string `json:"auditLogDestination"`
	// `CfnFileSystem.AuditLogConfigurationProperty.FileAccessAuditLogLevel`.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel"`
	// `CfnFileSystem.AuditLogConfigurationProperty.FileShareAccessAuditLogLevel`.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel"`
}

// TODO: EXAMPLE
//
type CfnFileSystem_DiskIopsConfigurationProperty struct {
	// `CfnFileSystem.DiskIopsConfigurationProperty.Iops`.
	Iops *float64 `json:"iops"`
	// `CfnFileSystem.DiskIopsConfigurationProperty.Mode`.
	Mode *string `json:"mode"`
}

// TODO: EXAMPLE
//
type CfnFileSystem_LustreConfigurationProperty struct {
	// `CfnFileSystem.LustreConfigurationProperty.AutoImportPolicy`.
	AutoImportPolicy *string `json:"autoImportPolicy"`
	// `CfnFileSystem.LustreConfigurationProperty.AutomaticBackupRetentionDays`.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// `CfnFileSystem.LustreConfigurationProperty.CopyTagsToBackups`.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// `CfnFileSystem.LustreConfigurationProperty.DailyAutomaticBackupStartTime`.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// `CfnFileSystem.LustreConfigurationProperty.DataCompressionType`.
	DataCompressionType *string `json:"dataCompressionType"`
	// `CfnFileSystem.LustreConfigurationProperty.DeploymentType`.
	DeploymentType *string `json:"deploymentType"`
	// `CfnFileSystem.LustreConfigurationProperty.DriveCacheType`.
	DriveCacheType *string `json:"driveCacheType"`
	// `CfnFileSystem.LustreConfigurationProperty.ExportPath`.
	ExportPath *string `json:"exportPath"`
	// `CfnFileSystem.LustreConfigurationProperty.ImportedFileChunkSize`.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize"`
	// `CfnFileSystem.LustreConfigurationProperty.ImportPath`.
	ImportPath *string `json:"importPath"`
	// `CfnFileSystem.LustreConfigurationProperty.PerUnitStorageThroughput`.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput"`
	// `CfnFileSystem.LustreConfigurationProperty.WeeklyMaintenanceStartTime`.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// TODO: EXAMPLE
//
type CfnFileSystem_OntapConfigurationProperty struct {
	// `CfnFileSystem.OntapConfigurationProperty.AutomaticBackupRetentionDays`.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// `CfnFileSystem.OntapConfigurationProperty.DailyAutomaticBackupStartTime`.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// `CfnFileSystem.OntapConfigurationProperty.DeploymentType`.
	DeploymentType *string `json:"deploymentType"`
	// `CfnFileSystem.OntapConfigurationProperty.DiskIopsConfiguration`.
	DiskIopsConfiguration interface{} `json:"diskIopsConfiguration"`
	// `CfnFileSystem.OntapConfigurationProperty.EndpointIpAddressRange`.
	EndpointIpAddressRange *string `json:"endpointIpAddressRange"`
	// `CfnFileSystem.OntapConfigurationProperty.FsxAdminPassword`.
	FsxAdminPassword *string `json:"fsxAdminPassword"`
	// `CfnFileSystem.OntapConfigurationProperty.PreferredSubnetId`.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// `CfnFileSystem.OntapConfigurationProperty.RouteTableIds`.
	RouteTableIds *[]*string `json:"routeTableIds"`
	// `CfnFileSystem.OntapConfigurationProperty.ThroughputCapacity`.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// `CfnFileSystem.OntapConfigurationProperty.WeeklyMaintenanceStartTime`.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// TODO: EXAMPLE
//
type CfnFileSystem_SelfManagedActiveDirectoryConfigurationProperty struct {
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.DnsIps`.
	DnsIps *[]*string `json:"dnsIps"`
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.DomainName`.
	DomainName *string `json:"domainName"`
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.FileSystemAdministratorsGroup`.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup"`
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.OrganizationalUnitDistinguishedName`.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName"`
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.Password`.
	Password *string `json:"password"`
	// `CfnFileSystem.SelfManagedActiveDirectoryConfigurationProperty.UserName`.
	UserName *string `json:"userName"`
}

// TODO: EXAMPLE
//
type CfnFileSystem_WindowsConfigurationProperty struct {
	// `CfnFileSystem.WindowsConfigurationProperty.ActiveDirectoryId`.
	ActiveDirectoryId *string `json:"activeDirectoryId"`
	// `CfnFileSystem.WindowsConfigurationProperty.Aliases`.
	Aliases *[]*string `json:"aliases"`
	// `CfnFileSystem.WindowsConfigurationProperty.AuditLogConfiguration`.
	AuditLogConfiguration interface{} `json:"auditLogConfiguration"`
	// `CfnFileSystem.WindowsConfigurationProperty.AutomaticBackupRetentionDays`.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// `CfnFileSystem.WindowsConfigurationProperty.CopyTagsToBackups`.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// `CfnFileSystem.WindowsConfigurationProperty.DailyAutomaticBackupStartTime`.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// `CfnFileSystem.WindowsConfigurationProperty.DeploymentType`.
	DeploymentType *string `json:"deploymentType"`
	// `CfnFileSystem.WindowsConfigurationProperty.PreferredSubnetId`.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// `CfnFileSystem.WindowsConfigurationProperty.SelfManagedActiveDirectoryConfiguration`.
	SelfManagedActiveDirectoryConfiguration interface{} `json:"selfManagedActiveDirectoryConfiguration"`
	// `CfnFileSystem.WindowsConfigurationProperty.ThroughputCapacity`.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// `CfnFileSystem.WindowsConfigurationProperty.WeeklyMaintenanceStartTime`.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

// Properties for defining a `AWS::FSx::FileSystem`.
//
// TODO: EXAMPLE
//
type CfnFileSystemProps struct {
	// `AWS::FSx::FileSystem.BackupId`.
	BackupId *string `json:"backupId"`
	// `AWS::FSx::FileSystem.FileSystemType`.
	FileSystemType *string `json:"fileSystemType"`
	// `AWS::FSx::FileSystem.FileSystemTypeVersion`.
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion"`
	// `AWS::FSx::FileSystem.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::FSx::FileSystem.LustreConfiguration`.
	LustreConfiguration interface{} `json:"lustreConfiguration"`
	// `AWS::FSx::FileSystem.OntapConfiguration`.
	OntapConfiguration interface{} `json:"ontapConfiguration"`
	// `AWS::FSx::FileSystem.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `AWS::FSx::FileSystem.StorageCapacity`.
	StorageCapacity *float64 `json:"storageCapacity"`
	// `AWS::FSx::FileSystem.StorageType`.
	StorageType *string `json:"storageType"`
	// `AWS::FSx::FileSystem.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::FSx::FileSystem.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::FSx::FileSystem.WindowsConfiguration`.
	WindowsConfiguration interface{} `json:"windowsConfiguration"`
}

// Properties that describe an existing FSx file system.
//
// TODO: EXAMPLE
//
// Experimental.
type FileSystemAttributes struct {
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName *string `json:"dnsName"`
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId *string `json:"fileSystemId"`
	// The security group of the file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
}

// A new or imported FSx file system.
// Experimental.
type FileSystemBase interface {
	awscdk.Resource
	IFileSystem
	Connections() awsec2.Connections
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	FileSystemId() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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
func (f *jsiiProxy_FileSystemBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FileSystemBase) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FileSystemBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FileSystemBase) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FileSystemBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
// Experimental.
type FileSystemProps struct {
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Experimental.
	BackupId *string `json:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	// Experimental.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
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
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
// Experimental.
type LustreConfiguration struct {
	// The type of backing file system deployment used by FSx.
	// Experimental.
	DeploymentType LustreDeploymentType `json:"deploymentType"`
	// The path in Amazon S3 where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same
	// Amazon S3 bucket as specified in ImportPath. If you only specify a bucket name, such as s3://import-bucket, you
	// get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is
	// overwritten on export. If you provide a custom prefix in the export path, such as
	// s3://import-bucket/[custom-optional-prefix], Amazon FSx exports the contents of your file system to that export
	// prefix in the Amazon S3 bucket.
	// Experimental.
	ExportPath *string `json:"exportPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// Allowed values are between 1 and 512,000.
	// Experimental.
	ImportedFileChunkSizeMiB *float64 `json:"importedFileChunkSizeMiB"`
	// The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// Must be of the format "s3://{bucketName}/optional-prefix" and cannot
	// exceed 900 characters.
	// Experimental.
	ImportPath *string `json:"importPath"`
	// Required for the PERSISTENT_1 deployment type, describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB.
	//
	// Valid values are 50, 100, 200.
	// Experimental.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput"`
	// The preferred day and time to perform weekly maintenance.
	//
	// The first digit is the day of the week, starting at 1
	// for Monday, then the following are hours and minutes in the UTC time zone, 24 hour clock. For example: '2:20:30'
	// is Tuesdays at 20:30.
	// Experimental.
	WeeklyMaintenanceStartTime LustreMaintenanceTime `json:"weeklyMaintenanceStartTime"`
}

// The different kinds of file system deployments used by Lustre.
// Experimental.
type LustreDeploymentType string

const (
	LustreDeploymentType_PERSISTENT_1 LustreDeploymentType = "PERSISTENT_1"
	LustreDeploymentType_SCRATCH_1 LustreDeploymentType = "SCRATCH_1"
	LustreDeploymentType_SCRATCH_2 LustreDeploymentType = "SCRATCH_2"
)

// The FSx for Lustre File System implementation of IFileSystem.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
// Experimental.
type LustreFileSystem interface {
	FileSystemBase
	Connections() awsec2.Connections
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	FileSystemId() *string
	MountName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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
func (l *jsiiProxy_LustreFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LustreFileSystem) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LustreFileSystem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LustreFileSystem) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LustreFileSystem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type LustreFileSystemProps struct {
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Experimental.
	BackupId *string `json:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	// Experimental.
	StorageCapacityGiB *float64 `json:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Additional configuration for FSx specific to Lustre.
	// Experimental.
	LustreConfiguration *LustreConfiguration `json:"lustreConfiguration"`
	// The subnet that the file system will be accessible from.
	// Experimental.
	VpcSubnet awsec2.ISubnet `json:"vpcSubnet"`
}

// Class for scheduling a weekly manitenance time.
//
// TODO: EXAMPLE
//
// Experimental.
type LustreMaintenanceTime interface {
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

// Converts a day, hour, and minute into a timestamp as used by FSx for Lustre's weeklyMaintenanceStartTime field.
// Experimental.
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
// Experimental.
type LustreMaintenanceTimeProps struct {
	// The day of the week for maintenance to be performed.
	// Experimental.
	Day Weekday `json:"day"`
	// The hour of the day (from 0-24) for maintenance to be performed.
	// Experimental.
	Hour *float64 `json:"hour"`
	// The minute of the hour (from 0-59) for maintenance to be performed.
	// Experimental.
	Minute *float64 `json:"minute"`
}

// Enum for representing all the days of the week.
// Experimental.
type Weekday string

const (
	Weekday_FRIDAY Weekday = "FRIDAY"
	Weekday_MONDAY Weekday = "MONDAY"
	Weekday_SATURDAY Weekday = "SATURDAY"
	Weekday_SUNDAY Weekday = "SUNDAY"
	Weekday_THURSDAY Weekday = "THURSDAY"
	Weekday_TUESDAY Weekday = "TUESDAY"
	Weekday_WEDNESDAY Weekday = "WEDNESDAY"
)

