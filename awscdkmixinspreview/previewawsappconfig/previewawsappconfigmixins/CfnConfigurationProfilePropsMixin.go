package previewawsappconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappconfig/previewawsappconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppConfig::ConfigurationProfile` resource creates a configuration profile that enables AWS AppConfig to access the configuration source.
//
// Valid configuration sources include AWS Systems Manager (SSM) documents, SSM Parameter Store parameters, and Amazon S3 . A configuration profile includes the following information.
//
// - The Uri location of the configuration data.
// - The AWS Identity and Access Management ( IAM ) role that provides access to the configuration data.
// - A validator for the configuration data. Available validators include either a JSON Schema or the Amazon Resource Name (ARN) of an AWS Lambda function.
//
// AWS AppConfig requires that you create resources and deploy a configuration in the following order:
//
// - Create an application
// - Create an environment
// - Create a configuration profile
// - Choose a pre-defined deployment strategy or create your own
// - Deploy the configuration
//
// For more information, see [AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationProfilePropsMixin(&CfnConfigurationProfileMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	DeletionProtectionCheck: jsii.String("deletionProtectionCheck"),
//   	Description: jsii.String("description"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	LocationUri: jsii.String("locationUri"),
//   	Name: jsii.String("name"),
//   	RetrievalRoleArn: jsii.String("retrievalRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Validators: []interface{}{
//   		&ValidatorsProperty{
//   			Content: jsii.String("content"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html
//
type CfnConfigurationProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationProfilePropsMixin
type jsiiProxy_CfnConfigurationProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationProfilePropsMixin) Props() *CfnConfigurationProfileMixinProps {
	var returns *CfnConfigurationProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::ConfigurationProfile`.
func NewCfnConfigurationProfilePropsMixin(props *CfnConfigurationProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnConfigurationProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::ConfigurationProfile`.
func NewCfnConfigurationProfilePropsMixin_Override(c CfnConfigurationProfilePropsMixin, props *CfnConfigurationProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnConfigurationProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnConfigurationProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnConfigurationProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfigurationProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

