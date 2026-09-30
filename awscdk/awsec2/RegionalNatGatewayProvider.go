package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provider for Regional NAT Gateways.
//
// Regional NAT Gateways provide automatic multi-AZ redundancy with a single gateway that scales across availability zones.
// Unlike zonal NAT gateways, a regional NAT gateway does not require a public subnet and is created at the VPC level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eIPRef IEIPRef
//
//   regionalNatGatewayProvider := awscdk.Aws_ec2.NewRegionalNatGatewayProvider(&RegionalNatGatewayProviderProps{
//   	AllocationId: jsii.String("allocationId"),
//   	AvailabilityZoneAddresses: []AvailabilityZoneAddress{
//   		&AvailabilityZoneAddress{
//   			AllocationIds: []*string{
//   				jsii.String("allocationIds"),
//   			},
//
//   			// the properties below are optional
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   		},
//   	},
//   	Eip: eIPRef,
//   	MaxDrainDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   })
//
// See: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html
//
type RegionalNatGatewayProvider interface {
	NatProvider
	// Return list of gateways spawned by the provider.
	//
	// A regional NAT gateway is not bound to a single Availability Zone, so `az`
	// is the sentinel string `'regional'` rather than a zone name.
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

// The jsii proxy struct for RegionalNatGatewayProvider
type jsiiProxy_RegionalNatGatewayProvider struct {
	jsiiProxy_NatProvider
}

func (j *jsiiProxy_RegionalNatGatewayProvider) ConfiguredGateways() *[]*GatewayConfig {
	var returns *[]*GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}


func NewRegionalNatGatewayProvider(props *RegionalNatGatewayProviderProps) RegionalNatGatewayProvider {
	_init_.Initialize()

	if err := validateNewRegionalNatGatewayProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RegionalNatGatewayProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewRegionalNatGatewayProvider_Override(r RegionalNatGatewayProvider, props *RegionalNatGatewayProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
		[]interface{}{props},
		r,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func RegionalNatGatewayProvider_Gateway(props *NatGatewayProps) NatProvider {
	_init_.Initialize()

	if err := validateRegionalNatGatewayProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns NatProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
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
func RegionalNatGatewayProvider_Instance(props *NatInstanceProps) NatInstanceProvider {
	_init_.Initialize()

	if err := validateRegionalNatGatewayProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
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
func RegionalNatGatewayProvider_InstanceV2(props *NatInstanceProps) NatInstanceProviderV2 {
	_init_.Initialize()

	if err := validateRegionalNatGatewayProvider_InstanceV2Parameters(props); err != nil {
		panic(err)
	}
	var returns NatInstanceProviderV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
		"instanceV2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use a Regional NAT Gateway to provide NAT services for your VPC.
//
// Regional NAT Gateways provide automatic multi-AZ redundancy with a single
// gateway that scales across availability zones. AWS automatically manages
// AZ coverage and EIP allocation.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html
//
func RegionalNatGatewayProvider_RegionalGateway(props *RegionalNatGatewayProviderProps) RegionalNatGatewayProvider {
	_init_.Initialize()

	if err := validateRegionalNatGatewayProvider_RegionalGatewayParameters(props); err != nil {
		panic(err)
	}
	var returns RegionalNatGatewayProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.RegionalNatGatewayProvider",
		"regionalGateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegionalNatGatewayProvider) ConfigureNat(options *ConfigureNatOptions) {
	if err := r.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureNat",
		[]interface{}{options},
	)
}

func (r *jsiiProxy_RegionalNatGatewayProvider) ConfigureSubnet(subnet PrivateSubnet) {
	if err := r.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

