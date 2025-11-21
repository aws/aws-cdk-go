package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverRouterInputProtocolConfigurationProperty := &FailoverRouterInputProtocolConfigurationProperty{
//   	Rist: &RistRouterInputConfigurationProperty{
//   		Port: jsii.Number(123),
//   		RecoveryLatencyMilliseconds: jsii.Number(123),
//   	},
//   	Rtp: &RtpRouterInputConfigurationProperty{
//   		Port: jsii.Number(123),
//
//   		// the properties below are optional
//   		ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   	},
//   	SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePort: jsii.Number(123),
//
//   		// the properties below are optional
//   		DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   			EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		StreamId: jsii.String("streamId"),
//   	},
//   	SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		Port: jsii.Number(123),
//
//   		// the properties below are optional
//   		DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   			EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.html
//
type CfnRouterInput_FailoverRouterInputProtocolConfigurationProperty struct {
	// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration-srtcaller
	//
	SrtCaller interface{} `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration-srtlistener
	//
	SrtListener interface{} `field:"optional" json:"srtListener" yaml:"srtListener"`
}

