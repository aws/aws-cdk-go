package previewawssupportappmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssupportapp/previewawssupportappmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// You can use the `AWS::SupportApp::SlackChannelConfiguration` resource to specify your AWS account when you configure the AWS Support App .
//
// This resource includes the following information:
//
// - The Slack channel name and ID
// - The team ID in Slack
// - The Amazon Resource Name (ARN) of the AWS Identity and Access Management ( IAM ) role
// - Whether you want the AWS Support App to notify you when your support cases are created, updated, resolved, or reopened
// - The case severity that you want to get notified for
//
// For more information, see the following topics in the *User Guide* :
//
// - [AWS Support App in Slack](https://docs.aws.amazon.com/awssupport/latest/user/aws-support-app-for-slack.html)
// - [Creating AWS Support App in Slack resources with AWS CloudFormation](https://docs.aws.amazon.com/awssupport/latest/user/creating-resources-with-cloudformation.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSlackChannelConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSlackChannelConfigurationPropsMixin(&CfnSlackChannelConfigurationMixinProps{
//   	ChannelId: jsii.String("channelId"),
//   	ChannelName: jsii.String("channelName"),
//   	ChannelRoleArn: jsii.String("channelRoleArn"),
//   	NotifyOnAddCorrespondenceToCase: jsii.Boolean(false),
//   	NotifyOnCaseSeverity: jsii.String("notifyOnCaseSeverity"),
//   	NotifyOnCreateOrReopenCase: jsii.Boolean(false),
//   	NotifyOnResolveCase: jsii.Boolean(false),
//   	TeamId: jsii.String("teamId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html
//
type CfnSlackChannelConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSlackChannelConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSlackChannelConfigurationPropsMixin
type jsiiProxy_CfnSlackChannelConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSlackChannelConfigurationPropsMixin) Props() *CfnSlackChannelConfigurationMixinProps {
	var returns *CfnSlackChannelConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SupportApp::SlackChannelConfiguration`.
func NewCfnSlackChannelConfigurationPropsMixin(props *CfnSlackChannelConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSlackChannelConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSlackChannelConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSlackChannelConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackChannelConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SupportApp::SlackChannelConfiguration`.
func NewCfnSlackChannelConfigurationPropsMixin_Override(c CfnSlackChannelConfigurationPropsMixin, props *CfnSlackChannelConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackChannelConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSlackChannelConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSlackChannelConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackChannelConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSlackChannelConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackChannelConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

