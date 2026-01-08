package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswisdom/previewawswisdommixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an association between an Amazon Connect Wisdom assistant and another resource.
//
// Currently, the only supported association is with a knowledge base. An assistant can have only a single association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAssistantAssociationPropsMixin(&CfnAssistantAssociationMixinProps{
//   	AssistantId: jsii.String("assistantId"),
//   	Association: &AssociationDataProperty{
//   		ExternalBedrockKnowledgeBaseConfig: &ExternalBedrockKnowledgeBaseConfigProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   		},
//   		KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	},
//   	AssociationType: jsii.String("associationType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html
//
type CfnAssistantAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAssistantAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssistantAssociationPropsMixin
type jsiiProxy_CfnAssistantAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAssistantAssociationPropsMixin) Props() *CfnAssistantAssociationMixinProps {
	var returns *CfnAssistantAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociationPropsMixin(props *CfnAssistantAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAssistantAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssistantAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssistantAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociationPropsMixin_Override(c CfnAssistantAssociationPropsMixin, props *CfnAssistantAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAssistantAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssistantAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssistantAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssistantAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAssistantAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

