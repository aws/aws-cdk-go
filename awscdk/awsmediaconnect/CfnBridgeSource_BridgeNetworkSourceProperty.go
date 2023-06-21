package awsmediaconnect


// The source of the bridge.
//
// A network source originates at your premises.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeNetworkSourceProperty := &BridgeNetworkSourceProperty{
//   	MulticastIp: jsii.String("multicastIp"),
//   	NetworkName: jsii.String("networkName"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   }
//
type CfnBridgeSource_BridgeNetworkSourceProperty struct {
	// The network source multicast IP.
	MulticastIp *string `field:"required" json:"multicastIp" yaml:"multicastIp"`
	// The network source's gateway network name.
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network source port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network source protocol.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

