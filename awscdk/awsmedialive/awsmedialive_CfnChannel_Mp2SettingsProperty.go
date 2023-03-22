package awsmedialive


// The configuration for this MP2 audio.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mp2SettingsProperty := &mp2SettingsProperty{
//   	bitrate: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	sampleRate: jsii.Number(123),
//   }
//
type CfnChannel_Mp2SettingsProperty struct {
	// The average bitrate in bits/second.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The MPEG2 Audio coding mode.
	//
	// Valid values are codingMode10 (for mono) or codingMode20 (for stereo).
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The sample rate in Hz.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

