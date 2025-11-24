package mixinsawsapplicationsignals

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsapplicationsignals/mixinsawsapplicationsignals/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// > If you have existing `AWS::ApplicationSignals::Discovery` resources that were created prior to the Application Map release, you will need to delete and recreate these resources in your account to enable Application Map.
//
// Enables this AWS account to be able to use CloudWatch Application Signals by creating the `AWSServiceRoleForCloudWatchApplicationSignals` service-linked role. This service-linked role has the following permissions:
//
// - `xray:GetServiceGraph`
// - `logs:StartQuery`
// - `logs:GetQueryResults`
// - `cloudwatch:GetMetricData`
// - `cloudwatch:ListMetrics`
// - `tag:GetResources`
// - `autoscaling:DescribeAutoScalingGroups`
//
// A service-linked CloudTrail event channel is created to process CloudTrail events and return change event information. This includes last deployment time, userName, eventName, and other event metadata.
//
// After completing this step, you still need to instrument your Java and Python applications to send data to Application Signals. For more information, see [Enabling Application Signals](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDiscoveryPropsMixin := awscdkmixinspreview.Mixins.NewCfnDiscoveryPropsMixin(&CfnDiscoveryMixinProps{
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-discovery.html
//
type CfnDiscoveryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDiscoveryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDiscoveryPropsMixin
type jsiiProxy_CfnDiscoveryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDiscoveryPropsMixin) Props() *CfnDiscoveryMixinProps {
	var returns *CfnDiscoveryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDiscoveryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApplicationSignals::Discovery`.
func NewCfnDiscoveryPropsMixin(props *CfnDiscoveryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDiscoveryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDiscoveryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDiscoveryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApplicationSignals::Discovery`.
func NewCfnDiscoveryPropsMixin_Override(c CfnDiscoveryPropsMixin, props *CfnDiscoveryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDiscoveryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDiscoveryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDiscoveryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDiscoveryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDiscoveryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

