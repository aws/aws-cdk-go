package awskms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::KMS::ReplicaKey` resource specifies a multi-Region replica key that is based on a multi-Region primary key.
//
// *Multi-Region keys* are an AWS  feature that lets you create multiple interoperable KMS keys in different AWS Regions . Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS Region and decrypt it in a different AWS Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
//
// A multi-Region *primary key* is a fully functional symmetric encryption KMS key, HMAC KMS key, or asymmetric KMS key that is also the model for replica keys in other AWS Regions . To create a multi-Region primary key, add an [AWS::KMS::Key](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html) resource to your CloudFormation stack. Set its `MultiRegion` property to true.
//
// A multi-Region *replica key* is a fully functional KMS key that has the same key ID and key material as a multi-Region primary key, but is located in a different AWS Region of the same AWS partition. There can be multiple replicas of a primary key, but each must be in a different AWS Region .
//
// When you create a replica key in CloudFormation , the replica key is created in the AWS Region represented by the endpoint you use for the request. If you try to replicate a multi-Region key into a Region in which the key type is not supported, the request will fail.
//
// A primary key and its replicas have the same key ID and key material. They also have the same key spec, key usage, key material origin, and automatic key rotation status. These properties are known as *shared properties* . If they change, AWS  synchronizes the change to all related multi-Region keys. All other properties of a replica key can differ, including its key policy, tags, aliases, and key state. AWS  does not synchronize these properties.
//
// *Regions*
//
// AWS  CloudFormation resources are available in all AWS Regions in which AWS  and CloudFormation are supported. You can use the `AWS::KMS::ReplicaKey` resource to create replica keys in all Regions that support multi-Region KMS keys. For details, see [Multi-Region keys in AWS](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyPolicy interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnReplicaKeyPropsMixin := awscdkcfnpropertymixins.Aws_kms.NewCfnReplicaKeyPropsMixin(&CfnReplicaKeyMixinProps{
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	KeyPolicy: keyPolicy,
//   	PendingWindowInDays: jsii.Number(123),
//   	PrimaryKeyArn: jsii.String("primaryKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html
//
type CfnReplicaKeyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnReplicaKeyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicaKeyPropsMixin
type jsiiProxy_CfnReplicaKeyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnReplicaKeyPropsMixin) Props() *CfnReplicaKeyMixinProps {
	var returns *CfnReplicaKeyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicaKeyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KMS::ReplicaKey`.
func NewCfnReplicaKeyPropsMixin(props *CfnReplicaKeyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnReplicaKeyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicaKeyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicaKeyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kms.CfnReplicaKeyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KMS::ReplicaKey`.
func NewCfnReplicaKeyPropsMixin_Override(c CfnReplicaKeyPropsMixin, props *CfnReplicaKeyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kms.CfnReplicaKeyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnReplicaKeyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicaKeyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kms.CfnReplicaKeyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicaKeyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kms.CfnReplicaKeyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicaKeyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnReplicaKeyPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

