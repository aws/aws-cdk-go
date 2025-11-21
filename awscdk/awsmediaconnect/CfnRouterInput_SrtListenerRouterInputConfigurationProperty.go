package awsmediaconnect


// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtListenerRouterInputConfigurationProperty := &SrtListenerRouterInputConfigurationProperty{
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   		EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtlistenerrouterinputconfiguration.html
//
type CfnRouterInput_SrtListenerRouterInputConfigurationProperty struct {
	// The minimum latency in milliseconds for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtlistenerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtlistenerrouterinputconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"required" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port number for the SRT protocol in listener mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtlistenerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtlistenerrouterinputconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-srtlistenerrouterinputconfiguration.html#cfn-mediaconnect-routerinput-srtlistenerrouterinputconfiguration-decryptionconfiguration
	//
	DecryptionConfiguration interface{} `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
}

