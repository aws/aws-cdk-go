package awsmediaconnect


// The settings for an output that pushes a stream to a destination using the RTMP (Real-Time Messaging Protocol) protocol, or RTMPS (RTMP over TLS) when TLS encryption is specified.
//
// These settings include the application and stream names, and optional TLS encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var public_ interface{}
//
//   rtmpPushSettingsProperty := &RtmpPushSettingsProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	StreamName: jsii.String("streamName"),
//   	TlsEncryption: &TlsEncryptionProperty{
//   		EncryptionConfiguration: &TlsEncryptionConfigurationProperty{
//   			Public: public_,
//   		},
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-rtmppushsettings.html
//
type CfnFlowOutputPropsMixin_RtmpPushSettingsProperty struct {
	// The name of the RTMP application on the destination server.
	//
	// Together with the stream name, the application name forms the RTMP URL path, in the pattern rtmp://destinationAddress/applicationName/streamName.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-rtmppushsettings.html#cfn-mediaconnect-flowoutput-rtmppushsettings-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The name of the RTMP stream that the output publishes to the destination application.
	//
	// The stream name forms the final segment of the RTMP URL path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-rtmppushsettings.html#cfn-mediaconnect-flowoutput-rtmppushsettings-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// The Transport Layer Security (TLS) encryption settings used to establish a secure connection to a destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-rtmppushsettings.html#cfn-mediaconnect-flowoutput-rtmppushsettings-tlsencryption
	//
	TlsEncryption interface{} `field:"optional" json:"tlsEncryption" yaml:"tlsEncryption"`
}

