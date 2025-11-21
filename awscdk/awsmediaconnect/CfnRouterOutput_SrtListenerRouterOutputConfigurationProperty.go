package awsmediaconnect


// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtListenerRouterOutputConfigurationProperty := &SrtListenerRouterOutputConfigurationProperty{
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	EncryptionConfiguration: &SrtEncryptionConfigurationProperty{
//   		EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html
//
type CfnRouterOutput_SrtListenerRouterOutputConfigurationProperty struct {
	// The minimum latency in milliseconds for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"required" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port number for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

