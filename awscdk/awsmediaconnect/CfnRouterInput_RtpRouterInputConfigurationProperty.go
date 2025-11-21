package awsmediaconnect


// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtpRouterInputConfigurationProperty := &RtpRouterInputConfigurationProperty{
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html
//
type CfnRouterInput_RtpRouterInputConfigurationProperty struct {
	// The port number used for the RTP protocol in the router input configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtprouterinputconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.html#cfn-mediaconnect-routerinput-rtprouterinputconfiguration-forwarderrorcorrection
	//
	ForwardErrorCorrection *string `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
}

