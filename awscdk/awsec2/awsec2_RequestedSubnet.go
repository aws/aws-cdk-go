package awsec2


// Subnet requested for allocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestedSubnet := &requestedSubnet{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	configuration: &subnetConfiguration{
//   		name: jsii.String("name"),
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//
//   		// the properties below are optional
//   		cidrMask: jsii.Number(123),
//   		mapPublicIpOnLaunch: jsii.Boolean(false),
//   		reserved: jsii.Boolean(false),
//   	},
//   	subnetConstructId: jsii.String("subnetConstructId"),
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

