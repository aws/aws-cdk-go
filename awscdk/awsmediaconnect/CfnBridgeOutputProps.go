package awsmediaconnect


// Properties for defining a `CfnBridgeOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBridgeOutputProps := &CfnBridgeOutputProps{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	Name: jsii.String("name"),
//   	NetworkOutput: &BridgeNetworkOutputProperty{
//   		IpAddress: jsii.String("ipAddress"),
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		Ttl: jsii.Number(123),
//   	},
//   }
//
type CfnBridgeOutputProps struct {
	// The ARN of the bridge that you want to describe.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The network output name.
	//
	// This name is used to reference the output and must be unique among outputs in this bridge.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Add a network output to an existing bridge.
	NetworkOutput interface{} `field:"required" json:"networkOutput" yaml:"networkOutput"`
}

