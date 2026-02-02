package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a key pair for use with an Amazon Elastic Compute Cloud instance as follows:.
//
// - To import an existing key pair, include the `PublicKeyMaterial` property.
// - To create a new key pair, omit the `PublicKeyMaterial` property.
//
// When you import an existing key pair, you specify the public key material for the key. We assume that you have the private key material for the key. AWS CloudFormation does not create or return the private key material when you import a key pair.
//
// When you create a new key pair, the private key is saved to AWS Systems Manager Parameter Store, using a parameter with the following name: `/ec2/keypair/{key_pair_id}` . For more information about retrieving private key, and the required permissions, see [Create a key pair using CloudFormation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-key-pairs.html#create-key-pair-cloudformation) in the *Amazon EC2 User Guide* .
//
// When CloudFormation deletes a key pair that was created or imported by a stack, it also deletes the parameter that was used to store the private key material in Parameter Store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKeyPairPropsMixin := awscdkmixinspreview.Mixins.NewCfnKeyPairPropsMixin(&CfnKeyPairMixinProps{
//   	KeyFormat: jsii.String("keyFormat"),
//   	KeyName: jsii.String("keyName"),
//   	KeyType: jsii.String("keyType"),
//   	PublicKeyMaterial: jsii.String("publicKeyMaterial"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-keypair.html
//
type CfnKeyPairPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnKeyPairMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKeyPairPropsMixin
type jsiiProxy_CfnKeyPairPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnKeyPairPropsMixin) Props() *CfnKeyPairMixinProps {
	var returns *CfnKeyPairMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyPairPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::KeyPair`.
func NewCfnKeyPairPropsMixin(props *CfnKeyPairMixinProps, options *mixins.CfnPropertyMixinOptions) CfnKeyPairPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKeyPairPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKeyPairPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnKeyPairPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::KeyPair`.
func NewCfnKeyPairPropsMixin_Override(c CfnKeyPairPropsMixin, props *CfnKeyPairMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnKeyPairPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnKeyPairPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKeyPairPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnKeyPairPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeyPairPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnKeyPairPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKeyPairPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnKeyPairPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

