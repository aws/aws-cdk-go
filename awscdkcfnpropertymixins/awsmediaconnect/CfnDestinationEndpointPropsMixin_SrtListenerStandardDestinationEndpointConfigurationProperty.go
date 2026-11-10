package awsmediaconnect


// The configuration settings for a destination endpoint that uses SRT in listener mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   srtListenerStandardDestinationEndpointConfigurationProperty := &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationIngestPort: jsii.Number(123),
//   	EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html
//
type CfnDestinationEndpointPropsMixin_SrtListenerStandardDestinationEndpointConfigurationProperty struct {
	// The IP address on which the destination endpoint listens for incoming connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The port on which the destination endpoint listens for incoming connections.
	//
	// The valid range is 0 to 65531.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration-destinationingestport
	//
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration-encryptionpassphrase
	//
	EncryptionPassphrase interface{} `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The minimum latency, in milliseconds, for the SRT connection.
	//
	// The valid range is 10 to 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The stream ID that identifies the content for the SRT connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-srtlistenerstandarddestinationendpointconfiguration-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

