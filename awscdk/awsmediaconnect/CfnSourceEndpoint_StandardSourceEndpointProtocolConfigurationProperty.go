package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardSourceEndpointProtocolConfigurationProperty := &StandardSourceEndpointProtocolConfigurationProperty{
//   	Rist: &RistStandardSourceEndpointConfigurationProperty{
//   		DestinationIngestPort: jsii.Number(123),
//   		RecoveryLatencyMilliseconds: jsii.Number(123),
//   	},
//   	Rtp: &RtpStandardSourceEndpointConfigurationProperty{
//   		DestinationIngestPort: jsii.Number(123),
//   		ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   	},
//   	SrtCaller: &SrtCallerStandardSourceEndpointConfigurationProperty{
//   		DestinationIngestPort: jsii.Number(123),
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourceListenerPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		StreamId: jsii.String("streamId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration.html
//
type CfnSourceEndpoint_StandardSourceEndpointProtocolConfigurationProperty struct {
	// The configuration settings for a source endpoint that uses the RIST protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a source endpoint that uses the RTP protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a source endpoint that uses SRT in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration-srtcaller
	//
	SrtCaller interface{} `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a source endpoint that uses SRT in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointprotocolconfiguration-srtlistener
	//
	SrtListener interface{} `field:"optional" json:"srtListener" yaml:"srtListener"`
}

