package mixinsawsmediaconnect


// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rtpRouterInputConfigurationProperty := &RtpRouterInputConfigurationProperty{
//   	ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_RtpRouterInputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtprouterinputconfiguration-forwarderrorcorrection
	//
	ForwardErrorCorrection *string `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
	// The port number used for the RTP protocol in the router input configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtprouterinputconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

