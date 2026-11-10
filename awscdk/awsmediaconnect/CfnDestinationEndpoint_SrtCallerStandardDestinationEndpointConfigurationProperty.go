package awsmediaconnect


// The configuration settings for a destination endpoint that uses SRT in caller mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtCallerStandardDestinationEndpointConfigurationProperty := &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	SourceListenerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration.html
//
type CfnDestinationEndpoint_SrtCallerStandardDestinationEndpointConfigurationProperty struct {
	// The minimum latency, in milliseconds, for the SRT connection.
	//
	// The valid range is 10 to 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"required" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port on which the source listens for the connection.
	//
	// The valid range is 3000 to 30000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration-sourcelistenerport
	//
	SourceListenerPort *float64 `field:"required" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtcallerstandarddestinationendpointconfiguration-encryptionpassphrase
	//
	EncryptionPassphrase interface{} `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
}

