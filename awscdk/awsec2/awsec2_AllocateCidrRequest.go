package awsec2


// Request for subnets Cidr to be allocated for a Vpc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allocateCidrRequest := &allocateCidrRequest{
//   	requestedSubnets: []requestedSubnet{
//   		&requestedSubnet{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			configuration: &subnetConfiguration{
//   				name: jsii.String("name"),
//   				subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//
//   				// the properties below are optional
//   				cidrMask: jsii.Number(123),
//   				mapPublicIpOnLaunch: jsii.Boolean(false),
//   				reserved: jsii.Boolean(false),
//   			},
//   			subnetConstructId: jsii.String("subnetConstructId"),
//   		},
//   	},
//   	vpcCidr: jsii.String("vpcCidr"),
//   }
//
type AllocateCidrRequest struct {
	// The Subnets to be allocated.
	RequestedSubnets *[]*RequestedSubnet `field:"required" json:"requestedSubnets" yaml:"requestedSubnets"`
	// The IPv4 CIDR block for this Vpc.
	VpcCidr *string `field:"required" json:"vpcCidr" yaml:"vpcCidr"`
}

