package previewawsappconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappconfig/previewawsappconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a new configuration in the AWS AppConfig hosted configuration store.
//
// Configurations must be 1 MB or smaller. The AWS AppConfig hosted configuration store provides the following benefits over other configuration store options.
//
// - You don't need to set up and configure other services such as Amazon Simple Storage Service ( Amazon S3 ) or Parameter Store.
// - You don't need to configure AWS Identity and Access Management ( IAM ) permissions to use the configuration store.
// - You can store configurations in any content type.
// - There is no cost to use the store.
// - You can create a configuration and add it to the store when you create a configuration profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHostedConfigurationVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnHostedConfigurationVersionPropsMixin(&CfnHostedConfigurationVersionMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	Content: jsii.String("content"),
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	LatestVersionNumber: jsii.Number(123),
//   	VersionLabel: jsii.String("versionLabel"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html
//
type CfnHostedConfigurationVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHostedConfigurationVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHostedConfigurationVersionPropsMixin
type jsiiProxy_CfnHostedConfigurationVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHostedConfigurationVersionPropsMixin) Props() *CfnHostedConfigurationVersionMixinProps {
	var returns *CfnHostedConfigurationVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostedConfigurationVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::HostedConfigurationVersion`.
func NewCfnHostedConfigurationVersionPropsMixin(props *CfnHostedConfigurationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHostedConfigurationVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHostedConfigurationVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHostedConfigurationVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnHostedConfigurationVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::HostedConfigurationVersion`.
func NewCfnHostedConfigurationVersionPropsMixin_Override(c CfnHostedConfigurationVersionPropsMixin, props *CfnHostedConfigurationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnHostedConfigurationVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHostedConfigurationVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHostedConfigurationVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnHostedConfigurationVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHostedConfigurationVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnHostedConfigurationVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHostedConfigurationVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnHostedConfigurationVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

