package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::Wisdom::Content Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContentPropsMixin := awscdkcfnpropertymixins.Aws_wisdom.NewCfnContentPropsMixin(&CfnContentMixinProps{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	Name: jsii.String("name"),
//   	OverrideLinkOutUri: jsii.String("overrideLinkOutUri"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	UploadId: jsii.String("uploadId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html
//
type CfnContentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContentPropsMixin
type jsiiProxy_CfnContentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContentPropsMixin) Props() *CfnContentMixinProps {
	var returns *CfnContentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::Content`.
func NewCfnContentPropsMixin(props *CfnContentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::Content`.
func NewCfnContentPropsMixin_Override(c CfnContentPropsMixin, props *CfnContentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnContentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

