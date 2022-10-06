package awsefs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EFS::FileSystem`.
//
// The `AWS::EFS::FileSystem` resource creates a new, empty file system in Amazon Elastic File System ( Amazon EFS ). You must create a mount target ( [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) ) to mount your EFS file system on an Amazon EC2 or other AWS cloud compute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystemPolicy interface{}
//
//   cfnFileSystem := awscdk.Aws_efs.NewCfnFileSystem(this, jsii.String("MyCfnFileSystem"), &cfnFileSystemProps{
//   	availabilityZoneName: jsii.String("availabilityZoneName"),
//   	backupPolicy: &backupPolicyProperty{
//   		status: jsii.String("status"),
//   	},
//   	bypassPolicyLockoutSafetyCheck: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	fileSystemPolicy: fileSystemPolicy,
//   	fileSystemTags: []elasticFileSystemTagProperty{
//   		&elasticFileSystemTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	lifecyclePolicies: []interface{}{
//   		&lifecyclePolicyProperty{
//   			transitionToIa: jsii.String("transitionToIa"),
//   			transitionToPrimaryStorageClass: jsii.String("transitionToPrimaryStorageClass"),
//   		},
//   	},
//   	performanceMode: jsii.String("performanceMode"),
//   	provisionedThroughputInMibps: jsii.Number(123),
//   	throughputMode: jsii.String("throughputMode"),
//   })
//
type CfnFileSystem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the EFS file system.
	//
	// Example: `arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-0123456789abcdef8`.
	AttrArn() *string
	// The ID of the EFS file system.
	//
	// For example: `fs-abcdef0123456789a`.
	AttrFileSystemId() *string
	// Used to create a file system that uses One Zone storage classes.
	//
	// It specifies the AWS Availability Zone in which to create the file system. Use the format `us-east-1a` to specify the Availability Zone. For more information about One Zone storage classes, see [Using EFS storage classes](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide* .
	//
	// > One Zone storage classes are not available in all Availability Zones in AWS Regions where Amazon EFS is available.
	AvailabilityZoneName() *string
	SetAvailabilityZoneName(val *string)
	// Use the `BackupPolicy` to turn automatic backups on or off for the file system.
	BackupPolicy() interface{}
	SetBackupPolicy(val interface{})
	// (Optional) A boolean that specifies whether or not to bypass the `FileSystemPolicy` lockout safety check.
	//
	// The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future `PutFileSystemPolicy` requests on this file system. Set `BypassPolicyLockoutSafetyCheck` to `True` only when you intend to prevent the IAM principal that is making the request from making subsequent `PutFileSystemPolicy` requests on this file system. The default value is `False` .
	BypassPolicyLockoutSafetyCheck() interface{}
	SetBypassPolicyLockoutSafetyCheck(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A Boolean value that, if true, creates an encrypted file system.
	//
	// When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing AWS KMS key . If you don't specify a KMS key , then the default KMS key for Amazon EFS , `/aws/elasticfilesystem` , is used to protect the encrypted file system.
	Encrypted() interface{}
	SetEncrypted(val interface{})
	// The `FileSystemPolicy` for the EFS file system.
	//
	// A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using IAM to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide* .
	FileSystemPolicy() interface{}
	SetFileSystemPolicy(val interface{})
	// The ID of the AWS KMS key to be used to protect the encrypted file system.
	//
	// This parameter is only required if you want to use a nondefault KMS key . If this parameter is not specified, the default KMS key for Amazon EFS is used. This ID can be in one of the following formats:
	//
	// - Key ID - A unique identifier of the key, for example `1234abcd-12ab-34cd-56ef-1234567890ab` .
	// - ARN - An Amazon Resource Name (ARN) for the key, for example `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` .
	// - Key alias - A previously created display name for a key, for example `alias/projectKey1` .
	// - Key alias ARN - An ARN for a key alias, for example `arn:aws:kms:us-west-2:444455556666:alias/projectKey1` .
	//
	// If `KmsKeyId` is specified, the `Encrypted` parameter must be set to true.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// An array of `LifecyclePolicy` objects that define the file system's `LifecycleConfiguration` object.
	//
	// A `LifecycleConfiguration` object informs EFS lifecycle management and intelligent tiering of the following:
	//
	// - When to move files in the file system from primary storage to the IA storage class.
	// - When to move files that are in IA storage to primary storage.
	//
	// > Amazon EFS requires that each `LifecyclePolicy` object have only a single transition. This means that in a request body, `LifecyclePolicies` needs to be structured as an array of `LifecyclePolicy` objects, one object for each transition, `TransitionToIA` , `TransitionToPrimaryStorageClass` . See the example requests in the following section for more information.
	LifecyclePolicies() interface{}
	SetLifecyclePolicies(val interface{})
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
	// The tree node.
	Node() constructs.Node
	// The performance mode of the file system.
	//
	// We recommend `generalPurpose` performance mode for most file systems. File systems using the `maxIO` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created.
	//
	// > The `maxIO` mode is not supported on file systems using One Zone storage classes.
	PerformanceMode() *string
	SetPerformanceMode(val *string)
	// The throughput, measured in MiB/s, that you want to provision for a file system that you're creating.
	//
	// Valid values are 1-1024. Required if `ThroughputMode` is set to `provisioned` . The upper limit for throughput is 1024 MiB/s. To increase this limit, contact AWS Support . For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide* .
	ProvisionedThroughputInMibps() *float64
	SetProvisionedThroughputInMibps(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Use to create one or more tags associated with the file system.
	//
	// Each tag is a user-defined key-value pair. Name your file system on creation by including a `"Key":"Name","Value":"{value}"` key-value pair. Each key must be unique. For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference Guide* .
	Tags() awscdk.TagManager
	// Specifies the throughput mode for the file system, either `bursting` or `provisioned` .
	//
	// If you set `ThroughputMode` to `provisioned` , you must also set a value for `ProvisionedThroughputInMibps` . After you create the file system, you can decrease your file system's throughput in Provisioned Throughput mode or change between the throughput modes, as long as it’s been more than 24 hours since the last decrease or throughput mode change. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide* .
	//
	// Default is `bursting` .
	ThroughputMode() *string
	SetThroughputMode(val *string)
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
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

func (j *jsiiProxy_CfnFileSystem) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) AttrFileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) AvailabilityZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) BackupPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) BypassPolicyLockoutSafetyCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPolicyLockoutSafetyCheck",
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

func (j *jsiiProxy_CfnFileSystem) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) FileSystemPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemPolicy",
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

func (j *jsiiProxy_CfnFileSystem) LifecyclePolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecyclePolicies",
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

func (j *jsiiProxy_CfnFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) PerformanceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystem) ProvisionedThroughputInMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInMibps",
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

func (j *jsiiProxy_CfnFileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
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

func (j *jsiiProxy_CfnFileSystem) ThroughputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughputMode",
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


// Create a new `AWS::EFS::FileSystem`.
func NewCfnFileSystem(scope constructs.Construct, id *string, props *CfnFileSystemProps) CfnFileSystem {
	_init_.Initialize()

	if err := validateNewCfnFileSystemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_efs.CfnFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EFS::FileSystem`.
