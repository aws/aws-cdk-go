package awskms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::KMS::Key`.
//
// The `AWS::KMS::Key` resource specifies an [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys) in AWS Key Management Service . You can use this resource to create symmetric encryption KMS keys, asymmetric KMS keys for encryption or signing, and symmetric HMAC KMS keys. You can use `AWS::KMS::Key` to create [multi-Region primary keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-primary-key) of all supported types. To replicate a multi-Region key, use the `AWS::KMS::ReplicaKey` resource.
//
// You cannot use the `AWS::KMS::Key` resource to specify a KMS key with [imported key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) or a KMS key in a [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) .
//
// > AWS KMS replaced the term *customer master key (CMK)* with *AWS KMS key* and *KMS key* . The concept has not changed. To prevent breaking changes, AWS KMS is keeping some variations of this term.
//
// You can use symmetric encryption KMS keys to encrypt and decrypt small amounts of data, but they are more commonly used to generate data keys and data key pairs. You can also use a symmetric encryption KMS key to encrypt data stored in AWS services that are [integrated with AWS KMS](https://docs.aws.amazon.com//kms/features/#AWS_Service_Integration) . For more information, see [Symmetric encryption KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#symmetric-cmks) in the *AWS Key Management Service Developer Guide* .
//
// You can use asymmetric KMS keys to encrypt and decrypt data or sign messages and verify signatures. To create an asymmetric key, you must specify an asymmetric `KeySpec` value and a `KeyUsage` value. For details, see [Asymmetric keys in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
//
// You can use HMAC KMS keys (which are also symmetric keys) to generate and verify hash-based message authentication codes. To create an HMAC key, you must specify an HMAC `KeySpec` value and a `KeyUsage` value of `GENERATE_VERIFY_MAC` . For details, see [HMAC keys in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html) in the *AWS Key Management Service Developer Guide* .
//
// You can also create symmetric encryption, asymmetric, and HMAC multi-Region primary keys. To create a multi-Region primary key, set the `MultiRegion` property to `true` . For information about multi-Region keys, see [Multi-Region keys in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
//
// > If you change the value of the `KeyUsage` , `KeySpec` , or `MultiRegion` property on an existing KMS key, the existing KMS key is [scheduled for deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) and a new KMS key is created with the specified value.
// >
// > While scheduled for deletion, the existing KMS key becomes unusable. If you don't [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the existing KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted.
//
// *Regions*
//
// AWS KMS CloudFormation resources are supported in all Regions in which AWS CloudFormation is supported. However, in the  (ap-southeast-3), you cannot use a CloudFormation template to create or manage asymmetric KMS keys or multi-Region KMS keys (primary or replica).
//
// Example:
//   var cfnTemplate cfnInclude
//
//   cfnKey := cfnTemplate.getResource(jsii.String("Key")).(cfnKey)
//   key := kms.key.fromCfnKey(cfnKey)
//
type CfnKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the KMS key, such as `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` .
	//
	// For information about the key ARN of a KMS key, see [Key ARN](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN) in the *AWS Key Management Service Developer Guide* .
	AttrArn() *string
	// The key ID of the KMS key, such as `1234abcd-12ab-34cd-56ef-1234567890ab` .
	//
	// For information about the key ID of a KMS key, see [Key ID](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id) in the *AWS Key Management Service Developer Guide* .
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
	// Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use.
	Description() *string
	SetDescription(val *string)
	// Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations.
	//
	// When `Enabled` is `true` , the *key state* of the KMS key is `Enabled` . When `Enabled` is `false` , the key state of the KMS key is `Disabled` . The default value is `true` .
	//
	// The actual key state of the KMS key might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html) , [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html) , or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations.
	//
	// For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* .
	Enabled() interface{}
	SetEnabled(val interface{})
	// Enables automatic rotation of the key material for the specified KMS key.
	//
	// By default, automatic key rotation is not enabled.
	//
	// AWS KMS supports automatic rotation only for symmetric encryption KMS keys ( `KeySpec` = `SYMMETRIC_DEFAULT` ). For asymmetric KMS keys and HMAC KMS keys, omit the `EnableKeyRotation` property or set it to `false` .
	//
	// To enable automatic key rotation of the key material for a multi-Region KMS key, set `EnableKeyRotation` to `true` on the primary key (created by using `AWS::KMS::Key` ). AWS KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate) in the *AWS Key Management Service Developer Guide* .
	//
	// When you enable automatic rotation, AWS KMS automatically creates new key material for the KMS key one year after the enable date and every year thereafter. AWS KMS retains all key material until you delete the KMS key. For detailed information about automatic key rotation, see [Rotating KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) in the *AWS Key Management Service Developer Guide* .
	EnableKeyRotation() interface{}
	SetEnableKeyRotation(val interface{})
	// The key policy that authorizes use of the KMS key. The key policy must conform to the following rules.
	//
	// - The key policy must allow the caller to make a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the **AWS Key Management Service Developer Guide** .
	// - Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to AWS KMS . When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to AWS KMS . For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *AWS Identity and Access Management User Guide* .
	//
	// If you are unsure of which policy to use, consider the *default key policy* . This is the key policy that AWS KMS applies to KMS keys that are created by using the CreateKey API with no specified key policy. It gives the AWS account that owns the key permission to perform all operations on the key. It also allows you write IAM policies to authorize access to the key. For details, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the *AWS Key Management Service Developer Guide* .
	//
	// A key policy document can include only the following characters:
	//
	// - Printable ASCII characters
	// - Printable characters in the Basic Latin and Latin-1 Supplement character set
	// - The tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ) special characters
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `32768`.
	KeyPolicy() interface{}
	SetKeyPolicy(val interface{})
	// Specifies the type of KMS key to create.
	//
	// The default value, `SYMMETRIC_DEFAULT` , creates a KMS key with a 256-bit symmetric key for encryption and decryption. You can't change the `KeySpec` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *AWS Key Management Service Developer Guide* .
	//
	// The `KeySpec` property determines the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see [AWS KMS condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms) in the *AWS Key Management Service Developer Guide* .
	//
	// > If you change the `KeySpec` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `KeySpec` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted. > [AWS services that are integrated with AWS KMS](https://docs.aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support encryption with asymmetric KMS keys. For help determining whether a KMS key is asymmetric, see [Identifying asymmetric KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in the *AWS Key Management Service Developer Guide* .
	//
	// AWS KMS supports the following key specs for KMS keys:
	//
	// - Symmetric encryption key (default)
	//
	// - `SYMMETRIC_DEFAULT` (AES-256-GCM)
	// - HMAC keys (symmetric)
	//
	// - `HMAC_224`
	// - `HMAC_256`
	// - `HMAC_384`
	// - `HMAC_512`
	// - Asymmetric RSA key pairs
	//
	// - `RSA_2048`
	// - `RSA_3072`
	// - `RSA_4096`
	// - Asymmetric NIST-recommended elliptic curve key pairs
	//
	// - `ECC_NIST_P256` (secp256r1)
	// - `ECC_NIST_P384` (secp384r1)
	// - `ECC_NIST_P521` (secp521r1)
	// - Other asymmetric elliptic curve key pairs
	//
	// - `ECC_SECG_P256K1` (secp256k1), commonly used for cryptocurrencies.
	KeySpec() *string
	SetKeySpec(val *string)
	// Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is `ENCRYPT_DECRYPT` . This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the `KeyUsage` value after the KMS key is created.
	//
	// > If you change the `KeyUsage` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `KeyUsage` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted.
	//
	// Select only one valid value.
	//
	// - For symmetric encryption KMS keys, omit the property or specify `ENCRYPT_DECRYPT` .
	// - For asymmetric KMS keys with RSA key material, specify `ENCRYPT_DECRYPT` or `SIGN_VERIFY` .
	// - For asymmetric KMS keys with ECC key material, specify `SIGN_VERIFY` .
	// - For HMAC KMS keys, specify `GENERATE_VERIFY_MAC` .
	KeyUsage() *string
	SetKeyUsage(val *string)
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
	// Creates a multi-Region primary key that you can replicate in other AWS Regions .
	//
	// You can't change the `MultiRegion` value after the KMS key is created.
	//
	// > If you change the `MultiRegion` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `Multi-Region` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted.
	//
	// For a multi-Region key, set to this property to `true` . For a single-Region key, omit this property or set it to `false` . The default value is `false` .
	//
	// *Multi-Region keys* are an AWS KMS feature that lets you create multiple interoperable KMS keys in different AWS Regions . Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS Region and decrypt it in a different AWS Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You can create a symmetric encryption, HMAC, or asymmetric multi-Region KMS key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.
	//
	// To create a replica of this primary key in a different AWS Region , create an [AWS::KMS::ReplicaKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html) resource in a CloudFormation stack in the replica Region. Specify the key ARN of this primary key.
	MultiRegion() interface{}
	SetMultiRegion(val interface{})
	// The tree node.
	Node() constructs.Node
	// Specifies the number of days in the waiting period before AWS KMS deletes a KMS key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// When you remove a KMS key from a CloudFormation stack, AWS KMS schedules the KMS key for deletion and starts the mandatory waiting period. The `PendingWindowInDays` property determines the length of waiting period. During the waiting period, the key state of KMS key is `Pending Deletion` or `Pending Replica Deletion` , which prevents the KMS key from being used in cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the KMS key.
	//
	// AWS KMS will not delete a [multi-Region primary key](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) that has replica keys. If you remove a multi-Region primary key from a CloudFormation stack, its key state changes to `PendingReplicaDeletion` so it cannot be replicated or used in cryptographic operations. This state can persist indefinitely. When the last of its replica keys is deleted, the key state of the primary key changes to `PendingDeletion` and the waiting period specified by `PendingWindowInDays` begins. When this waiting period expires, AWS KMS deletes the primary key. For details, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You cannot use a CloudFormation template to cancel deletion of the KMS key after you remove it from the stack, regardless of the waiting period. If you specify a KMS key in your template, even one with the same name, CloudFormation creates a new KMS key. To cancel deletion of a KMS key, use the AWS KMS console or the [CancelKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_CancelKeyDeletion.html) operation.
	//
	// For information about the `Pending Deletion` and `Pending Replica Deletion` key states, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* . For more information about deleting KMS keys, see the [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operation in the *AWS Key Management Service API Reference* and [Deleting KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// *Minimum* : 7
	//
	// *Maximum* : 30.
	PendingWindowInDays() *float64
	SetPendingWindowInDays(val *float64)
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
	// For information about tags in AWS KMS , see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) in the *AWS Key Management Service Developer Guide* . For information about tags in CloudFormation, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnKey
type jsiiProxy_CfnKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnKey) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) AttrKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) EnableKeyRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKeyRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) KeyPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) KeySpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) KeyUsage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) MultiRegion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) PendingWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKey) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::KMS::Key`.
func NewCfnKey(scope constructs.Construct, id *string, props *CfnKeyProps) CfnKey {
	_init_.Initialize()

	if err := validateNewCfnKeyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.CfnKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KMS::Key`.
func NewCfnKey_Override(c CfnKey, scope constructs.Construct, id *string, props *CfnKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.CfnKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnKey)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetEnableKeyRotation(val interface{}) {
	if err := j.validateSetEnableKeyRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKeyRotation",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetKeyPolicy(val interface{}) {
	if err := j.validateSetKeyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetKeySpec(val *string) {
	_jsii_.Set(
		j,
		"keySpec",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetKeyUsage(val *string) {
	_jsii_.Set(
		j,
		"keyUsage",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetMultiRegion(val interface{}) {
	if err := j.validateSetMultiRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiRegion",
		val,
	)
}

func (j *jsiiProxy_CfnKey)SetPendingWindowInDays(val *float64) {
	_jsii_.Set(
		j,
		"pendingWindowInDays",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKey_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnKey_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnKey",
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
func CfnKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.CfnKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kms.CfnKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKey) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnKey) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnKey) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnKey) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnKey) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnKey) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnKey) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnKey) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnKey) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKey) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

