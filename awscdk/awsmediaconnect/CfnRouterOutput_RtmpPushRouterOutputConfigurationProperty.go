package awsmediaconnect


// The configuration settings for a router output that pushes a stream to a destination using the RTMP (Real-Time Messaging Protocol) protocol, or RTMPS (RTMP over TLS) when TLS encryption is specified.
//
// These settings include the destination address and port, the application and stream names, and optional TLS encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var public_ interface{}
//
//   rtmpPushRouterOutputConfigurationProperty := &RtmpPushRouterOutputConfigurationProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPort: jsii.Number(123),
//   	StreamName: jsii.String("streamName"),
//
//   	// the properties below are optional
//   	TlsEncryption: &TlsEncryptionProperty{
//   		EncryptionConfiguration: &TlsEncryptionConfigurationProperty{
//   			Public: public_,
//   		},
//
//   		// the properties below are optional
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html
//
type CfnRouterOutput_RtmpPushRouterOutputConfigurationProperty struct {
	// The name of the RTMP application on the destination server.
	//
	// Together with the stream name, the application name forms the RTMP URL path, in the pattern rtmp://destinationAddress/applicationName/streamName.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration-applicationname
	//
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The IP address or hostname of the destination RTMP server that the router output pushes the stream to.
	//
	// Provide only the server address; specify the application and stream names separately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// The TCP port on the destination RTMP server.
	//
	// For RTMP, valid values range from 1024 to 65535. For RTMPS (RTMP over TLS), valid values are 443 or 1024 to 65535. RTMP typically uses port 1935, and RTMPS typically uses port 443.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration-destinationport
	//
	DestinationPort *float64 `field:"required" json:"destinationPort" yaml:"destinationPort"`
	// The name of the RTMP stream that the output publishes to the destination application.
	//
	// The stream name forms the final segment of the RTMP URL path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration-streamname
	//
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// The Transport Layer Security (TLS) encryption settings used to establish a secure connection to a destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-rtmppushrouteroutputconfiguration-tlsencryption
	//
	TlsEncryption interface{} `field:"optional" json:"tlsEncryption" yaml:"tlsEncryption"`
}

