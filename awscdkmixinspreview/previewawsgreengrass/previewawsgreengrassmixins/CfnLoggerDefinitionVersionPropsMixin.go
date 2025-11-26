package previewawsgreengrassmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrass/previewawsgreengrassmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::LoggerDefinitionVersion` resource represents a logger definition version for AWS IoT Greengrass .
//
// A logger definition version contains a list of loggers.
//
// > To create a logger definition version, you must specify the ID of the logger definition that you want to associate with the version. For information about creating a logger definition, see [`AWS::Greengrass::LoggerDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html) .
// >
// > After you create a logger definition version that contains the loggers you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoggerDefinitionVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnLoggerDefinitionVersionPropsMixin(&CfnLoggerDefinitionVersionMixinProps{
//   	LoggerDefinitionId: jsii.String("loggerDefinitionId"),
//   	Loggers: []interface{}{
//   		&LoggerProperty{
//   			Component: jsii.String("component"),
//   			Id: jsii.String("id"),
//   			Level: jsii.String("level"),
//   			Space: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html
//
type CfnLoggerDefinitionVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoggerDefinitionVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoggerDefinitionVersionPropsMixin
type jsiiProxy_CfnLoggerDefinitionVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) Props() *CfnLoggerDefinitionVersionMixinProps {
	var returns *CfnLoggerDefinitionVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::LoggerDefinitionVersion`.
func NewCfnLoggerDefinitionVersionPropsMixin(props *CfnLoggerDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoggerDefinitionVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoggerDefinitionVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoggerDefinitionVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnLoggerDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::LoggerDefinitionVersion`.
func NewCfnLoggerDefinitionVersionPropsMixin_Override(c CfnLoggerDefinitionVersionPropsMixin, props *CfnLoggerDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnLoggerDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoggerDefinitionVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoggerDefinitionVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnLoggerDefinitionVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggerDefinitionVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnLoggerDefinitionVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLoggerDefinitionVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

