package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsvpclattice/previewawsvpclatticemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a target group.
//
// A target group is a collection of targets, or compute resources, that run your application or service. A target group can only be used by a single service.
//
// For more information, see [Target groups](https://docs.aws.amazon.com/vpc-lattice/latest/ug/target-groups.html) in the *Amazon VPC Lattice User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTargetGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnTargetGroupPropsMixin(&CfnTargetGroupMixinProps{
//   	Config: &TargetGroupConfigProperty{
//   		HealthCheck: &HealthCheckConfigProperty{
//   			Enabled: jsii.Boolean(false),
//   			HealthCheckIntervalSeconds: jsii.Number(123),
//   			HealthCheckTimeoutSeconds: jsii.Number(123),
//   			HealthyThresholdCount: jsii.Number(123),
//   			Matcher: &MatcherProperty{
//   				HttpCode: jsii.String("httpCode"),
//   			},
//   			Path: jsii.String("path"),
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			ProtocolVersion: jsii.String("protocolVersion"),
//   			UnhealthyThresholdCount: jsii.Number(123),
//   		},
//   		IpAddressType: jsii.String("ipAddressType"),
//   		LambdaEventStructureVersion: jsii.String("lambdaEventStructureVersion"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   		VpcIdentifier: jsii.String("vpcIdentifier"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Id: jsii.String("id"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html
//
type CfnTargetGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTargetGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTargetGroupPropsMixin
type jsiiProxy_CfnTargetGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTargetGroupPropsMixin) Props() *CfnTargetGroupMixinProps {
	var returns *CfnTargetGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VpcLattice::TargetGroup`.
func NewCfnTargetGroupPropsMixin(props *CfnTargetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTargetGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTargetGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTargetGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnTargetGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VpcLattice::TargetGroup`.
func NewCfnTargetGroupPropsMixin_Override(c CfnTargetGroupPropsMixin, props *CfnTargetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnTargetGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTargetGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTargetGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnTargetGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTargetGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnTargetGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTargetGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTargetGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

