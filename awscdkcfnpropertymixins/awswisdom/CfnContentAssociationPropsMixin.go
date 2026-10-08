package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::Wisdom::ContentAssociation Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContentAssociationPropsMixin := awscdkcfnpropertymixins.Aws_wisdom.NewCfnContentAssociationPropsMixin(&CfnContentAssociationMixinProps{
//   	Association: &ContentAssociationContentsProperty{
//   		AmazonConnectGuideAssociation: &AmazonConnectGuideAssociationDataProperty{
//   			FlowId: jsii.String("flowId"),
//   		},
//   	},
//   	AssociationType: jsii.String("associationType"),
//   	ContentId: jsii.String("contentId"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html
//
type CfnContentAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContentAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContentAssociationPropsMixin
type jsiiProxy_CfnContentAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContentAssociationPropsMixin) Props() *CfnContentAssociationMixinProps {
	var returns *CfnContentAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContentAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::ContentAssociation`.
func NewCfnContentAssociationPropsMixin(props *CfnContentAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContentAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContentAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContentAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::ContentAssociation`.
func NewCfnContentAssociationPropsMixin_Override(c CfnContentAssociationPropsMixin, props *CfnContentAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContentAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContentAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContentAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContentAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContentAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

