package mixinsawsmediaconnect


// Create a bridge with the ingress bridge type.
//
// An ingress bridge is a ground-to-cloud bridge. The content originates at your premises and is delivered to the cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressGatewayBridgeProperty := &IngressGatewayBridgeProperty{
//   	MaxBitrate: jsii.Number(123),
//   	MaxOutputs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-ingressgatewaybridge.html
//
type CfnBridgePropsMixin_IngressGatewayBridgeProperty struct {
	// The maximum expected bitrate (in bps) of the ingress bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-ingressgatewaybridge.html#cfn-mediaconnect-bridge-ingressgatewaybridge-maxbitrate
	//
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum number of outputs on the ingress bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-ingressgatewaybridge.html#cfn-mediaconnect-bridge-ingressgatewaybridge-maxoutputs
	//
	MaxOutputs *float64 `field:"optional" json:"maxOutputs" yaml:"maxOutputs"`
}

