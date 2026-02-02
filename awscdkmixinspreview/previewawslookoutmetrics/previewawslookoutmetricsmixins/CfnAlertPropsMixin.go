package previewawslookoutmetricsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslookoutmetrics/previewawslookoutmetricsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > End of support notice: On Oct 9, 2025, AWS will end support for Amazon Lookout for Metrics.
//
// After Oct 9, 2025, you will no longer be able to access the Amazon Lookout for Metrics console or Amazon Lookout for Metrics resources. For more information, see [Amazon Lookout for Metrics end of support](https://docs.aws.amazon.com//blogs/machine-learning/transitioning-off-amazon-lookout-for-metrics/) .
//
// The `AWS::LookoutMetrics::Alert` type creates an alert for an anomaly detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlertPropsMixin := awscdkmixinspreview.Mixins.NewCfnAlertPropsMixin(&CfnAlertMixinProps{
//   	Action: &ActionProperty{
//   		LambdaConfiguration: &LambdaConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		SnsConfiguration: &SNSConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	AlertDescription: jsii.String("alertDescription"),
//   	AlertName: jsii.String("alertName"),
//   	AlertSensitivityThreshold: jsii.Number(123),
//   	AnomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html
//
type CfnAlertPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAlertMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlertPropsMixin
type jsiiProxy_CfnAlertPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAlertPropsMixin) Props() *CfnAlertMixinProps {
	var returns *CfnAlertMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlertPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LookoutMetrics::Alert`.
func NewCfnAlertPropsMixin(props *CfnAlertMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAlertPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlertPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlertPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LookoutMetrics::Alert`.
func NewCfnAlertPropsMixin_Override(c CfnAlertPropsMixin, props *CfnAlertMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAlertPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlertPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlertPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlertPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAlertPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

