package awsmediaconnect


// The configuration settings for a source endpoint that uses SRT in listener mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtListenerStandardSourceEndpointConfigurationProperty := &SrtListenerStandardSourceEndpointConfigurationProperty{
//   	MinimumLatencyMilliseconds: jsii.Number(123),
//   	SourceAddress: jsii.String("sourceAddress"),
//   	SourceListenerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html
//
type CfnSourceEndpoint_SrtListenerStandardSourceEndpointConfigurationProperty struct {
	// The minimum latency, in milliseconds, for the SRT connection.
	//
	// The valid range is 10 to 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration-minimumlatencymilliseconds
	//
	MinimumLatencyMilliseconds *float64 `field:"required" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The IP address that the source endpoint listens on for incoming connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration-sourceaddress
	//
	SourceAddress *string `field:"required" json:"sourceAddress" yaml:"sourceAddress"`
	// The port that the source endpoint listens on for incoming connections.
	//
	// The valid range is 1024 to 65535.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration-sourcelistenerport
	//
	SourceListenerPort *float64 `field:"required" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration-encryptionpassphrase
	//
	EncryptionPassphrase interface{} `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The stream ID that identifies the content for the SRT connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-srtlistenerstandardsourceendpointconfiguration-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

