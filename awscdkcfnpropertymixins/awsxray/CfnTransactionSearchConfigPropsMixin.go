package awsxray

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransactionSearchConfigPropsMixin := awscdkcfnpropertymixins.Aws_xray.NewCfnTransactionSearchConfigPropsMixin(&CfnTransactionSearchConfigMixinProps{
//   	IndexingPercentage: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html
//
type CfnTransactionSearchConfigPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransactionSearchConfigMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransactionSearchConfigPropsMixin
type jsiiProxy_CfnTransactionSearchConfigPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransactionSearchConfigPropsMixin) Props() *CfnTransactionSearchConfigMixinProps {
	var returns *CfnTransactionSearchConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransactionSearchConfigPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::XRay::TransactionSearchConfig`.
func NewCfnTransactionSearchConfigPropsMixin(props *CfnTransactionSearchConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransactionSearchConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransactionSearchConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransactionSearchConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_xray.CfnTransactionSearchConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::XRay::TransactionSearchConfig`.
func NewCfnTransactionSearchConfigPropsMixin_Override(c CfnTransactionSearchConfigPropsMixin, props *CfnTransactionSearchConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_xray.CfnTransactionSearchConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransactionSearchConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransactionSearchConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_xray.CfnTransactionSearchConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransactionSearchConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_xray.CfnTransactionSearchConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransactionSearchConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransactionSearchConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

