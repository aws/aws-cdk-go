package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for VBR rate control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitrate Bitrate
//
//   vbrRateControlProps := &VbrRateControlProps{
//   	Bitrate: bitrate,
//   	MaxBitrate: bitrate,
//   }
//
// Experimental.
type VbrRateControlProps struct {
	// The average bitrate.
	// Experimental.
	Bitrate awscdk.Bitrate `field:"required" json:"bitrate" yaml:"bitrate"`
	// The maximum bitrate.
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
}

