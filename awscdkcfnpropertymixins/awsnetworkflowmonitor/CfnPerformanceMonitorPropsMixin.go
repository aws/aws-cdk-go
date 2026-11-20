package awsnetworkflowmonitor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnetworkflowmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a performance monitor to monitor network performance between a set of targets for a workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPerformanceMonitorPropsMixin := awscdkcfnpropertymixins.Aws_networkflowmonitor.NewCfnPerformanceMonitorPropsMixin(&CfnPerformanceMonitorMixinProps{
//   	MonitorName: jsii.String("monitorName"),
//   	ScopeArn: jsii.String("scopeArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&PerformanceMonitorTargetProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html
//
type CfnPerformanceMonitorPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPerformanceMonitorMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPerformanceMonitorPropsMixin
type jsiiProxy_CfnPerformanceMonitorPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPerformanceMonitorPropsMixin) Props() *CfnPerformanceMonitorMixinProps {
	var returns *CfnPerformanceMonitorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPerformanceMonitorPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFlowMonitor::PerformanceMonitor`.
func NewCfnPerformanceMonitorPropsMixin(props *CfnPerformanceMonitorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPerformanceMonitorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPerformanceMonitorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPerformanceMonitorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkflowmonitor.CfnPerformanceMonitorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFlowMonitor::PerformanceMonitor`.
func NewCfnPerformanceMonitorPropsMixin_Override(c CfnPerformanceMonitorPropsMixin, props *CfnPerformanceMonitorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkflowmonitor.CfnPerformanceMonitorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPerformanceMonitorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPerformanceMonitorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_networkflowmonitor.CfnPerformanceMonitorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPerformanceMonitorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_networkflowmonitor.CfnPerformanceMonitorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPerformanceMonitorPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPerformanceMonitorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

