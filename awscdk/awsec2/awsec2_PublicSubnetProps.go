package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicSubnetProps := &publicSubnetProps{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	cidrBlock: jsii.String("cidrBlock"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	mapPublicIpOnLaunch: jsii.Boolean(false),
//   }
//
type PublicSubnetProps struct {
	// The availability zone for the subnet.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The CIDR notation for this subnet.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The VPC which this subnet is part of.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Controls if a public IP is associated to an instance at launch.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
}

