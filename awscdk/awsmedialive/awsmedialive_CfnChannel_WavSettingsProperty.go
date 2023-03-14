package awsmedialive


// The setup of WAV audio in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wavSettingsProperty := &wavSettingsProperty{
//   	bitDepth: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	sampleRate: jsii.Number(123),
//   }
//
type CfnChannel_WavSettingsProperty struct {
	// Bits per sample.
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// The audio coding mode for the WAV audio.
	//
	// The mode determines the number of channels in the audio.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sample rate in Hz.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

