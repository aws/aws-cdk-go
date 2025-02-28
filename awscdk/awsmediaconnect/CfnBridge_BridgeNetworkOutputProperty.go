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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html
//
type CfnBridge_BridgeNetworkOutputProperty struct {
	// The network output IP Address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-ipaddress
	//
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The network output name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network output's gateway network name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-networkname
	//
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network output port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network output protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The network output TTL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgenetworkoutput.html#cfn-mediaconnect-bridge-bridgenetworkoutput-ttl
	//
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

