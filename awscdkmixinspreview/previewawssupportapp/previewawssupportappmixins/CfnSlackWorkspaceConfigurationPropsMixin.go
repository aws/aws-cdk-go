package previewawssupportappmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssupportapp/previewawssupportappmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// You can use the `AWS::SupportApp::SlackWorkspaceConfiguration` resource to specify your Slack workspace configuration.
//
// This resource configures your AWS account so that you can use the specified Slack workspace in the AWS Support App . This resource includes the following information:
//
// - The team ID for the Slack workspace
// - The version ID of the resource to use with AWS CloudFormation
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
//   cfnSlackWorkspaceConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSlackWorkspaceConfigurationPropsMixin(&CfnSlackWorkspaceConfigurationMixinProps{
//   	TeamId: jsii.String("teamId"),
//   	VersionId: jsii.String("versionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html
//
type CfnSlackWorkspaceConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSlackWorkspaceConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSlackWorkspaceConfigurationPropsMixin
type jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin) Props() *CfnSlackWorkspaceConfigurationMixinProps {
	var returns *CfnSlackWorkspaceConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SupportApp::SlackWorkspaceConfiguration`.
func NewCfnSlackWorkspaceConfigurationPropsMixin(props *CfnSlackWorkspaceConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSlackWorkspaceConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSlackWorkspaceConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackWorkspaceConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SupportApp::SlackWorkspaceConfiguration`.
func NewCfnSlackWorkspaceConfigurationPropsMixin_Override(c CfnSlackWorkspaceConfigurationPropsMixin, props *CfnSlackWorkspaceConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackWorkspaceConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSlackWorkspaceConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSlackWorkspaceConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackWorkspaceConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSlackWorkspaceConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_supportapp.mixins.CfnSlackWorkspaceConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSlackWorkspaceConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

