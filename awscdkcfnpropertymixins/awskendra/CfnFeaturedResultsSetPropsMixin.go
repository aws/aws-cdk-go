package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Kendra::FeaturedResultsSet.
//
// A set of featured results that are displayed at the top of your search results. Featured results are placed above all other results for certain queries. If there's an exact match of a query, then one or more specific documents are featured in the search results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFeaturedResultsSetPropsMixin := awscdkcfnpropertymixins.Aws_kendra.NewCfnFeaturedResultsSetPropsMixin(&CfnFeaturedResultsSetMixinProps{
//   	Description: jsii.String("description"),
//   	FeaturedDocuments: []interface{}{
//   		&FeaturedDocumentProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	FeaturedResultsSetName: jsii.String("featuredResultsSetName"),
//   	IndexId: jsii.String("indexId"),
//   	QueryTexts: []*string{
//   		jsii.String("queryTexts"),
//   	},
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html
//
type CfnFeaturedResultsSetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFeaturedResultsSetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFeaturedResultsSetPropsMixin
type jsiiProxy_CfnFeaturedResultsSetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFeaturedResultsSetPropsMixin) Props() *CfnFeaturedResultsSetMixinProps {
	var returns *CfnFeaturedResultsSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeaturedResultsSetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kendra::FeaturedResultsSet`.
func NewCfnFeaturedResultsSetPropsMixin(props *CfnFeaturedResultsSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFeaturedResultsSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFeaturedResultsSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFeaturedResultsSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFeaturedResultsSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kendra::FeaturedResultsSet`.
func NewCfnFeaturedResultsSetPropsMixin_Override(c CfnFeaturedResultsSetPropsMixin, props *CfnFeaturedResultsSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFeaturedResultsSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFeaturedResultsSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFeaturedResultsSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFeaturedResultsSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeaturedResultsSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_kendra.CfnFeaturedResultsSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeaturedResultsSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFeaturedResultsSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

