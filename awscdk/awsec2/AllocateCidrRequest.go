package awsec2


// Request for subnets CIDR to be allocated for a Vpc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allocateCidrRequest := &AllocateCidrRequest{
//   	RequestedSubnets: []requestedSubnet{
//   		&requestedSubnet{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			Configuration: &SubnetConfiguration{
//   				Name: jsii.String("name"),
//   				SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//
//   				// the properties below are optional
//   				CidrMask: jsii.Number(123),
//   				Ipv6AssignAddressOnCreation: jsii.Boolean(false),
//   				MapPublicIpOnLaunch: jsii.Boolean(false),
//   				Reserved: jsii.Boolean(false),
//   			},
//   			SubnetConstructId: jsii.String("subnetConstructId"),
//   		},
//   	},
//   	VpcCidr: jsii.String("vpcCidr"),
//   }
//
type AllocateCidrRequest struct {
	// The Subnets to be allocated.
	RequestedSubnets *[]*RequestedSubnet `field:"required" json:"requestedSubnets" yaml:"requestedSubnets"`
	// The IPv4 CIDR block for this Vpc.
	VpcCidr *string `field:"required" json:"vpcCidr" yaml:"vpcCidr"`
}

