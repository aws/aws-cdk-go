package awsmedialive


// The configuration of a MediaStore container as the destination for an HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsMediaStoreSettingsProperty := &hlsMediaStoreSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	mediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_HlsMediaStoreSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// When set to temporal, output files are stored in non-persistent memory for faster reading and writing.
	MediaStoreStorageClass *string `field:"optional" json:"mediaStoreStorageClass" yaml:"mediaStoreStorageClass"`
	// The number of retry attempts that are made before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

