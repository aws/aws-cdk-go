package awsresiliencehubv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsresiliencehubv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a service function within a Resilience Hub service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnServiceFunctionPropsMixin := awscdkcfnpropertymixins.Aws_resiliencehubv2.NewCfnServiceFunctionPropsMixin(&CfnServiceFunctionMixinProps{
//   	Criticality: jsii.String("criticality"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceArn: jsii.String("serviceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html
//
type CfnServiceFunctionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServiceFunctionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceFunctionPropsMixin
type jsiiProxy_CfnServiceFunctionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServiceFunctionPropsMixin) Props() *CfnServiceFunctionMixinProps {
	var returns *CfnServiceFunctionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceFunctionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResilienceHubV2::ServiceFunction`.
func NewCfnServiceFunctionPropsMixin(props *CfnServiceFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServiceFunctionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceFunctionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceFunctionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResilienceHubV2::ServiceFunction`.
func NewCfnServiceFunctionPropsMixin_Override(c CfnServiceFunctionPropsMixin, props *CfnServiceFunctionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServiceFunctionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceFunctionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceFunctionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceFunctionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServiceFunctionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

