package mixinsawsmedialive


// The settings for one RTMP output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rtmpOutputSettingsProperty := &RtmpOutputSettingsProperty{
//   	CertificateMode: jsii.String("certificateMode"),
//   	ConnectionRetryInterval: jsii.Number(123),
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	NumRetries: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpoutputsettings.html
//
type CfnChannelPropsMixin_RtmpOutputSettingsProperty struct {
	// If set to verifyAuthenticity, verifies the TLS certificate chain to a trusted certificate authority (CA).
	//
	// This causes RTMPS outputs with self-signed certificates to fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpoutputsettings.html#cfn-medialive-channel-rtmpoutputsettings-certificatemode
	//
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying a connection to the Flash Media server if the connection is lost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpoutputsettings.html#cfn-medialive-channel-rtmpoutputsettings-connectionretryinterval
	//
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The RTMP endpoint excluding the stream name (for example, rtmp://host/appname).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpoutputsettings.html#cfn-medialive-channel-rtmpoutputsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of retry attempts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpoutputsettings.html#cfn-medialive-channel-rtmpoutputsettings-numretries
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}

