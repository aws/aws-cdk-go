package awsec2


// Options passed by the VPC when NAT needs to be configured.
//
// Don't use this directly, the VPC will pass it automatically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var privateSubnet PrivateSubnet
//   var publicSubnet PublicSubnet
//   var vpc Vpc
//
//   configureNatOptions := &ConfigureNatOptions{
//   	NatSubnets: []PublicSubnet{
//   		publicSubnet,
//   	},
//   	PrivateSubnets: []PrivateSubnet{
//   		privateSubnet,
//   	},
//   	Vpc: vpc,
//   }
//
type ConfigureNatOptions struct {
	// The public subnets where the NAT providers need to be placed.
	//
	// For Regional NAT Gateways, this should be an empty array as they
	// don't require public subnets.
	NatSubnets *[]PublicSubnet `field:"required" json:"natSubnets" yaml:"natSubnets"`
	// The private subnets that need to route through the NAT providers.
	//
	// There may be more private subnets than public subnets with NAT providers.
	PrivateSubnets *[]PrivateSubnet `field:"required" json:"privateSubnets" yaml:"privateSubnets"`
	// The VPC we're configuring NAT for.
	Vpc Vpc `field:"required" json:"vpc" yaml:"vpc"`
}

