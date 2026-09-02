package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an Archive (S3) output group.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Archive(&ArchiveOutputGroupProps{
//   	Name: jsii.String("archive"),
//   	Destinations: []S3OutputDestination{
//   		medialive.S3OutputDestination_ToBucket(bucket, jsii.String("archive/recording")),
//   	},
//   	RolloverInterval: awscdk.Duration_Seconds(jsii.Number(600)),
//   	Outputs: []ArchiveOutputDefinition{
//   		&ArchiveOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("archive_out"),
//   		},
//   	},
//   })
//
// Experimental.
type ArchiveOutputGroupProps struct {
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
	// The S3 canned ACL to apply to each archive output.
	// Default: - no canned ACL.
	//
	// Experimental.
	ArchiveS3CannedAcl S3CannedAcl `field:"optional" json:"archiveS3CannedAcl" yaml:"archiveS3CannedAcl"`
	// The outputs for this Archive output group.
	// Default: - no initial outputs.
	//
	// Experimental.
	Outputs *[]*ArchiveOutputDefinition `field:"optional" json:"outputs" yaml:"outputs"`
	// The duration of each archive file (rollover interval).
	// Default: Duration.seconds(300)
	//
	// Experimental.
	RolloverInterval awscdk.Duration `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

