package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsses/previewawssesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration sets let you create groups of rules that you can apply to the emails you send using Amazon SES.
//
// For more information about using configuration sets, see [Using Amazon SES Configuration Sets](https://docs.aws.amazon.com/ses/latest/dg/using-configuration-sets.html) in the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/) .
//
// > *Required permissions:*
// >
// > To apply any of the resource options, you will need to have the corresponding AWS Identity and Access Management (IAM) SES API v2 permissions:
// >
// > - `ses:GetConfigurationSet`
// >
// > - (This permission is replacing the v1 *ses:DescribeConfigurationSet* permission which will not work with these v2 resource options.)
// > - `ses:PutConfigurationSetDeliveryOptions`
// > - `ses:PutConfigurationSetReputationOptions`
// > - `ses:PutConfigurationSetSendingOptions`
// > - `ses:PutConfigurationSetSuppressionOptions`
// > - `ses:PutConfigurationSetTrackingOptions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationSetPropsMixin(&CfnConfigurationSetMixinProps{
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		MaxDeliverySeconds: jsii.Number(123),
//   		SendingPoolName: jsii.String("sendingPoolName"),
//   		TlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	Name: jsii.String("name"),
//   	ReputationOptions: &ReputationOptionsProperty{
//   		ReputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	SendingOptions: &SendingOptionsProperty{
//   		SendingEnabled: jsii.Boolean(false),
//   	},
//   	SuppressionOptions: &SuppressionOptionsProperty{
//   		SuppressedReasons: []*string{
//   			jsii.String("suppressedReasons"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   		HttpsPolicy: jsii.String("httpsPolicy"),
//   	},
//   	VdmOptions: &VdmOptionsProperty{
//   		DashboardOptions: &DashboardOptionsProperty{
//   			EngagementMetrics: jsii.String("engagementMetrics"),
//   		},
//   		GuardianOptions: &GuardianOptionsProperty{
//   			OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html
//
type CfnConfigurationSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationSetPropsMixin
type jsiiProxy_CfnConfigurationSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationSetPropsMixin) Props() *CfnConfigurationSetMixinProps {
	var returns *CfnConfigurationSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin(props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin_Override(c CfnConfigurationSetPropsMixin, props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfigurationSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

