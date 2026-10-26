package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An association between an IPAM and a Regional Internet Registry (RIR) for delegated Resource Public Key Infrastructure (RPKI) management.
//
// Use it to have IPAM create Route Origin Authorizations (ROAs) for IP address prefixes registered with the internet registry. The IPAM must be in the Advanced tier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIpamInternetRegistryAssociationPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnIpamInternetRegistryAssociationPropsMixin(&CfnIpamInternetRegistryAssociationMixinProps{
//   	Description: jsii.String("description"),
//   	IpamId: jsii.String("ipamId"),
//   	OrganizationHandle: jsii.String("organizationHandle"),
//   	Rir: jsii.String("rir"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html
//
type CfnIpamInternetRegistryAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIpamInternetRegistryAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIpamInternetRegistryAssociationPropsMixin
type jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin) Props() *CfnIpamInternetRegistryAssociationMixinProps {
	var returns *CfnIpamInternetRegistryAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::IpamInternetRegistryAssociation`.
func NewCfnIpamInternetRegistryAssociationPropsMixin(props *CfnIpamInternetRegistryAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIpamInternetRegistryAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIpamInternetRegistryAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamInternetRegistryAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::IpamInternetRegistryAssociation`.
func NewCfnIpamInternetRegistryAssociationPropsMixin_Override(c CfnIpamInternetRegistryAssociationPropsMixin, props *CfnIpamInternetRegistryAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamInternetRegistryAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIpamInternetRegistryAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIpamInternetRegistryAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamInternetRegistryAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIpamInternetRegistryAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamInternetRegistryAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIpamInternetRegistryAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

