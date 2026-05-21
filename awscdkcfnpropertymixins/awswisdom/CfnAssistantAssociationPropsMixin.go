package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an association between an Connect Customer Wisdom assistant and another resource.
//
// Currently, the only supported association is with a knowledge base. An assistant can have only a single association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAssistantAssociationPropsMixin := awscdkcfnpropertymixins.Aws_wisdom.NewCfnAssistantAssociationPropsMixin(&CfnAssistantAssociationMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html
//
type CfnAssistantAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAssistantAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssistantAssociationPropsMixin
type jsiiProxy_CfnAssistantAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnAssistantAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociationPropsMixin(props *CfnAssistantAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAssistantAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssistantAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssistantAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociationPropsMixin_Override(c CfnAssistantAssociationPropsMixin, props *CfnAssistantAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAssistantAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssistantAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssistantAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

