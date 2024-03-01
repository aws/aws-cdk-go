package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NAT providers.
//
// Determines what type of NAT provider to create, either NAT gateways or NAT
// instance.
//
// Example:
//   natInstanceProvider := ec2.NatProvider_Instance(&NatInstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T4G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.NewAmazonLinuxImage(),
//   	CreditSpecification: ec2.CpuCredits_UNLIMITED,
//   })
//   ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
//   	NatGatewayProvider: natInstanceProvider,
//   })
//
type NatProvider interface {
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


func NewNatProvider_Override(n NatProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatProvider",
		nil, // no parameters
		n,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func NatProvider_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	if err := validateNatProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns NatProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatProvider",
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
func NatProvider_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateNatProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.NatProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NatProvider) ConfigureNat(options *ConfigureNatOptions) {
	if err := n.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureNat",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_NatProvider) ConfigureSubnet(subnet PrivateSubnet) {
	if err := n.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

