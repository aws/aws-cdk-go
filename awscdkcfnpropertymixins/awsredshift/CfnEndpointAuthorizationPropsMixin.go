package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEndpointAuthorizationPropsMixin := awscdkcfnpropertymixins.Aws_redshift.NewCfnEndpointAuthorizationPropsMixin(&CfnEndpointAuthorizationMixinProps{
//   	Account: jsii.String("account"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	Force: jsii.Boolean(false),
//   	VpcIds: []interface{}{
//   		jsii.String("vpcIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointauthorization.html
//
type CfnEndpointAuthorizationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEndpointAuthorizationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointAuthorizationPropsMixin
type jsiiProxy_CfnEndpointAuthorizationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEndpointAuthorizationPropsMixin) Props() *CfnEndpointAuthorizationMixinProps {
	var returns *CfnEndpointAuthorizationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorizationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::EndpointAuthorization`.
func NewCfnEndpointAuthorizationPropsMixin(props *CfnEndpointAuthorizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEndpointAuthorizationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointAuthorizationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointAuthorizationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAuthorizationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::EndpointAuthorization`.
func NewCfnEndpointAuthorizationPropsMixin_Override(c CfnEndpointAuthorizationPropsMixin, props *CfnEndpointAuthorizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAuthorizationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEndpointAuthorizationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointAuthorizationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAuthorizationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointAuthorizationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAuthorizationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorizationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorizationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

