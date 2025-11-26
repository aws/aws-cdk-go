package previewawsmedialivemixins


// The configuration of HLS Basic Put Settings.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsBasicPutSettingsProperty := &HlsBasicPutSettingsProperty{
//   	ConnectionRetryInterval: jsii.Number(123),
//   	FilecacheDuration: jsii.Number(123),
//   	NumRetries: jsii.Number(123),
//   	RestartDelay: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsbasicputsettings.html
//
type CfnChannelPropsMixin_HlsBasicPutSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsbasicputsettings.html#cfn-medialive-channel-hlsbasicputsettings-connectionretryinterval
	//
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsbasicputsettings.html#cfn-medialive-channel-hlsbasicputsettings-filecacheduration
	//
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The number of retry attempts that MediaLive makes before the channel is put into an error state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsbasicputsettings.html#cfn-medialive-channel-hlsbasicputsettings-numretries
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsbasicputsettings.html#cfn-medialive-channel-hlsbasicputsettings-restartdelay
	//
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

