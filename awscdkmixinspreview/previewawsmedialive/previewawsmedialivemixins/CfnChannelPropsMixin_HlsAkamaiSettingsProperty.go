package previewawsmedialivemixins


// The Akamai settings in an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsAkamaiSettingsProperty := &HlsAkamaiSettingsProperty{
//   	ConnectionRetryInterval: jsii.Number(123),
//   	FilecacheDuration: jsii.Number(123),
//   	HttpTransferMode: jsii.String("httpTransferMode"),
//   	NumRetries: jsii.Number(123),
//   	RestartDelay: jsii.Number(123),
//   	Salt: jsii.String("salt"),
//   	Token: jsii.String("token"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html
//
type CfnChannelPropsMixin_HlsAkamaiSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-connectionretryinterval
	//
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-filecacheduration
	//
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Specifies whether to use chunked transfer encoding to Akamai.
	//
	// To enable this feature, contact Akamai.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-httptransfermode
	//
	HttpTransferMode *string `field:"optional" json:"httpTransferMode" yaml:"httpTransferMode"`
	// The number of retry attempts that will be made before the channel is put into an error state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-numretries
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-restartdelay
	//
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// The salt for authenticated Akamai.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-salt
	//
	Salt *string `field:"optional" json:"salt" yaml:"salt"`
	// The token parameter for authenticated Akamai.
	//
	// If this is not specified, _gda_ is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsakamaisettings.html#cfn-medialive-channel-hlsakamaisettings-token
	//
	Token *string `field:"optional" json:"token" yaml:"token"`
}

