package awsmediaconnectalpha


// NDI Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkInterface NetworkInterface
//   var role Role
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   ndiDiscoveryServerConfig := &NdiDiscoveryServerConfig{
//   	DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   	VpcInterface: &VpcInterfaceConfig{
//   		Name: jsii.String("name"),
//   		Role: role,
//   		SecurityGroups: []ISecurityGroup{
//   			securityGroup,
//   		},
//   		Subnet: subnet,
//
//   		// the properties below are optional
//   		NetworkInterfaceIds: []*string{
//   			jsii.String("networkInterfaceIds"),
//   		},
//   		NetworkInterfaceType: networkInterface,
//   	},
//
//   	// the properties below are optional
//   	DiscoveryServerPort: jsii.Number(123),
//   }
//
// Experimental.
type NdiDiscoveryServerConfig struct {
	// The unique network address of the NDI discovery server.
	// Experimental.
	DiscoveryServerAddress *string `field:"required" json:"discoveryServerAddress" yaml:"discoveryServerAddress"`
	// The VPC interface that the NDI discovery server uses to reach the flow.
	// Experimental.
	VpcInterface *VpcInterfaceConfig `field:"required" json:"vpcInterface" yaml:"vpcInterface"`
	// The port for the NDI discovery server.
	//
	// Defaults to 5959 if a custom port isn't specified.
	// Default: - undefined; when omitted, MediaConnect applies 5959 at deploy time.
	//
	// Experimental.
	DiscoveryServerPort *float64 `field:"optional" json:"discoveryServerPort" yaml:"discoveryServerPort"`
}

