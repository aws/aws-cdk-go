package awsmediaconnectalpha


// A named flow source for an egress bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var flow Flow
//   var networkInterface NetworkInterface
//   var role Role
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   bridgeFlowInput := &BridgeFlowInput{
//   	Name: jsii.String("name"),
//   	Source: &BridgeFlowSource{
//   		Flow: flow,
//
//   		// the properties below are optional
//   		VpcInterface: &VpcInterfaceConfig{
//   			Name: jsii.String("name"),
//   			Role: role,
//   			SecurityGroups: []ISecurityGroup{
//   				securityGroup,
//   			},
//   			Subnet: subnet,
//
//   			// the properties below are optional
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: networkInterface,
//   		},
//   	},
//   }
//
// Experimental.
type BridgeFlowInput struct {
	// The name of the flow source.
	//
	// Must be unique among sources on the bridge.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The flow source configuration describing where the bridge consumes content from.
	// Experimental.
	Source *BridgeFlowSource `field:"required" json:"source" yaml:"source"`
}

