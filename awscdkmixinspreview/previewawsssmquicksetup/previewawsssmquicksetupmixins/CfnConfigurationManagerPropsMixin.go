package previewawsssmquicksetupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssmquicksetup/previewawsssmquicksetupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Quick Setup configuration manager resource.
//
// This object is a collection of desired state configurations for multiple configuration definitions and summaries describing the deployments of those definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationManagerPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationManagerPropsMixin(&CfnConfigurationManagerMixinProps{
//   	ConfigurationDefinitions: []interface{}{
//   		&ConfigurationDefinitionProperty{
//   			Id: jsii.String("id"),
//   			LocalDeploymentAdministrationRoleArn: jsii.String("localDeploymentAdministrationRoleArn"),
//   			LocalDeploymentExecutionRoleName: jsii.String("localDeploymentExecutionRoleName"),
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			Type: jsii.String("type"),
//   			TypeVersion: jsii.String("typeVersion"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html
//
type CfnConfigurationManagerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationManagerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationManagerPropsMixin
type jsiiProxy_CfnConfigurationManagerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationManagerPropsMixin) Props() *CfnConfigurationManagerMixinProps {
	var returns *CfnConfigurationManagerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationManagerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMQuickSetup::ConfigurationManager`.
func NewCfnConfigurationManagerPropsMixin(props *CfnConfigurationManagerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationManagerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationManagerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationManagerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmquicksetup.mixins.CfnConfigurationManagerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMQuickSetup::ConfigurationManager`.
func NewCfnConfigurationManagerPropsMixin_Override(c CfnConfigurationManagerPropsMixin, props *CfnConfigurationManagerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmquicksetup.mixins.CfnConfigurationManagerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationManagerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationManagerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssmquicksetup.mixins.CfnConfigurationManagerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationManagerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssmquicksetup.mixins.CfnConfigurationManagerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationManagerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfigurationManagerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

