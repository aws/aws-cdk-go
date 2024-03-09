package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Modern NAT provider which uses NAT Instances.
//
// The instance uses Amazon Linux 2023 as the operating system.
//
// Example:
//   var instanceType instanceType
//
//
//   provider := ec2.NatProvider_InstanceV2(&NatInstanceProps{
//   	InstanceType: InstanceType,
//   	DefaultAllowedTraffic: ec2.NatTrafficDirection_OUTBOUND_ONLY,
//   })
//   ec2.NewVpc(this, jsii.String("TheVPC"), &VpcProps{
//   	NatGatewayProvider: provider,
//   })
//   provider.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(80)))
//
type NatInstanceProviderV2 interface {
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

// The jsii proxy struct for NatInstanceProviderV2
type jsiiProxy_NatInstanceProviderV2 struct {
	jsiiProxy_NatProvider
	jsiiProxy_IConnectable
}

func (j *jsiiProxy_NatInstanceProviderV2) ConfiguredGateways() *[]*GatewayConfig {
	var returns *[]*GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NatInstanceProviderV2) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NatInstanceProviderV2) SecurityGroup() ISecurityGroup {
	var returns ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}


func NewNatInstanceProviderV2(props *NatInstanceProps) NatInstanceProviderV2 {
	_init_.Initialize()

	if err := validateNewNatInstanceProviderV2Parameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NatInstanceProviderV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNatInstanceProviderV2_Override(n NatInstanceProviderV2, props *NatInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
		[]interface{}{props},
		n,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func NatInstanceProviderV2_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	if err := validateNatInstanceProviderV2_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns NatProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
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
func NatInstanceProviderV2_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateNatInstanceProviderV2_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
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
func NatInstanceProviderV2_InstanceV2(props *NatInstanceProps) NatInstanceProviderV2 {
	_init_.Initialize()

	if err := validateNatInstanceProviderV2_InstanceV2Parameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProviderV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
		"instanceV2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NatInstanceProviderV2) ConfigureNat(options *ConfigureNatOptions) {
	if err := n.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureNat",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_NatInstanceProviderV2) ConfigureSubnet(subnet PrivateSubnet) {
	if err := n.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

