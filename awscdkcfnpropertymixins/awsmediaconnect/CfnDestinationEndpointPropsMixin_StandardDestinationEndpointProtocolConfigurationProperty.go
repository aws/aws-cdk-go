package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   standardDestinationEndpointProtocolConfigurationProperty := &StandardDestinationEndpointProtocolConfigurationProperty{
//   	Rist: &RistStandardDestinationEndpointConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationIngestPort: jsii.Number(123),
//   	},
//   	Rtp: &RtpStandardDestinationEndpointConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationIngestPort: jsii.Number(123),
//   		ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   	},
//   	SrtCaller: &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   		EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		SourceListenerPort: jsii.Number(123),
//   	},
//   	SrtListener: &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationIngestPort: jsii.Number(123),
//   		EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		MinimumLatencyMilliseconds: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration.html
//
type CfnDestinationEndpointPropsMixin_StandardDestinationEndpointProtocolConfigurationProperty struct {
	// The configuration settings for a destination endpoint that uses the RIST protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration-rist
	//
	Rist interface{} `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a destination endpoint that uses the RTP protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration-rtp
	//
	Rtp interface{} `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a destination endpoint that uses SRT in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration-srtcaller
	//
	SrtCaller interface{} `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a destination endpoint that uses SRT in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointprotocolconfiguration-srtlistener
	//
	SrtListener interface{} `field:"optional" json:"srtListener" yaml:"srtListener"`
}