func NewCfnFileSystem_Override(c CfnFileSystem, scope constructs.Construct, id *string, props *CfnFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_efs.CfnFileSystem",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetAvailabilityZoneName(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneName",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetBackupPolicy(val interface{}) {
	if err := j.validateSetBackupPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetBypassPolicyLockoutSafetyCheck(val interface{}) {
	if err := j.validateSetBypassPolicyLockoutSafetyCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassPolicyLockoutSafetyCheck",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetFileSystemPolicy(val interface{}) {
	if err := j.validateSetFileSystemPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemPolicy",
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

func (j *jsiiProxy_CfnFileSystem)SetLifecyclePolicies(val interface{}) {
	if err := j.validateSetLifecyclePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecyclePolicies",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetPerformanceMode(val *string) {
	_jsii_.Set(
		j,
		"performanceMode",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetProvisionedThroughputInMibps(val *float64) {
	_jsii_.Set(
		j,
		"provisionedThroughputInMibps",
		val,
	)
}

func (j *jsiiProxy_CfnFileSystem)SetThroughputMode(val *string) {
	_jsii_.Set(
		j,
		"throughputMode",
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
		"aws-cdk-lib.aws_efs.CfnFileSystem",
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
		"aws-cdk-lib.aws_efs.CfnFileSystem",
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
		"aws-cdk-lib.aws_efs.CfnFileSystem",
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
		"aws-cdk-lib.aws_efs.CfnFileSystem",
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

func (c *jsiiProxy_CfnFileSystem) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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

