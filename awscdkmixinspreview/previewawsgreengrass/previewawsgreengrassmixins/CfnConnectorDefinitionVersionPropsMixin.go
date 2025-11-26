package previewawsgreengrassmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrass/previewawsgreengrassmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::ConnectorDefinitionVersion` resource represents a connector definition version for AWS IoT Greengrass .
//
// A connector definition version contains a list of connectors.
//
// > To create a connector definition version, you must specify the ID of the connector definition that you want to associate with the version. For information about creating a connector definition, see [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html) .
// >
// > After you create a connector definition version that contains the connectors you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnConnectorDefinitionVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectorDefinitionVersionPropsMixin(&CfnConnectorDefinitionVersionMixinProps{
//   	ConnectorDefinitionId: jsii.String("connectorDefinitionId"),
//   	Connectors: []interface{}{
//   		&ConnectorProperty{
//   			ConnectorArn: jsii.String("connectorArn"),
//   			Id: jsii.String("id"),
//   			Parameters: parameters,
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html
//
type CfnConnectorDefinitionVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectorDefinitionVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectorDefinitionVersionPropsMixin
type jsiiProxy_CfnConnectorDefinitionVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectorDefinitionVersionPropsMixin) Props() *CfnConnectorDefinitionVersionMixinProps {
	var returns *CfnConnectorDefinitionVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorDefinitionVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::ConnectorDefinitionVersion`.
func NewCfnConnectorDefinitionVersionPropsMixin(props *CfnConnectorDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectorDefinitionVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectorDefinitionVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorDefinitionVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnConnectorDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::ConnectorDefinitionVersion`.
func NewCfnConnectorDefinitionVersionPropsMixin_Override(c CfnConnectorDefinitionVersionPropsMixin, props *CfnConnectorDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnConnectorDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectorDefinitionVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorDefinitionVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnConnectorDefinitionVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorDefinitionVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnConnectorDefinitionVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorDefinitionVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConnectorDefinitionVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

