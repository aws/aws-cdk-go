package awstranslate

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstranslate/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A parallel data resource in Amazon Translate used to customize machine translation output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnParallelDataPropsMixin := awscdkcfnpropertymixins.Aws_translate.NewCfnParallelDataPropsMixin(&CfnParallelDataMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	ParallelDataConfig: &ParallelDataConfigProperty{
//   		Format: jsii.String("format"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-paralleldata.html
//
type CfnParallelDataPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnParallelDataMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnParallelDataPropsMixin
type jsiiProxy_CfnParallelDataPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnParallelDataPropsMixin) Props() *CfnParallelDataMixinProps {
	var returns *CfnParallelDataMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParallelDataPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Translate::ParallelData`.
func NewCfnParallelDataPropsMixin(props *CfnParallelDataMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnParallelDataPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnParallelDataPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnParallelDataPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnParallelDataPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Translate::ParallelData`.
func NewCfnParallelDataPropsMixin_Override(c CfnParallelDataPropsMixin, props *CfnParallelDataMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnParallelDataPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnParallelDataPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnParallelDataPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnParallelDataPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnParallelDataPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnParallelDataPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnParallelDataPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnParallelDataPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

