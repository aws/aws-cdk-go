package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkfirewall/previewawsnetworkfirewallmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the logging configuration to define the destinations and logging options for an firewall.
//
// You must change the logging configuration by changing one `LogDestinationConfig` setting at a time in your `LogDestinationConfigs` .
//
// You can make only one of the following changes to your logging configuration resource:
//
// - Create a new log destination object by adding a single `LogDestinationConfig` array element to `LogDestinationConfigs` .
// - Delete a log destination object by removing a single `LogDestinationConfig` array element from `LogDestinationConfigs` .
// - Change the `LogDestination` setting in a single `LogDestinationConfig` array element.
//
// You can't change the `LogDestinationType` or `LogType` in a `LogDestinationConfig` . To change these settings, delete the existing `LogDestinationConfig` object and create a new one, in two separate modifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoggingConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnLoggingConfigurationPropsMixin(&CfnLoggingConfigurationMixinProps{
//   	EnableMonitoringDashboard: jsii.Boolean(false),
//   	FirewallArn: jsii.String("firewallArn"),
//   	FirewallName: jsii.String("firewallName"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogDestinationConfigs: []interface{}{
//   			&LogDestinationConfigProperty{
//   				LogDestination: map[string]*string{
//   					"logDestinationKey": jsii.String("logDestination"),
//   				},
//   				LogDestinationType: jsii.String("logDestinationType"),
//   				LogType: jsii.String("logType"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html
//
type CfnLoggingConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoggingConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoggingConfigurationPropsMixin
type jsiiProxy_CfnLoggingConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoggingConfigurationPropsMixin) Props() *CfnLoggingConfigurationMixinProps {
	var returns *CfnLoggingConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::LoggingConfiguration`.
func NewCfnLoggingConfigurationPropsMixin(props *CfnLoggingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoggingConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoggingConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::LoggingConfiguration`.
func NewCfnLoggingConfigurationPropsMixin_Override(c CfnLoggingConfigurationPropsMixin, props *CfnLoggingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoggingConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoggingConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggingConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggingConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLoggingConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

