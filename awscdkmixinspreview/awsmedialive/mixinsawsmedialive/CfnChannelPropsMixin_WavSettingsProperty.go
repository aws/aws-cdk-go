package mixinsawsmedialive


// The setup of WAV audio in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   wavSettingsProperty := &WavSettingsProperty{
//   	BitDepth: jsii.Number(123),
//   	CodingMode: jsii.String("codingMode"),
//   	SampleRate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html
//
type CfnChannelPropsMixin_WavSettingsProperty struct {
	// Bits per sample.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-bitdepth
	//
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// The audio coding mode for the WAV audio.
	//
	// The mode determines the number of channels in the audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-codingmode
	//
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sample rate in Hz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-samplerate
	//
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

