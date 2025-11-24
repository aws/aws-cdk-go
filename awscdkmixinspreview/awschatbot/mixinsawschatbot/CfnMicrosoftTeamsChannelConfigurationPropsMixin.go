package mixinsawschatbot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awschatbot/mixinsawschatbot/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// > AWS Chatbot is now  . [Learn more](https://docs.aws.amazon.com//chatbot/latest/adminguide/service-rename.html) >  > `Type` attribute values remain unchanged.
//
// The `AWS::Chatbot::MicrosoftTeamsChannelConfiguration` resource configures a Microsoft Teams channel to allow users to use  with CloudFormation templates.
//
// This resource requires some setup to be done in the  in chat applications console. To provide the required Microsoft Teams team and tenant IDs, you must perform the initial authorization flow with Microsoft Teams in the  in chat applications console, then copy and paste the IDs from the console. For more details, see steps 1-3 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *in chat applications Administrator Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMicrosoftTeamsChannelConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnMicrosoftTeamsChannelConfigurationPropsMixin(&CfnMicrosoftTeamsChannelConfigurationMixinProps{
//   	ConfigurationName: jsii.String("configurationName"),
//   	CustomizationResourceArns: []*string{
//   		jsii.String("customizationResourceArns"),
//   	},
//   	GuardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	SnsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TeamId: jsii.String("teamId"),
//   	TeamsChannelId: jsii.String("teamsChannelId"),
//   	TeamsChannelName: jsii.String("teamsChannelName"),
//   	TeamsTenantId: jsii.String("teamsTenantId"),
//   	UserRoleRequired: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html
//
type CfnMicrosoftTeamsChannelConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMicrosoftTeamsChannelConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMicrosoftTeamsChannelConfigurationPropsMixin
type jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin) Props() *CfnMicrosoftTeamsChannelConfigurationMixinProps {
	var returns *CfnMicrosoftTeamsChannelConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`.
func NewCfnMicrosoftTeamsChannelConfigurationPropsMixin(props *CfnMicrosoftTeamsChannelConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMicrosoftTeamsChannelConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMicrosoftTeamsChannelConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnMicrosoftTeamsChannelConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`.
func NewCfnMicrosoftTeamsChannelConfigurationPropsMixin_Override(c CfnMicrosoftTeamsChannelConfigurationPropsMixin, props *CfnMicrosoftTeamsChannelConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnMicrosoftTeamsChannelConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMicrosoftTeamsChannelConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMicrosoftTeamsChannelConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnMicrosoftTeamsChannelConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMicrosoftTeamsChannelConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnMicrosoftTeamsChannelConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

