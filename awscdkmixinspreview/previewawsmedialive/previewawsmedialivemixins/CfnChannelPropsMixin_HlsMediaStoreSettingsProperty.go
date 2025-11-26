package previewawsmedialivemixins


// The configuration of a MediaStore container as the destination for an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsMediaStoreSettingsProperty := &HlsMediaStoreSettingsProperty{
//   	ConnectionRetryInterval: jsii.Number(123),
//   	FilecacheDuration: jsii.Number(123),
//   	MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   	NumRetries: jsii.Number(123),
//   	RestartDelay: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html
//
type CfnChannelPropsMixin_HlsMediaStoreSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html#cfn-medialive-channel-hlsmediastoresettings-connectionretryinterval
	//
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html#cfn-medialive-channel-hlsmediastoresettings-filecacheduration
	//
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// When set to temporal, output files are stored in non-persistent memory for faster reading and writing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html#cfn-medialive-channel-hlsmediastoresettings-mediastorestorageclass
	//
	MediaStoreStorageClass *string `field:"optional" json:"mediaStoreStorageClass" yaml:"mediaStoreStorageClass"`
	// The number of retry attempts that are made before the channel is put into an error state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html#cfn-medialive-channel-hlsmediastoresettings-numretries
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsmediastoresettings.html#cfn-medialive-channel-hlsmediastoresettings-restartdelay
	//
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

