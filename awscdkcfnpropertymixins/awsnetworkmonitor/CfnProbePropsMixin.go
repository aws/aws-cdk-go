package awsnetworkmonitor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnetworkmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon CloudWatch Network Monitor probe that measures network traffic between a source subnet and a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProbePropsMixin := awscdkcfnpropertymixins.Aws_networkmonitor.NewCfnProbePropsMixin(&CfnProbeMixinProps{
//   	Destination: jsii.String("destination"),
//   	DestinationPort: jsii.Number(123),
//   	MonitorName: jsii.String("monitorName"),
//   	PacketSize: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SourceArn: jsii.String("sourceArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html
//
type CfnProbePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProbeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProbePropsMixin
type jsiiProxy_CfnProbePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProbePropsMixin) Props() *CfnProbeMixinProps {
	var returns *CfnProbeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProbePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkMonitor::Probe`.
func NewCfnProbePropsMixin(props *CfnProbeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProbePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProbePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProbePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmonitor.CfnProbePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkMonitor::Probe`.
func NewCfnProbePropsMixin_Override(c CfnProbePropsMixin, props *CfnProbeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkmonitor.CfnProbePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProbePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProbePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_networkmonitor.CfnProbePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProbePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_networkmonitor.CfnProbePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProbePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProbePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

