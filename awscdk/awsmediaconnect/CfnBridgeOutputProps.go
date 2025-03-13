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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html
//
type CfnBridgeOutputProps struct {
	// The ARN of the bridge that you want to describe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-bridgearn
	//
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The network output name.
	//
	// This name is used to reference the output and must be unique among outputs in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Add a network output to an existing bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-networkoutput
	//
	NetworkOutput interface{} `field:"required" json:"networkOutput" yaml:"networkOutput"`
}

