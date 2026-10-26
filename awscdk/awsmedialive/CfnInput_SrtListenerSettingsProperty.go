package awsmedialive


// The configuration for a source that uses SRT as the connection protocol.
//
// In terms of establishing the connection, MediaLive is the listener and the upstream system is the caller.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtListenerSettingsProperty := &SrtListenerSettingsProperty{
//   	Decryption: &SrtListenerDecryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   	},
//   	MinimumLatency: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettings.html
//
type CfnInput_SrtListenerSettingsProperty struct {
	// The decryption settings for the SRT listener source.
	//
	// Present only if the source has decryption enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettings.html#cfn-medialive-input-srtlistenersettings-decryption
	//
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// The preferred latency (in milliseconds) for implementing packet loss and recovery.
	//
	// Packet recovery is a key feature of SRT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettings.html#cfn-medialive-input-srtlistenersettings-minimumlatency
	//
	MinimumLatency *float64 `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// The stream ID, if the upstream system uses this identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettings.html#cfn-medialive-input-srtlistenersettings-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

