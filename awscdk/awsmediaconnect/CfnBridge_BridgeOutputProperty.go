package awsmediaconnect


// The output of the bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeOutputProperty := &BridgeOutputProperty{
//   	NetworkOutput: &BridgeNetworkOutputProperty{
//   		IpAddress: jsii.String("ipAddress"),
//   		Name: jsii.String("name"),
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		Ttl: jsii.Number(123),
//   	},
//   }
//
type CfnBridge_BridgeOutputProperty struct {
	// The output of the bridge.
	//
	// A network output is delivered to your premises.
	NetworkOutput interface{} `field:"optional" json:"networkOutput" yaml:"networkOutput"`
}

