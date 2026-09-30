package awscases

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscases/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Cases::RelatedItem.
//
// Creates a related item (comments) and associates it with a case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRelatedItemPropsMixin := awscdkcfnpropertymixins.Aws_cases.NewCfnRelatedItemPropsMixin(&CfnRelatedItemMixinProps{
//   	CaseId: jsii.String("caseId"),
//   	Content: &RelatedItemContentProperty{
//   		Comment: &CommentContentProperty{
//   			Body: jsii.String("body"),
//   			ContentType: jsii.String("contentType"),
//   		},
//   	},
//   	DomainId: jsii.String("domainId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html
//
type CfnRelatedItemPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRelatedItemMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRelatedItemPropsMixin
type jsiiProxy_CfnRelatedItemPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRelatedItemPropsMixin) Props() *CfnRelatedItemMixinProps {
	var returns *CfnRelatedItemMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRelatedItemPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cases::RelatedItem`.
func NewCfnRelatedItemPropsMixin(props *CfnRelatedItemMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRelatedItemPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRelatedItemPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRelatedItemPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnRelatedItemPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cases::RelatedItem`.
func NewCfnRelatedItemPropsMixin_Override(c CfnRelatedItemPropsMixin, props *CfnRelatedItemMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnRelatedItemPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRelatedItemPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRelatedItemPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnRelatedItemPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRelatedItemPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnRelatedItemPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRelatedItemPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRelatedItemPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

