package previewawsroute53recoverycontrolmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53recoverycontrol/previewawsroute53recoverycontrolmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a routing control in Amazon Route 53 Application Recovery Controller.
//
// Routing control states are maintained on the highly reliable cluster data plane.
//
// To get or update the state of the routing control, you must specify a cluster endpoint, which is an endpoint URL and an AWS Region. For more information, see [Code examples](https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples.html) in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutingControlPropsMixin := awscdkmixinspreview.Mixins.NewCfnRoutingControlPropsMixin(&CfnRoutingControlMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html
//
type CfnRoutingControlPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRoutingControlMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRoutingControlPropsMixin
type jsiiProxy_CfnRoutingControlPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRoutingControlPropsMixin) Props() *CfnRoutingControlMixinProps {
	var returns *CfnRoutingControlMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoutingControlPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53RecoveryControl::RoutingControl`.
func NewCfnRoutingControlPropsMixin(props *CfnRoutingControlMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRoutingControlPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRoutingControlPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRoutingControlPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnRoutingControlPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53RecoveryControl::RoutingControl`.
func NewCfnRoutingControlPropsMixin_Override(c CfnRoutingControlPropsMixin, props *CfnRoutingControlMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnRoutingControlPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRoutingControlPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRoutingControlPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnRoutingControlPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoutingControlPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnRoutingControlPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRoutingControlPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRoutingControlPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

