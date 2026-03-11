package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an new set of frequently asked question (FAQ) questions and answers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFaqPropsMixin := awscdkcfnpropertymixins.Aws_kendra.NewCfnFaqPropsMixin(&CfnFaqMixinProps{
//   	Description: jsii.String("description"),
//   	FileFormat: jsii.String("fileFormat"),
//   	IndexId: jsii.String("indexId"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3Path: &S3PathProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-faq.html
//
type CfnFaqPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFaqMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFaqPropsMixin
type jsiiProxy_CfnFaqPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFaqPropsMixin) Props() *CfnFaqMixinProps {
	var returns *CfnFaqMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaqPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kendra::Faq`.
func NewCfnFaqPropsMixin(props *CfnFaqMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFaqPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFaqPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFaqPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFaqPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kendra::Faq`.
func NewCfnFaqPropsMixin_Override(c CfnFaqPropsMixin, props *CfnFaqMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFaqPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFaqPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFaqPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFaqPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFaqPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFaqPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFaqPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFaqPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

