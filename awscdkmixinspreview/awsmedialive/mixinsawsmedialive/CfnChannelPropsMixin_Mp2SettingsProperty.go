package mixinsawsmedialive


// The configuration for this MP2 audio.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mp2SettingsProperty := &Mp2SettingsProperty{
//   	Bitrate: jsii.Number(123),
//   	CodingMode: jsii.String("codingMode"),
//   	SampleRate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html
//
type CfnChannelPropsMixin_Mp2SettingsProperty struct {
	// The average bitrate in bits/second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The MPEG2 Audio coding mode.
	//
	// Valid values are codingMode10 (for mono) or codingMode20 (for stereo).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-codingmode
	//
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The sample rate in Hz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-samplerate
	//
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

