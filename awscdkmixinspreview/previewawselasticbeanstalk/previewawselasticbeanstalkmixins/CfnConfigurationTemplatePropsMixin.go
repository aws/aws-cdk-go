package previewawselasticbeanstalkmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticbeanstalk/previewawselasticbeanstalkmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify an AWS Elastic Beanstalk configuration template by using the AWS::ElasticBeanstalk::ConfigurationTemplate resource in an AWS CloudFormation template.
//
// The AWS::ElasticBeanstalk::ConfigurationTemplate resource is an AWS Elastic Beanstalk resource type that specifies an Elastic Beanstalk configuration template, associated with a specific Elastic Beanstalk application. You define application configuration settings in a configuration template. You can then use the configuration template to deploy different versions of the application with the same configuration settings.
//
// > The Elastic Beanstalk console and documentation often refer to configuration templates as *saved configurations* . When you set configuration options in a saved configuration (configuration template), Elastic Beanstalk applies them with a particular precedence as part of applying options from multiple sources. For more information, see [Configuration Options](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationTemplatePropsMixin(&CfnConfigurationTemplateMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Description: jsii.String("description"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	OptionSettings: []interface{}{
//   		&ConfigurationOptionSettingProperty{
//   			Namespace: jsii.String("namespace"),
//   			OptionName: jsii.String("optionName"),
//   			ResourceName: jsii.String("resourceName"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PlatformArn: jsii.String("platformArn"),
//   	SolutionStackName: jsii.String("solutionStackName"),
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		TemplateName: jsii.String("templateName"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html
//
type CfnConfigurationTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationTemplatePropsMixin
type jsiiProxy_CfnConfigurationTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationTemplatePropsMixin) Props() *CfnConfigurationTemplateMixinProps {
	var returns *CfnConfigurationTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticBeanstalk::ConfigurationTemplate`.
func NewCfnConfigurationTemplatePropsMixin(props *CfnConfigurationTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticBeanstalk::ConfigurationTemplate`.
func NewCfnConfigurationTemplatePropsMixin_Override(c CfnConfigurationTemplatePropsMixin, props *CfnConfigurationTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnConfigurationTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticbeanstalk.mixins.CfnConfigurationTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfigurationTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

