package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a continuous deployment policy that routes a subset of production traffic from a primary distribution to a staging distribution.
//
// After you create and update a staging distribution, you can use a continuous deployment policy to incrementally move traffic to the staging distribution. This enables you to test changes to a distribution's configuration before moving all of your production traffic to the new configuration.
//
// For more information, see [Using CloudFront continuous deployment to safely test CDN configuration changes](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContinuousDeploymentPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnContinuousDeploymentPolicyPropsMixin(&CfnContinuousDeploymentPolicyMixinProps{
//   	ContinuousDeploymentPolicyConfig: &ContinuousDeploymentPolicyConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		SingleHeaderPolicyConfig: &SingleHeaderPolicyConfigProperty{
//   			Header: jsii.String("header"),
//   			Value: jsii.String("value"),
//   		},
//   		SingleWeightPolicyConfig: &SingleWeightPolicyConfigProperty{
//   			SessionStickinessConfig: &SessionStickinessConfigProperty{
//   				IdleTtl: jsii.Number(123),
//   				MaximumTtl: jsii.Number(123),
//   			},
//   			Weight: jsii.Number(123),
//   		},
//   		StagingDistributionDnsNames: []*string{
//   			jsii.String("stagingDistributionDnsNames"),
//   		},
//   		TrafficConfig: &TrafficConfigProperty{
//   			SingleHeaderConfig: &SingleHeaderConfigProperty{
//   				Header: jsii.String("header"),
//   				Value: jsii.String("value"),
//   			},
//   			SingleWeightConfig: &SingleWeightConfigProperty{
//   				SessionStickinessConfig: &SessionStickinessConfigProperty{
//   					IdleTtl: jsii.Number(123),
//   					MaximumTtl: jsii.Number(123),
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-continuousdeploymentpolicy.html
//
type CfnContinuousDeploymentPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnContinuousDeploymentPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContinuousDeploymentPolicyPropsMixin
type jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin) Props() *CfnContinuousDeploymentPolicyMixinProps {
	var returns *CfnContinuousDeploymentPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::ContinuousDeploymentPolicy`.
func NewCfnContinuousDeploymentPolicyPropsMixin(props *CfnContinuousDeploymentPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnContinuousDeploymentPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContinuousDeploymentPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::ContinuousDeploymentPolicy`.
func NewCfnContinuousDeploymentPolicyPropsMixin_Override(c CfnContinuousDeploymentPolicyPropsMixin, props *CfnContinuousDeploymentPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnContinuousDeploymentPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContinuousDeploymentPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContinuousDeploymentPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

