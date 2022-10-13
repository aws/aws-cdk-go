package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NAT provider which uses NAT Instances.
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
type NatInstanceProvider interface {
	NatProvider
	IConnectable
	// Return list of gateways spawned by the provider.
	ConfiguredGateways() *[]*GatewayConfig
	// Manage the Security Groups associated with the NAT instances.
	Connections() Connections
	// The Security Group associated with the NAT instances.
	SecurityGroup() ISecurityGroup
	// Called by the VPC to configure NAT.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureNat(options *ConfigureNatOptions)
	// Configures subnet with the gateway.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureSubnet(subnet PrivateSubnet)
}

// The jsii proxy struct for NatInstanceProvider
type jsiiProxy_NatInstanceProvider struct {
	jsiiProxy_NatProvider
	jsiiProxy_IConnectable
}

func (j *jsiiProxy_NatInstanceProvider) ConfiguredGateways() *[]*GatewayConfig {
	var returns *[]*GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NatInstanceProvider) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NatInstanceProvider) SecurityGroup() ISecurityGroup {
	var returns ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}


func NewNatInstanceProvider(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateNewNatInstanceProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NatInstanceProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNatInstanceProvider_Override(n NatInstanceProvider, props *NatInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceProvider",
		[]interface{}{props},
		n,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func NatInstanceProvider_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	if err := validateNatInstanceProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns NatProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatInstanceProvider",
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
func NatInstanceProvider_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateNatInstanceProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatInstanceProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NatInstanceProvider) ConfigureNat(options *ConfigureNatOptions) {
	if err := n.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureNat",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_NatInstanceProvider) ConfigureSubnet(subnet PrivateSubnet) {
	if err := n.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

