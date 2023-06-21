package awsmediaconnect


// The output of the bridge.
//
// A network output is delivered to your premises.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeNetworkOutputProperty := &BridgeNetworkOutputProperty{
//   	IpAddress: jsii.String("ipAddress"),
//   	Name: jsii.String("name"),
//   	NetworkName: jsii.String("networkName"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	Ttl: jsii.Number(123),
//   }
//
type CfnBridge_BridgeNetworkOutputProperty struct {
	// The network output IP Address.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The network output name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network output's gateway network name.
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network output port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network output protocol.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The network output TTL.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

