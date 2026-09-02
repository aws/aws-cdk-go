package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A block list used for query suggestions for an Amazon Kendra index.
//
// A block list contains words or phrases that should not appear as query suggestions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnQuerySuggestionsBlockListPropsMixin := awscdkcfnpropertymixins.Aws_kendra.NewCfnQuerySuggestionsBlockListPropsMixin(&CfnQuerySuggestionsBlockListMixinProps{
//   	Description: jsii.String("description"),
//   	IndexId: jsii.String("indexId"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	SourceS3Path: &S3PathProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-querysuggestionsblocklist.html
//
type CfnQuerySuggestionsBlockListPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnQuerySuggestionsBlockListMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQuerySuggestionsBlockListPropsMixin
type jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin) Props() *CfnQuerySuggestionsBlockListMixinProps {
	var returns *CfnQuerySuggestionsBlockListMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kendra::QuerySuggestionsBlockList`.
func NewCfnQuerySuggestionsBlockListPropsMixin(props *CfnQuerySuggestionsBlockListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnQuerySuggestionsBlockListPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQuerySuggestionsBlockListPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnQuerySuggestionsBlockListPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kendra::QuerySuggestionsBlockList`.
func NewCfnQuerySuggestionsBlockListPropsMixin_Override(c CfnQuerySuggestionsBlockListPropsMixin, props *CfnQuerySuggestionsBlockListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnQuerySuggestionsBlockListPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnQuerySuggestionsBlockListPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQuerySuggestionsBlockListPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnQuerySuggestionsBlockListPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQuerySuggestionsBlockListPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnQuerySuggestionsBlockListPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnQuerySuggestionsBlockListPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

