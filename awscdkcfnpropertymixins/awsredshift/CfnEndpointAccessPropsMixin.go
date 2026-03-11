package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Redshift-managed VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEndpointAccessPropsMixin := awscdkcfnpropertymixins.Aws_redshift.NewCfnEndpointAccessPropsMixin(&CfnEndpointAccessMixinProps{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	EndpointName: jsii.String("endpointName"),
//   	ResourceOwner: jsii.String("resourceOwner"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	VpcSecurityGroupIds: []interface{}{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-endpointaccess.html
//
type CfnEndpointAccessPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEndpointAccessMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointAccessPropsMixin
type jsiiProxy_CfnEndpointAccessPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEndpointAccessPropsMixin) Props() *CfnEndpointAccessMixinProps {
	var returns *CfnEndpointAccessMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccessPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::EndpointAccess`.
func NewCfnEndpointAccessPropsMixin(props *CfnEndpointAccessMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEndpointAccessPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointAccessPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointAccessPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAccessPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::EndpointAccess`.
func NewCfnEndpointAccessPropsMixin_Override(c CfnEndpointAccessPropsMixin, props *CfnEndpointAccessMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAccessPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEndpointAccessPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointAccessPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAccessPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointAccessPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnEndpointAccessPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointAccessPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointAccessPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

