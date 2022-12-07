package awskms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::KMS::ReplicaKey`.
//
// The `AWS::KMS::ReplicaKey` resource specifies a multi-Region replica key that is based on a multi-Region primary key.
//
// *Multi-Region keys* are an AWS KMS feature that lets you create multiple interoperable KMS keys in different AWS Regions . Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS Region and decrypt it in a different AWS Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
//
// A multi-Region *primary key* is a fully functional symmetric encryption KMS key, HMAC KMS key, or asymmetric KMS key that is also the model for replica keys in other AWS Regions . To create a multi-Region primary key, add an [AWS::KMS::Key](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html) resource to your CloudFormation stack. Set its `MultiRegion` property to true.
//
// A multi-Region *replica key* is a fully functional KMS key that has the same key ID and key material as a multi-Region primary key, but is located in a different AWS Region of the same AWS partition. There can be multiple replicas of a primary key, but each must be in a different AWS Region .
//
// When you create a replica key in AWS CloudFormation , the replica key is created in the AWS Region represented by the endpoint you use for the request. If you try to replicate a multi-Region key into a Region in which the key type is not supported, the request will fail.
//
// > HMAC KMS keys are not supported in all AWS Regions . For a list of supported Regions, see [HMAC keys in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html#hmac-regions) in the *AWS Key Management Service Developer Guide* .
//
// A primary key and its replicas have the same key ID and key material. They also have the same key spec, key usage, key material origin, and automatic key rotation status. These properties are known as *shared properties* . If they change, AWS KMS synchronizes the change to all related multi-Region keys. All other properties of a replica key can differ, including its key policy, tags, aliases, and key state. AWS KMS does not synchronize these properties.
//
// *Regions*
//
// AWS KMS CloudFormation resources are supported in all Regions in which AWS CloudFormation is supported. However, in the  (ap-southeast-3), you cannot use a CloudFormation template to create or manage multi-Region KMS keys (primary or replica).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyPolicy interface{}
//
//   cfnReplicaKey := awscdk.Aws_kms.NewCfnReplicaKey(this, jsii.String("MyCfnReplicaKey"), &cfnReplicaKeyProps{
//   	keyPolicy: keyPolicy,
//   	primaryKeyArn: jsii.String("primaryKeyArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	pendingWindowInDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnReplicaKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the replica key, such as `arn:aws:kms:us-west-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab` .
	//
	// The key ARNs of related multi-Region keys differ only in the Region value. For information about the key ARNs of multi-Region keys, see [How multi-Region keys work](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-how-it-works) in the *AWS Key Management Service Developer Guide* .
	AttrArn() *string
	// The key ID of the replica key, such as `mrk-1234abcd12ab34cd56ef1234567890ab` .
	//
	// Related multi-Region keys have the same key ID. For information about the key IDs of multi-Region keys, see [How multi-Region keys work](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-how-it-works) in the *AWS Key Management Service Developer Guide* .
	AttrKeyId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the KMS key.
	//
	// The default value is an empty string (no description).
	//
	// The description is not a shared property of multi-Region keys. You can specify the same description or a different description for each key in a set of related multi-Region keys. AWS Key Management Service does not synchronize this property.
	Description() *string
	SetDescription(val *string)
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations.
	//
	// When `Enabled` is `true` , the *key state* of the KMS key is `Enabled` . When `Enabled` is `false` , the key state of the KMS key is `Disabled` . The default value is `true` .
	//
	// The actual key state of the replica might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html) , [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html) , or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations. Also, while the replica key is being created, its key state is `Creating` . When the process is complete, the key state of the replica key changes to `Enabled` .
	//
	// For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* .
	Enabled() interface{}
	SetEnabled(val interface{})
	// The key policy that authorizes use of the replica key.
	//
	// The key policy is not a shared property of multi-Region keys. You can specify the same key policy or a different key policy for each key in a set of related multi-Region keys. AWS KMS does not synchronize this property.
	//
	// The key policy must conform to the following rules.
	//
	// - The key policy must give the caller [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) permission on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the **AWS Key Management Service Developer Guide** .
	// - Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to AWS KMS . When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to AWS KMS . For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *AWS Identity and Access Management User Guide* .
	//
	// A key policy document can include only the following characters:
	//
	// - Printable ASCII characters from the space character ( `\ u0020` ) through the end of the ASCII character range.
	// - Printable characters in the Basic Latin and Latin-1 Supplement character set (through `\ u00FF` ).
	// - The tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ) special characters
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `32768`.
	KeyPolicy() interface{}
	SetKeyPolicy(val interface{})
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
	// Specifies the number of days in the waiting period before AWS KMS deletes a replica key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// When you remove a replica key from a CloudFormation stack, AWS KMS schedules the replica key for deletion and starts the mandatory waiting period. The `PendingWindowInDays` property determines the length of waiting period. During the waiting period, the key state of replica key is `Pending Deletion` , which prevents it from being used in cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the replica key.
	//
	// If the KMS key is a multi-Region primary key with replica keys, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	//
	// You cannot use a CloudFormation template to cancel deletion of the replica after you remove it from the stack, regardless of the waiting period. However, if you specify a replica key in your template that is based on the same primary key as the original replica key, CloudFormation creates a new replica key with the same key ID, key material, and other shared properties of the original replica key. This new replica key can decrypt ciphertext that was encrypted under the original replica key, or any related multi-Region key.
	//
	// For detailed information about deleting multi-Region keys, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide* .
	//
	// For information about the `PendingDeletion` key state, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* . For more information about deleting KMS keys, see the [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operation in the *AWS Key Management Service API Reference* and [Deleting KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// *Minimum* : 7
	//
	// *Maximum* : 30.
	PendingWindowInDays() *float64
	SetPendingWindowInDays(val *float64)
	// Specifies the multi-Region primary key to replicate.
	//
	// The primary key must be in a different AWS Region of the same AWS partition. You can create only one replica of a given primary key in each AWS Region .
	//
	// > If you change the `PrimaryKeyArn` value of a replica key, the existing replica key is scheduled for deletion and a new replica key is created based on the specified primary key. While it is scheduled for deletion, the existing replica key becomes unusable. You can cancel the scheduled deletion of the key outside of CloudFormation.
	// >
	// > However, if you inadvertently delete a replica key, you can decrypt ciphertext encrypted by that replica key by using any related multi-Region key. If necessary, you can recreate the replica in the same Region after the previous one is completely deleted. For details, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide*
	//
	// Specify the key ARN of an existing multi-Region primary key. For example, `arn:aws:kms:us-east-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab` .
	PrimaryKeyArn() *string
	SetPrimaryKeyArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Assigns one or more tags to the replica key.
	//
	// > Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Tags are not a shared property of multi-Region keys. You can specify the same tags or different tags for each key in a set of related multi-Region keys. AWS KMS does not synchronize this property.
	//
	// Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, AWS KMS replaces the current tag value with the specified one.
	//
	// When you assign tags to an AWS resource, AWS generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) .
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

// The jsii proxy struct for CfnReplicaKey
type jsiiProxy_CfnReplicaKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicaKey) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) AttrKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) KeyPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) PendingWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) PrimaryKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKey) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::KMS::ReplicaKey`.
func NewCfnReplicaKey(scope constructs.Construct, id *string, props *CfnReplicaKeyProps) CfnReplicaKey {
	_init_.Initialize()

	if err := validateNewCfnReplicaKeyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicaKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KMS::ReplicaKey`.
func NewCfnReplicaKey_Override(c CfnReplicaKey, scope constructs.Construct, id *string, props *CfnReplicaKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicaKey)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnReplicaKey)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicaKey)SetKeyPolicy(val interface{}) {
	if err := j.validateSetKeyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnReplicaKey)SetPendingWindowInDays(val *float64) {
	_jsii_.Set(
		j,
		"pendingWindowInDays",
		val,
	)
}

func (j *jsiiProxy_CfnReplicaKey)SetPrimaryKeyArn(val *string) {
	if err := j.validateSetPrimaryKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicaKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicaKey_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicaKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnReplicaKey_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
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
func CfnReplicaKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicaKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicaKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kms.CfnReplicaKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicaKey) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicaKey) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicaKey) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicaKey) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicaKey) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicaKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicaKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnReplicaKey) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnReplicaKey) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnReplicaKey) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicaKey) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicaKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnReplicaKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicaKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicaKey) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

