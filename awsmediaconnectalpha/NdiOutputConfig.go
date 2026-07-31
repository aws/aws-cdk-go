package awsmediaconnectalpha


// Configuration options for NDI outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   ndiOutputConfig := &NdiOutputConfig{
//   	NdiProgramName: jsii.String("ndiProgramName"),
//   	NdiSpeedHqQuality: jsii.Number(123),
//   }
//
// Experimental.
type NdiOutputConfig struct {
	// A suffix for the names of the NDI sources that the flow creates.
	// Default: - the output name is used.
	//
	// Experimental.
	NdiProgramName *string `field:"optional" json:"ndiProgramName" yaml:"ndiProgramName"`
	// A quality setting for the NDI Speed HQ encoder, expressed as a percentage.
	//
	// Valid range: 100-200. Higher values produce higher quality and larger bitrate.
	// See: https://aws.amazon.com/about-aws/whats-new/2025/03/aws-elemental-mediaconnect-support-ndi-outputs/
	//
	// Default: - the MediaConnect service default.
	//
	// Experimental.
	NdiSpeedHqQuality *float64 `field:"optional" json:"ndiSpeedHqQuality" yaml:"ndiSpeedHqQuality"`
}

