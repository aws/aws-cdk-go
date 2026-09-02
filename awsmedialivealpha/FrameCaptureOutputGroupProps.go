package awsmedialivealpha


// Properties for a Frame Capture output group.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_FrameCapture(&FrameCaptureOutputGroupProps{
//   	Name: jsii.String("thumbnails"),
//   	Destinations: []S3OutputDestination{
//   		medialive.S3OutputDestination_ToBucket(bucket, jsii.String("thumbnails/live")),
//   	},
//   	Outputs: []FrameCaptureOutputDefinition{
//   		&FrameCaptureOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("thumb"),
//   		},
//   	},
//   })
//
// Experimental.
type FrameCaptureOutputGroupProps struct {
	// The destinations for this output group — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations.
	// Experimental.
	Destinations *[]S3OutputDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The S3 canned ACL to apply to each frame capture output.
	// Default: - no canned ACL.
	//
	// Experimental.
	FrameCaptureS3CannedAcl S3CannedAcl `field:"optional" json:"frameCaptureS3CannedAcl" yaml:"frameCaptureS3CannedAcl"`
	// The outputs for this Frame Capture output group.
	// Default: - no initial outputs.
	//
	// Experimental.
	Outputs *[]*FrameCaptureOutputDefinition `field:"optional" json:"outputs" yaml:"outputs"`
}

