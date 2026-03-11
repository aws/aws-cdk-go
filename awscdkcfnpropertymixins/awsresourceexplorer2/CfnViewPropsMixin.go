package awsresourceexplorer2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a view that users can query by using the [Search](https://docs.aws.amazon.com/resource-explorer/latest/apireference/API_Search.html) operation. Results from queries that you make using this view include only resources that match the view's `Filters` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnViewPropsMixin := awscdkcfnpropertymixins.Aws_resourceexplorer2.NewCfnViewPropsMixin(&CfnViewMixinProps{
//   	Filters: &FiltersProperty{
//   		FilterString: jsii.String("filterString"),
//   	},
//   	IncludedProperties: []interface{}{
//   		&IncludedPropertyProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	ViewName: jsii.String("viewName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html
//
type CfnViewPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnViewMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnViewPropsMixin
type jsiiProxy_CfnViewPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnViewPropsMixin) Props() *CfnViewMixinProps {
	var returns *CfnViewMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnViewPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResourceExplorer2::View`.
func NewCfnViewPropsMixin(props *CfnViewMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnViewPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnViewPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnViewPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnViewPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResourceExplorer2::View`.
func NewCfnViewPropsMixin_Override(c CfnViewPropsMixin, props *CfnViewMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnViewPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnViewPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnViewPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnViewPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnViewPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnViewPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnViewPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnViewPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

