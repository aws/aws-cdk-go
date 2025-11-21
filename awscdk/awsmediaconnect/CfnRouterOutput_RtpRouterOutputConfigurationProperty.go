package awsmediaconnect


// The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtpRouterOutputConfigurationProperty := &RtpRouterOutputConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtprouteroutputconfiguration.html
//
type CfnRouterOutput_RtpRouterOutputConfigurationProperty struct {
	// The destination IP address for the RTP protocol in the router output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtprouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtprouteroutputconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// The destination port number for the RTP protocol in the router output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtprouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtprouteroutputconfiguration-destinationport
	//
	DestinationPort *float64 `field:"required" json:"destinationPort" yaml:"destinationPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtprouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtprouteroutputconfiguration-forwarderrorcorrection
	//
	ForwardErrorCorrection *string `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
}

