package awsdataexchange

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdataexchange/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DataExchange::EntitledDataSets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEntitledDataSetsPropsMixin := awscdkcfnpropertymixins.Aws_dataexchange.NewCfnEntitledDataSetsPropsMixin(&CfnEntitledDataSetsMixinProps{
//   	AssetType: jsii.String("assetType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-entitleddatasets.html
//
type CfnEntitledDataSetsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEntitledDataSetsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEntitledDataSetsPropsMixin
type jsiiProxy_CfnEntitledDataSetsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEntitledDataSetsPropsMixin) Props() *CfnEntitledDataSetsMixinProps {
	var returns *CfnEntitledDataSetsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntitledDataSetsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataExchange::EntitledDataSets`.
func NewCfnEntitledDataSetsPropsMixin(props *CfnEntitledDataSetsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEntitledDataSetsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEntitledDataSetsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEntitledDataSetsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEntitledDataSetsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataExchange::EntitledDataSets`.
func NewCfnEntitledDataSetsPropsMixin_Override(c CfnEntitledDataSetsPropsMixin, props *CfnEntitledDataSetsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEntitledDataSetsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEntitledDataSetsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEntitledDataSetsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEntitledDataSetsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEntitledDataSetsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEntitledDataSetsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEntitledDataSetsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEntitledDataSetsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

