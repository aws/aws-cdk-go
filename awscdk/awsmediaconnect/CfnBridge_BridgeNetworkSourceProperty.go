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
//   	Name: jsii.String("name"),
//   	NetworkName: jsii.String("networkName"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   		MulticastSourceIp: jsii.String("multicastSourceIp"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html
//
type CfnBridge_BridgeNetworkSourceProperty struct {
	// The network source multicast IP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-multicastip
	//
	MulticastIp *string `field:"required" json:"multicastIp" yaml:"multicastIp"`
	// The name of the network source.
	//
	// This name is used to reference the source and must be unique among sources in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network source's gateway network name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-networkname
	//
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network source port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network source protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The settings related to the multicast source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworksource.html#cfn-mediaconnect-bridge-bridgenetworksource-multicastsourcesettings
	//
	MulticastSourceSettings interface{} `field:"optional" json:"multicastSourceSettings" yaml:"multicastSourceSettings"`
}

