package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for EAC3 Atmos codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitrate Bitrate
//   var eac3AtmosCodingMode Eac3AtmosCodingMode
//   var eac3AtmosDrcLine Eac3AtmosDrcLine
//   var eac3AtmosDrcRf Eac3AtmosDrcRf
//
//   eac3AtmosSettingsProps := &Eac3AtmosSettingsProps{
//   	Bitrate: bitrate,
//   	CodingMode: eac3AtmosCodingMode,
//   	DialNorm: jsii.Number(123),
//   	DrcLine: eac3AtmosDrcLine,
//   	DrcRf: eac3AtmosDrcRf,
//   	HeightTrim: jsii.Number(123),
//   	SurroundTrim: jsii.Number(123),
//   }
//
// Experimental.
type Eac3AtmosSettingsProps struct {
	// The average bitrate.
	// Default: - service default.
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The coding mode (e.g. CODING_MODE_5_1_4, CODING_MODE_7_1_4, CODING_MODE_9_1_6).
	// Default: Eac3AtmosCodingMode.CODING_MODE_5_1_4
	//
	// Experimental.
	CodingMode Eac3AtmosCodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The dialogue normalization level (1–31).
	// Default: - service default.
	//
	// Experimental.
	DialNorm *float64 `field:"optional" json:"dialNorm" yaml:"dialNorm"`
	// Sets the Dolby dynamic range compression line mode profile.
	// Default: - service default.
	//
	// Experimental.
	DrcLine Eac3AtmosDrcLine `field:"optional" json:"drcLine" yaml:"drcLine"`
	// Sets the Dolby dynamic range compression RF mode profile.
	// Default: - service default.
	//
	// Experimental.
	DrcRf Eac3AtmosDrcRf `field:"optional" json:"drcRf" yaml:"drcRf"`
	// Height channel trim level.
	// Default: - service default.
	//
	// Experimental.
	HeightTrim *float64 `field:"optional" json:"heightTrim" yaml:"heightTrim"`
	// Surround channel trim level.
	// Default: - service default.
	//
	// Experimental.
	SurroundTrim *float64 `field:"optional" json:"surroundTrim" yaml:"surroundTrim"`
}

