package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a secret stored in the AWS Secrets Manager that can be used to authenticate with a cluster using a user name and a password.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnBatchScramSecretPropsMixin := awscdkcfnpropertymixins.Aws_msk.NewCfnBatchScramSecretPropsMixin(&CfnBatchScramSecretMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	SecretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html
//
type CfnBatchScramSecretPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBatchScramSecretMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBatchScramSecretPropsMixin
type jsiiProxy_CfnBatchScramSecretPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnBatchScramSecretPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecretPropsMixin(props *CfnBatchScramSecretMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBatchScramSecretPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBatchScramSecretPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBatchScramSecretPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecretPropsMixin_Override(c CfnBatchScramSecretPropsMixin, props *CfnBatchScramSecretMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBatchScramSecretPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBatchScramSecretPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBatchScramSecretPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

