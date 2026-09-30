package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Properties for a Regional NAT Gateway Provider.
//
// Regional NAT Gateways provide automatic multi-AZ redundancy with a single
// gateway that scales across availability zones.
//
// Example:
//   eip := ec2.NewCfnEIP(this, jsii.String("NatEip"))
//
//   ec2.NewVpc(this, jsii.String("Vpc1"), &VpcProps{
//   	NatGatewayProvider: ec2.NatProvider_RegionalGateway(&RegionalNatGatewayProviderProps{
//   		Eip: *Eip,
//   	}),
//   })
//
//   ec2.NewVpc(this, jsii.String("Vpc2"), &VpcProps{
//   	NatGatewayProvider: ec2.NatProvider_*RegionalGateway(&RegionalNatGatewayProviderProps{
//   		AllocationId: eip.attrAllocationId,
//   	}),
//   })
//
type RegionalNatGatewayProviderProps struct {
	// The allocation ID of the Elastic IP address to use for this NAT gateway.
	//
	// Cannot be specified together with `eip`.
	// Ignored when `availabilityZoneAddresses` is specified.
	// Default: - A new EIP is automatically allocated.
	//
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Specifies which Availability Zones you want the NAT gateway to support and the Elastic IP addresses to use in each AZ.
	//
	// When specified, `allocationId` and `eip` are ignored.
	// This enables manual mode for Regional NAT Gateway where you control EIP allocation per AZ.
	//
	// In manual mode the gateway only serves the listed Availability Zones and does not
	// expand to other zones automatically, so every Availability Zone that contains a
	// private subnet routed through this gateway must be listed. When the zone names are
	// known at synthesis time, a private subnet in an unlisted zone is rejected with an
	// error. Entries that use `availabilityZoneId` or unresolved tokens cannot be checked
	// and are assumed to be correct.
	// Default: - Automatic mode: AWS manages AZ coverage and EIP allocation.
	//
	AvailabilityZoneAddresses *[]*AvailabilityZoneAddress `field:"optional" json:"availabilityZoneAddresses" yaml:"availabilityZoneAddresses"`
	// Reference to an existing EIP to use for this NAT gateway.
	//
	// Cannot be specified together with `allocationId`.
	// Ignored when `availabilityZoneAddresses` is specified.
	// Default: - A new EIP is automatically allocated.
	//
	Eip interfacesawsec2.IEIPRef `field:"optional" json:"eip" yaml:"eip"`
	// Maximum amount of time to wait before forcibly releasing IP addresses if connections are still in progress.
	// Default: Duration.seconds(350)
	//
	MaxDrainDuration awscdk.Duration `field:"optional" json:"maxDrainDuration" yaml:"maxDrainDuration"`
}

