package awsmediaconnect


// Create a bridge with the egress bridge type.
//
// An egress bridge is a cloud-to-ground bridge. The content comes from an existing MediaConnect flow and is delivered to your premises.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressGatewayBridgeProperty := &EgressGatewayBridgeProperty{
//   	MaxBitrate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-egressgatewaybridge.html
//
type CfnBridge_EgressGatewayBridgeProperty struct {
	// The maximum expected bitrate (in bps) of the egress bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-egressgatewaybridge.html#cfn-mediaconnect-bridge-egressgatewaybridge-maxbitrate
	//
	MaxBitrate *float64 `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
}

