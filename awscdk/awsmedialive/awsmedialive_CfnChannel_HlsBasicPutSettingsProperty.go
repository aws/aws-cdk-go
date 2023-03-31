package awsmedialive


// The configuration of HLS Basic Put Settings.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsBasicPutSettingsProperty := &hlsBasicPutSettingsProperty{
//   	connectionRetryInterval: jsii.Number(123),
//   	filecacheDuration: jsii.Number(123),
//   	numRetries: jsii.Number(123),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_HlsBasicPutSettingsProperty struct {
	// The number of seconds to wait before retrying a connection to the CDN if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size, in seconds, of the file cache for streaming outputs.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The number of retry attempts that MediaLive makes before the channel is put into an error state.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

