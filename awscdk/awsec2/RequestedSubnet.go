package awsec2


// Subnet requested for allocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestedSubnet := &RequestedSubnet{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Configuration: &SubnetConfiguration{
//   		Name: jsii.String("name"),
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//
//   		// the properties below are optional
//   		CidrMask: jsii.Number(123),
//   		Ipv6AssignAddressOnCreation: jsii.Boolean(false),
//   		MapPublicIpOnLaunch: jsii.Boolean(false),
//   		Reserved: jsii.Boolean(false),
//   	},
//   	SubnetConstructId: jsii.String("subnetConstructId"),
//   }
//
type RequestedSubnet struct {
	// The availability zone for the subnet.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Specify configuration parameters for a single subnet group in a VPC.
	Configuration *SubnetConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Id for the Subnet construct.
	SubnetConstructId *string `field:"required" json:"subnetConstructId" yaml:"subnetConstructId"`
}

