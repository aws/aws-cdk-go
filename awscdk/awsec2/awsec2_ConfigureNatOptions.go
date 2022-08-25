package awsec2


// Options passed by the VPC when NAT needs to be configured.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var privateSubnet privateSubnet
//   var publicSubnet publicSubnet
//   var vpc vpc
//
//   configureNatOptions := &configureNatOptions{
//   	natSubnets: []*publicSubnet{
//   		publicSubnet,
//   	},
//   	privateSubnets: []*privateSubnet{
//   		privateSubnet,
//   	},
//   	vpc: vpc,
//   }
//
type ConfigureNatOptions struct {
	// The public subnets where the NAT providers need to be placed.
	NatSubnets *[]PublicSubnet `field:"required" json:"natSubnets" yaml:"natSubnets"`
	// The private subnets that need to route through the NAT providers.
	//
	// There may be more private subnets than public subnets with NAT providers.
	PrivateSubnets *[]PrivateSubnet `field:"required" json:"privateSubnets" yaml:"privateSubnets"`
	// The VPC we're configuring NAT for.
	Vpc Vpc `field:"required" json:"vpc" yaml:"vpc"`
}

