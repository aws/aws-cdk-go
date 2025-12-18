package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a network interface in an Amazon EC2 instance for AWS CloudFormation .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInterfacePropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkInterfacePropsMixin(&CfnNetworkInterfaceMixinProps{
//   	ConnectionTrackingSpecification: &ConnectionTrackingSpecificationProperty{
//   		TcpEstablishedTimeout: jsii.Number(123),
//   		UdpStreamTimeout: jsii.Number(123),
//   		UdpTimeout: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	EnablePrimaryIpv6: jsii.Boolean(false),
//   	GroupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	InterfaceType: jsii.String("interfaceType"),
//   	Ipv4PrefixCount: jsii.Number(123),
//   	Ipv4Prefixes: []interface{}{
//   		&Ipv4PrefixSpecificationProperty{
//   			Ipv4Prefix: jsii.String("ipv4Prefix"),
//   		},
//   	},
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&InstanceIpv6AddressProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	Ipv6PrefixCount: jsii.Number(123),
//   	Ipv6Prefixes: []interface{}{
//   		&Ipv6PrefixSpecificationProperty{
//   			Ipv6Prefix: jsii.String("ipv6Prefix"),
//   		},
//   	},
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PrivateIpAddresses: []interface{}{
//   		&PrivateIpAddressSpecificationProperty{
//   			Primary: jsii.Boolean(false),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	PublicIpDnsHostnameTypeSpecification: jsii.String("publicIpDnsHostnameTypeSpecification"),
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SourceDestCheck: jsii.Boolean(false),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html
//
type CfnNetworkInterfacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkInterfaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkInterfacePropsMixin
type jsiiProxy_CfnNetworkInterfacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNetworkInterfacePropsMixin) Props() *CfnNetworkInterfaceMixinProps {
	var returns *CfnNetworkInterfaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkInterfacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkInterface`.
func NewCfnNetworkInterfacePropsMixin(props *CfnNetworkInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkInterfacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkInterfacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkInterfacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkInterface`.
func NewCfnNetworkInterfacePropsMixin_Override(c CfnNetworkInterfacePropsMixin, props *CfnNetworkInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkInterfacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkInterfacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkInterfacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkInterfacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNetworkInterfacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

