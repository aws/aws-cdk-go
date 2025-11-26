package previewawsmediaconnectmixins


// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtCallerRouterInputConfigurationProperty := &SrtCallerRouterInputConfigurationProperty{
//   	DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   		EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	SourceAddress: jsii.String("sourceAddress"),
//   	SourcePort: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_SrtCallerRouterInputConfigurationProperty struct {
	// Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtcallerrouterinputconfiguration-decryptionconfiguration
	//
	DecryptionConfiguration interface{} `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
	// The minimum latency in milliseconds for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtcallerrouterinputconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The source IP address for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtcallerrouterinputconfiguration-sourceaddress
	//
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// The source port number for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtcallerrouterinputconfiguration-sourceport
	//
	SourcePort *float64 `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// The stream ID for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtcallerrouterinputconfiguration-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

