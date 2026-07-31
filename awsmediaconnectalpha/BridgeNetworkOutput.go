package awsmediaconnectalpha


// A named network output for an egress bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var bridgeOutputConfiguration BridgeOutputConfiguration
//
//   bridgeNetworkOutput := &BridgeNetworkOutput{
//   	Name: jsii.String("name"),
//   	Output: bridgeOutputConfiguration,
//   }
//
// Experimental.
type BridgeNetworkOutput struct {
	// The name of the network output. Must be unique among outputs on the bridge.
	//
	// Used as the physical name of the underlying CFN resource.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network configuration describing where this output sends content.
	// Experimental.
	Output BridgeOutputConfiguration `field:"required" json:"output" yaml:"output"`
}

