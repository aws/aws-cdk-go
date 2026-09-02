package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for M3U8 container settings.
//
// PID properties accept a decimal or hexadecimal value (and, where noted, ranges or
// comma-separated lists). Interval properties are `Duration` values rendered as milliseconds.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   	},
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("video"),
//   			HlsSettings: medialive.HlsSettings_Standard(&StandardHlsSettingsProps{
//   				M3u8Settings: medialive.M3u8Settings_Of(&M3u8SettingsProps{
//   					Scte35Behavior: medialive.M3u8Scte35Behavior_PASSTHROUGH(),
//   					ProgramNum: jsii.Number(1),
//   				}),
//   			}),
//   		},
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				audio,
//   			},
//   			OutputName: jsii.String("audio"),
//   			HlsSettings: medialive.HlsSettings_AudioOnly(&AudioOnlyHlsSettingsProps{
//   				AudioGroupId: jsii.String("program"),
//   				AudioOnlyImage: medialive.FileLocation_FromBucket(bucket, jsii.String("art/cover.png")),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type M3u8SettingsProps struct {
	// The number of audio frames to insert for each PES packet.
	// Default: - service default.
	//
	// Experimental.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID(s) of the elementary audio streams.
	//
	// Accepts ranges and comma separation.
	// Default: - service default.
	//
	// Experimental.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// KLV data passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	KlvBehavior M3u8KlvBehavior `field:"optional" json:"klvBehavior" yaml:"klvBehavior"`
	// The PID(s) of the KLV data streams.
	// Default: - service default.
	//
	// Experimental.
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// Nielsen ID3 passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	NielsenId3Behavior M3u8NielsenId3Behavior `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The interval between instances of the PAT in the output.
	//
	// A value of 0 writes the PAT once
	// per segment file.
	// Default: - service default.
	//
	// Experimental.
	PatInterval awscdk.Duration `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Controls insertion of the Program Clock Reference (PCR).
	// Default: - service default.
	//
	// Experimental.
	PcrControl M3u8PcrControl `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum interval between Program Clock References (PCRs).
	// Default: - service default.
	//
	// Experimental.
	PcrPeriod awscdk.Duration `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR).
	//
	// Defaults to the video PID.
	// Default: - same as the video PID.
	//
	// Experimental.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The interval between instances of the PMT in the output.
	//
	// A value of 0 writes the PMT once
	// per segment file.
	// Default: - service default.
	//
	// Experimental.
	PmtInterval awscdk.Duration `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID of the Program Map Table (PMT).
	// Default: - service default.
	//
	// Experimental.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the PMT.
	// Default: - service default.
	//
	// Experimental.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// SCTE-35 passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	Scte35Behavior M3u8Scte35Behavior `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// The PID of the SCTE-35 stream.
	// Default: - service default.
	//
	// Experimental.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Timed-metadata passthrough behavior.
	// Default: - service default.
	//
	// Experimental.
	TimedMetadataBehavior M3u8TimedMetadataBehavior `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed-metadata stream.
	//
	// Valid values are 32 (0x20)..8182 (0x1ff6).
	// Default: - service default.
	//
	// Experimental.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the PMT.
	// Default: - service default.
	//
	// Experimental.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream.
	// Default: - service default.
	//
	// Experimental.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

