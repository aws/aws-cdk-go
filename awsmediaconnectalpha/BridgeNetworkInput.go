package awsmediaconnectalpha


// A named network source for an ingress bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var bridgeProtocol BridgeProtocol
//   var gatewayNetwork GatewayNetwork
//
//   bridgeNetworkInput := &BridgeNetworkInput{
//   	Name: jsii.String("name"),
//   	Source: &BridgeNetworkSource{
//   		MulticastIp: jsii.String("multicastIp"),
//   		Network: gatewayNetwork,
//   		Port: jsii.Number(123),
//   		Protocol: bridgeProtocol,
//
//   		// the properties below are optional
//   		MulticastSourceIp: jsii.String("multicastSourceIp"),
//   	},
//   }
//
// Experimental.
type BridgeNetworkInput struct {
	// The name of the network source.
	//
	// Must be unique among sources on the bridge.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network source configuration describing the multicast endpoint the bridge listens to.
	// Experimental.
	Source *BridgeNetworkSource `field:"required" json:"source" yaml:"source"`
}

