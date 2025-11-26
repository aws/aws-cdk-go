package previewawsmediaconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routerInputProtocolConfigurationProperty := &RouterInputProtocolConfigurationProperty{
//   	Rist: &RistRouterInputConfigurationProperty{
//   		Port: jsii.Number(123),
//   		RecoveryLatencyMilliseconds: jsii.Number(123),
//   	},
//   	Rtp: &RtpRouterInputConfigurationProperty{
//   		ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   		Port: jsii.Number(123),
//   	},
//   	SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   		DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   			EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePort: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   	},
//   	SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   		DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputprotocolconfiguration.html
//
type CfnRouterInputPropsMixin_RouterInputProtocolConfigurationProperty struct {
	// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-routerinputprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-routerinputprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-routerinputprotocolconfiguration-srtcaller
	//
	SrtCaller interface{} `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputprotocolconfiguration.html#cfn-mediaconnect-routerinput-routerinputprotocolconfiguration-srtlistener
	//
	SrtListener interface{} `field:"optional" json:"srtListener" yaml:"srtListener"`
}

