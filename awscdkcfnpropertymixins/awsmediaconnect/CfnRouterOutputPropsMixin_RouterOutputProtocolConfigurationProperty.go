package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var public_ interface{}
//
//   routerOutputProtocolConfigurationProperty := &RouterOutputProtocolConfigurationProperty{
//   	Rist: &RistRouterOutputConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPort: jsii.Number(123),
//   	},
//   	RtmpPush: &RtmpPushRouterOutputConfigurationProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPort: jsii.Number(123),
//   		StreamName: jsii.String("streamName"),
//   		TlsEncryption: &TlsEncryptionProperty{
//   			EncryptionConfiguration: &TlsEncryptionConfigurationProperty{
//   				Public: public_,
//   			},
//   			EncryptionType: jsii.String("encryptionType"),
//   		},
//   	},
//   	Rtp: &RtpRouterOutputConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPort: jsii.Number(123),
//   		ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   	},
//   	SrtCaller: &SrtCallerRouterOutputConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPort: jsii.Number(123),
//   		EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   			EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   	},
//   	SrtListener: &SrtListenerRouterOutputConfigurationProperty{
//   		EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   			EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		Port: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html
//
type CfnRouterOutputPropsMixin_RouterOutputProtocolConfigurationProperty struct {
	// The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a router output that pushes a stream to a destination using the RTMP (Real-Time Messaging Protocol) protocol, or RTMPS (RTMP over TLS) when TLS encryption is specified.
	//
	// These settings include the destination address and port, the application and stream names, and optional TLS encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputprotocolconfiguration-rtmppush
	//
	RtmpPush interface{} `field:"optional" json:"rtmpPush" yaml:"rtmpPush"`
	// The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputprotocolconfiguration-srtcaller
	//
	SrtCaller interface{} `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-routeroutputprotocolconfiguration.html#cfn-mediaconnect-routeroutput-routeroutputprotocolconfiguration-srtlistener
	//
	SrtListener interface{} `field:"optional" json:"srtListener" yaml:"srtListener"`
}

