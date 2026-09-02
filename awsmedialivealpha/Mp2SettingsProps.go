package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for MP2 codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var audioSampleRate AudioSampleRate
//   var bitrate Bitrate
//   var mp2CodingMode Mp2CodingMode
//
//   mp2SettingsProps := &Mp2SettingsProps{
//   	Bitrate: bitrate,
//   	CodingMode: mp2CodingMode,
//   	SampleRate: audioSampleRate,
//   }
//
// Experimental.
type Mp2SettingsProps struct {
	// The average bitrate.
	// Default: - service default.
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The MPEG2 Audio coding mode.
	// Default: Mp2CodingMode.CODING_MODE_2_0
	//
	// Experimental.
	CodingMode Mp2CodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The sample rate.
	// Default: AudioSampleRate.HZ_48000
	//
	// Experimental.
	SampleRate AudioSampleRate `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

