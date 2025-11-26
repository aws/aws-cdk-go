package previewawsmediaconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mergeRouterInputProtocolConfigurationProperty := &MergeRouterInputProtocolConfigurationProperty{
//   	Rist: &RistRouterInputConfigurationProperty{
//   		Port: jsii.Number(123),
//   		RecoveryLatencyMilliseconds: jsii.Number(123),
//   	},
//   	Rtp: &RtpRouterInputConfigurationProperty{
//   		ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   		Port: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputprotocolconfiguration.html
//
type CfnRouterInputPropsMixin_MergeRouterInputProtocolConfigurationProperty struct {
	// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-mergerouterinputprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mergerouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-mergerouterinputprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
}

