package previewawsmediaconnectmixins


// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtCallerRouterOutputConfigurationProperty := &SrtCallerRouterOutputConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPort: jsii.Number(123),
//   	EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   		EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_SrtCallerRouterOutputConfigurationProperty struct {
	// The destination IP address for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The destination port number for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration-destinationport
	//
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The minimum latency in milliseconds for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The stream ID for the SRT protocol in caller mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

