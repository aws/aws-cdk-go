package awsresourceexplorer2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Sets the specified view as the default for the AWS Region in which you call this operation.
//
// If a user makes a search query that doesn't explicitly specify the view to use, Resource Explorer chooses this default view automatically for searches performed in this AWS Region .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDefaultViewAssociationPropsMixin := awscdkcfnpropertymixins.Aws_resourceexplorer2.NewCfnDefaultViewAssociationPropsMixin(&CfnDefaultViewAssociationMixinProps{
//   	ViewArn: jsii.String("viewArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-defaultviewassociation.html
//
type CfnDefaultViewAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDefaultViewAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDefaultViewAssociationPropsMixin
type jsiiProxy_CfnDefaultViewAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDefaultViewAssociationPropsMixin) Props() *CfnDefaultViewAssociationMixinProps {
	var returns *CfnDefaultViewAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDefaultViewAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResourceExplorer2::DefaultViewAssociation`.
func NewCfnDefaultViewAssociationPropsMixin(props *CfnDefaultViewAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDefaultViewAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDefaultViewAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDefaultViewAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnDefaultViewAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResourceExplorer2::DefaultViewAssociation`.
func NewCfnDefaultViewAssociationPropsMixin_Override(c CfnDefaultViewAssociationPropsMixin, props *CfnDefaultViewAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnDefaultViewAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDefaultViewAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDefaultViewAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnDefaultViewAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDefaultViewAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_resourceexplorer2.CfnDefaultViewAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDefaultViewAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDefaultViewAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

