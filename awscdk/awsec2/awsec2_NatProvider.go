package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NAT providers.
//
// Determines what type of NAT provider to create, either NAT gateways or NAT
// instance.
//
// Example:
//   // Configure the `natGatewayProvider` when defining a Vpc
//   natGatewayProvider := ec2.natProvider.instance(&natInstanceProps{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   })
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
//   	natGatewayProvider: natGatewayProvider,
//
//   	// The 'natGateways' parameter now controls the number of NAT instances
//   	natGateways: jsii.Number(2),
//   })
//
// Experimental.
type NatProvider interface {
	// Return list of gateways spawned by the provider.
	// Experimental.
	ConfiguredGateways() *[]*GatewayConfig
	// Called by the VPC to configure NAT.
	//
	// Don't call this directly, the VPC will call it automatically.
	// Experimental.
	ConfigureNat(options *ConfigureNatOptions)
	// Configures subnet with the gateway.
	//
	// Don't call this directly, the VPC will call it automatically.
	// Experimental.
	ConfigureSubnet(subnet PrivateSubnet)
}

// The jsii proxy struct for NatProvider
type jsiiProxy_NatProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_NatProvider) ConfiguredGateways() *[]*GatewayConfig {
	var returns *[]*GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}


// Experimental.
func NewNatProvider_Override(n NatProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.NatProvider",
		nil, // no parameters
		n,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
// Experimental.
func NatProvider_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	var returns NatProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.NatProvider",
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
// Experimental.
func NatProvider_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.NatProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NatProvider) ConfigureNat(options *ConfigureNatOptions) {
	_jsii_.InvokeVoid(
		n,
		"configureNat",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_NatProvider) ConfigureSubnet(subnet PrivateSubnet) {
	_jsii_.InvokeVoid(
		n,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

