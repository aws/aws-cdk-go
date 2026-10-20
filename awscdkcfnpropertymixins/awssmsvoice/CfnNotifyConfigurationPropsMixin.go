package awssmsvoice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A notify configuration defines the settings used to send templated messages with AWS End User Messaging SMS managed messaging, including the display name, use case, enabled channels, and enabled countries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNotifyConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_smsvoice.NewCfnNotifyConfigurationPropsMixin(&CfnNotifyConfigurationMixinProps{
//   	DefaultTemplateId: jsii.String("defaultTemplateId"),
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	DisplayName: jsii.String("displayName"),
//   	EnabledChannels: []*string{
//   		jsii.String("enabledChannels"),
//   	},
//   	EnabledCountries: []*string{
//   		jsii.String("enabledCountries"),
//   	},
//   	PoolId: jsii.String("poolId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCase: jsii.String("useCase"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-notifyconfiguration.html
//
type CfnNotifyConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNotifyConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNotifyConfigurationPropsMixin
type jsiiProxy_CfnNotifyConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNotifyConfigurationPropsMixin) Props() *CfnNotifyConfigurationMixinProps {
	var returns *CfnNotifyConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotifyConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::NotifyConfiguration`.
func NewCfnNotifyConfigurationPropsMixin(props *CfnNotifyConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNotifyConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNotifyConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNotifyConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnNotifyConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::NotifyConfiguration`.
func NewCfnNotifyConfigurationPropsMixin_Override(c CfnNotifyConfigurationPropsMixin, props *CfnNotifyConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnNotifyConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNotifyConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNotifyConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnNotifyConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotifyConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnNotifyConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotifyConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNotifyConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

