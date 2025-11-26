package previewawsmediaconnectmixins


// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtListenerRouterOutputConfigurationProperty := &SrtListenerRouterOutputConfigurationProperty{
//   	EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   		EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_SrtListenerRouterOutputConfigurationProperty struct {
	// Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The minimum latency in milliseconds for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port number for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

