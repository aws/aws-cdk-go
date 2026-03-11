package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a CIDR collection in the current AWS account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCidrCollectionPropsMixin := awscdkcfnpropertymixins.Aws_route53.NewCfnCidrCollectionPropsMixin(&CfnCidrCollectionMixinProps{
//   	Locations: []interface{}{
//   		&LocationProperty{
//   			CidrList: []*string{
//   				jsii.String("cidrList"),
//   			},
//   			LocationName: jsii.String("locationName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html
//
type CfnCidrCollectionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCidrCollectionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCidrCollectionPropsMixin
type jsiiProxy_CfnCidrCollectionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCidrCollectionPropsMixin) Props() *CfnCidrCollectionMixinProps {
	var returns *CfnCidrCollectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCidrCollectionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::CidrCollection`.
func NewCfnCidrCollectionPropsMixin(props *CfnCidrCollectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCidrCollectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCidrCollectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCidrCollectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::CidrCollection`.
func NewCfnCidrCollectionPropsMixin_Override(c CfnCidrCollectionPropsMixin, props *CfnCidrCollectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCidrCollectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCidrCollectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCidrCollectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCidrCollectionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCidrCollectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

