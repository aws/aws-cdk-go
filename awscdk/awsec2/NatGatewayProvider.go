package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provider for NAT Gateways.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natGatewayProvider := awscdk.Aws_ec2.NewNatGatewayProvider(&NatGatewayProps{
//   	EipAllocationIds: []*string{
//   		jsii.String("eipAllocationIds"),
//   	},
//   })
//
type NatGatewayProvider interface {
	NatProvider
	// Return list of gateways spawned by the provider.
	ConfiguredGateways() *[]*GatewayConfig
	// Called by the VPC to configure NAT.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureNat(options *ConfigureNatOptions)
	// Configures subnet with the gateway.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureSubnet(subnet PrivateSubnet)
}

// The jsii proxy struct for NatGatewayProvider
type jsiiProxy_NatGatewayProvider struct {
	jsiiProxy_NatProvider
}

func (j *jsiiProxy_NatGatewayProvider) ConfiguredGateways() *[]*GatewayConfig {
	var returns *[]*GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}


func NewNatGatewayProvider(props *NatGatewayProps) NatGatewayProvider {
	_init_.Initialize()

	if err := validateNewNatGatewayProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NatGatewayProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatGatewayProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNatGatewayProvider_Override(n NatGatewayProvider, props *NatGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatGatewayProvider",
		[]interface{}{props},
		n,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func NatGatewayProvider_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	if err := validateNatGatewayProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns NatProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatGatewayProvider",
		"gateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use NAT instances to provide NAT services for your VPC.
//
// NAT instances are managed by you, but in return allow more configuration.
//
// Be aware that instances created using this provider will not be
// automatically replaced if they are stopped for any reason. You should implement
// your own NatProvider based on AutoScaling groups if you need that.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html
//
// Deprecated: use instanceV2. 'instance' is deprecated since NatInstanceProvider
// uses a instance image that has reached EOL on Dec 31 2023.
func NatGatewayProvider_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateNatGatewayProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatGatewayProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use NAT instances to provide NAT services for your VPC.
//
// NAT instances are managed by you, but in return allow more configuration.
//
// Be aware that instances created using this provider will not be
// automatically replaced if they are stopped for any reason. You should implement
// your own NatProvider based on AutoScaling groups if you need that.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html
//
func NatGatewayProvider_InstanceV2(props *NatInstanceProps) NatInstanceProviderV2 {
	_init_.Initialize()

	if err := validateNatGatewayProvider_InstanceV2Parameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProviderV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatGatewayProvider",
		"instanceV2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NatGatewayProvider) ConfigureNat(options *ConfigureNatOptions) {
	if err := n.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureNat",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_NatGatewayProvider) ConfigureSubnet(subnet PrivateSubnet) {
	if err := n.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

