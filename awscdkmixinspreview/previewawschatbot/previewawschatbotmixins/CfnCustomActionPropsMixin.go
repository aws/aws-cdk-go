package previewawschatbotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawschatbot/previewawschatbotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > AWS Chatbot is now  .
//
// [Learn more](https://docs.aws.amazon.com//chatbot/latest/adminguide/service-rename.html)
// >
// > `Type` attribute values remain unchanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomActionPropsMixin := awscdkmixinspreview.Mixins.NewCfnCustomActionPropsMixin(&CfnCustomActionMixinProps{
//   	ActionName: jsii.String("actionName"),
//   	AliasName: jsii.String("aliasName"),
//   	Attachments: []interface{}{
//   		&CustomActionAttachmentProperty{
//   			ButtonText: jsii.String("buttonText"),
//   			Criteria: []interface{}{
//   				&CustomActionAttachmentCriteriaProperty{
//   					Operator: jsii.String("operator"),
//   					Value: jsii.String("value"),
//   					VariableName: jsii.String("variableName"),
//   				},
//   			},
//   			NotificationType: jsii.String("notificationType"),
//   			Variables: map[string]*string{
//   				"variablesKey": jsii.String("variables"),
//   			},
//   		},
//   	},
//   	Definition: &CustomActionDefinitionProperty{
//   		CommandText: jsii.String("commandText"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html
//
type CfnCustomActionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCustomActionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomActionPropsMixin
type jsiiProxy_CfnCustomActionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCustomActionPropsMixin) Props() *CfnCustomActionMixinProps {
	var returns *CfnCustomActionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Chatbot::CustomAction`.
func NewCfnCustomActionPropsMixin(props *CfnCustomActionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCustomActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnCustomActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Chatbot::CustomAction`.
func NewCfnCustomActionPropsMixin_Override(c CfnCustomActionPropsMixin, props *CfnCustomActionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnCustomActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCustomActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnCustomActionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomActionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_chatbot.mixins.CfnCustomActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomActionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomActionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

