package awsmediaconnect


// Create a bridge with the ingress bridge type.
//
// An ingress bridge is a ground-to-cloud bridge. The content originates at your premises and is delivered to the cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressGatewayBridgeProperty := &IngressGatewayBridgeProperty{
//   	MaxBitrate: jsii.Number(123),
//   	MaxOutputs: jsii.Number(123),
//   }
//
type CfnBridge_IngressGatewayBridgeProperty struct {
	// The maximum expected bitrate (in bps) of the ingress bridge.
	MaxBitrate *float64 `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum number of outputs on the ingress bridge.
	MaxOutputs *float64 `field:"required" json:"maxOutputs" yaml:"maxOutputs"`
}

