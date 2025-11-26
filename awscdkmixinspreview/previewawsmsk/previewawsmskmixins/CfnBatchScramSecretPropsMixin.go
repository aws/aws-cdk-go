package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmsk/previewawsmskmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a secret stored in the AWS Secrets Manager that can be used to authenticate with a cluster using a user name and a password.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBatchScramSecretPropsMixin := awscdkmixinspreview.Mixins.NewCfnBatchScramSecretPropsMixin(&CfnBatchScramSecretMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	SecretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html
//
type CfnBatchScramSecretPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBatchScramSecretMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBatchScramSecretPropsMixin
type jsiiProxy_CfnBatchScramSecretPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBatchScramSecretPropsMixin) Props() *CfnBatchScramSecretMixinProps {
	var returns *CfnBatchScramSecretMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecretPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecretPropsMixin(props *CfnBatchScramSecretMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBatchScramSecretPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBatchScramSecretPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBatchScramSecretPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecretPropsMixin_Override(c CfnBatchScramSecretPropsMixin, props *CfnBatchScramSecretMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBatchScramSecretPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBatchScramSecretPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBatchScramSecretPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBatchScramSecretPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecretPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

