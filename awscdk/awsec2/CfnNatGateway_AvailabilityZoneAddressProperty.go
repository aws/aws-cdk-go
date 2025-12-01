package awsec2


// For regional NAT gateways only: The configuration specifying which Elastic IP address (EIP) to use for handling outbound NAT traffic from a specific Availability Zone.
//
// A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.
//
// For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the *Amazon VPC User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZoneAddressProperty := &AvailabilityZoneAddressProperty{
//   	AllocationIds: []*string{
//   		jsii.String("allocationIds"),
//   	},
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html
//
type CfnNatGateway_AvailabilityZoneAddressProperty struct {
	// The allocation IDs of the Elastic IP addresses (EIPs) to be used for handling outbound NAT traffic in this specific Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-allocationids
	//
	AllocationIds *[]*string `field:"required" json:"allocationIds" yaml:"allocationIds"`
	// For regional NAT gateways only: The Availability Zone where this specific NAT gateway configuration will be active.
	//
	// Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ.
	//
	// A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// For regional NAT gateways only: The ID of the Availability Zone where this specific NAT gateway configuration will be active.
	//
	// Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ. Use this instead of AvailabilityZone for consistent identification of AZs across AWS Regions.
	//
	// A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-natgateway-availabilityzoneaddress.html#cfn-ec2-natgateway-availabilityzoneaddress-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
}

