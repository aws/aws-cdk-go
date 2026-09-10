package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An application status check monitors an HTTP or HTTPS endpoint on Amazon EC2 instances and reports application-layer health.
//
// Configure a health check with a Protocol, Port, and Path, and then associate it with EC2 instances via AWS::EC2::ApplicationStatusCheckInstanceAssociation or AWS::EC2::ApplicationStatusCheckTagAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnApplicationStatusCheckPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnApplicationStatusCheckPropsMixin(&CfnApplicationStatusCheckMixinProps{
//   	Aggregation: jsii.String("aggregation"),
//   	DeviceIndex: jsii.Number(123),
//   	FailureThreshold: jsii.Number(123),
//   	HealthCheckPaths: []interface{}{
//   		&HealthCheckPathProperty{
//   			Destinations: []interface{}{
//   				&HealthCheckPathDestinationProperty{
//   					SecurityGroupId: jsii.String("securityGroupId"),
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			Source: &HealthCheckPathSourceProperty{
//   				SecurityGroupId: jsii.String("securityGroupId"),
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   	},
//   	InitializationGracePeriodSeconds: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IpScope: jsii.String("ipScope"),
//   	IpVersion: jsii.String("ipVersion"),
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	StatusCodeMatcher: jsii.String("statusCodeMatcher"),
//   	SuccessThreshold: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-applicationstatuscheck.html
//
type CfnApplicationStatusCheckPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnApplicationStatusCheckMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationStatusCheckPropsMixin
type jsiiProxy_CfnApplicationStatusCheckPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnApplicationStatusCheckPropsMixin) Props() *CfnApplicationStatusCheckMixinProps {
	var returns *CfnApplicationStatusCheckMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationStatusCheckPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::ApplicationStatusCheck`.
func NewCfnApplicationStatusCheckPropsMixin(props *CfnApplicationStatusCheckMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnApplicationStatusCheckPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationStatusCheckPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationStatusCheckPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnApplicationStatusCheckPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::ApplicationStatusCheck`.
func NewCfnApplicationStatusCheckPropsMixin_Override(c CfnApplicationStatusCheckPropsMixin, props *CfnApplicationStatusCheckMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnApplicationStatusCheckPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnApplicationStatusCheckPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationStatusCheckPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnApplicationStatusCheckPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationStatusCheckPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnApplicationStatusCheckPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationStatusCheckPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApplicationStatusCheckPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

