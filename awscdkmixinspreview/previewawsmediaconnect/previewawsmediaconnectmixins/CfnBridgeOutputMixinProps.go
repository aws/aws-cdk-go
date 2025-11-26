package previewawsmediaconnectmixins


// Properties for CfnBridgeOutputPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBridgeOutputMixinProps := &CfnBridgeOutputMixinProps{
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
type CfnBridgeOutputMixinProps struct {
	// The Amazon Resource Name (ARN) of the bridge that you want to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-bridgearn
	//
	BridgeArn *string `field:"optional" json:"bridgeArn" yaml:"bridgeArn"`
	// The network output name.
	//
	// This name is used to reference the output and must be unique among outputs in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network output of the bridge.
	//
	// A network output is delivered to your premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgeoutput.html#cfn-mediaconnect-bridgeoutput-networkoutput
	//
	NetworkOutput interface{} `field:"optional" json:"networkOutput" yaml:"networkOutput"`
}

