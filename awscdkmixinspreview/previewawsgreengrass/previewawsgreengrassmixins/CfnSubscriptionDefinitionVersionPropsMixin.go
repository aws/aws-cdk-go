package previewawsgreengrassmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrass/previewawsgreengrassmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::SubscriptionDefinitionVersion` resource represents a subscription definition version for AWS IoT Greengrass .
//
// A subscription definition version contains a list of subscriptions.
//
// > To create a subscription definition version, you must specify the ID of the subscription definition that you want to associate with the version. For information about creating a subscription definition, see [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) .
// >
// > After you create a subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriptionDefinitionVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnSubscriptionDefinitionVersionPropsMixin(&CfnSubscriptionDefinitionVersionMixinProps{
//   	SubscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   	Subscriptions: []interface{}{
//   		&SubscriptionProperty{
//   			Id: jsii.String("id"),
//   			Source: jsii.String("source"),
//   			Subject: jsii.String("subject"),
//   			Target: jsii.String("target"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html
//
type CfnSubscriptionDefinitionVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSubscriptionDefinitionVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriptionDefinitionVersionPropsMixin
type jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin) Props() *CfnSubscriptionDefinitionVersionMixinProps {
	var returns *CfnSubscriptionDefinitionVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::SubscriptionDefinitionVersion`.
func NewCfnSubscriptionDefinitionVersionPropsMixin(props *CfnSubscriptionDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSubscriptionDefinitionVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriptionDefinitionVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnSubscriptionDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::SubscriptionDefinitionVersion`.
func NewCfnSubscriptionDefinitionVersionPropsMixin_Override(c CfnSubscriptionDefinitionVersionPropsMixin, props *CfnSubscriptionDefinitionVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnSubscriptionDefinitionVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSubscriptionDefinitionVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriptionDefinitionVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnSubscriptionDefinitionVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionDefinitionVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnSubscriptionDefinitionVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSubscriptionDefinitionVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

