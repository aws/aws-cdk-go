package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &CfnChannelProps{
//   	Authorized: jsii.Boolean(false),
//   	ContainerFormat: jsii.String("containerFormat"),
//   	InsecureIngest: jsii.Boolean(false),
//   	LatencyMode: jsii.String("latencyMode"),
//   	MultitrackInputConfiguration: &MultitrackInputConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		MaximumResolution: jsii.String("maximumResolution"),
//   		Policy: jsii.String("policy"),
//   	},
//   	Name: jsii.String("name"),
//   	Preset: jsii.String("preset"),
//   	RecordingConfigurationArn: jsii.String("recordingConfigurationArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html
//
type CfnChannelProps struct {
	// Whether the channel is authorized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-authorized
	//
	// Default: - false.
	//
	Authorized interface{} `field:"optional" json:"authorized" yaml:"authorized"`
	// Indicates which content-packaging format is used (MPEG-TS or fMP4).
	//
	// If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-containerformat
	//
	// Default: - "TS".
	//
	ContainerFormat *string `field:"optional" json:"containerFormat" yaml:"containerFormat"`
	// Whether the channel allows insecure ingest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-insecureingest
	//
	// Default: - false.
	//
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Channel latency mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-latencymode
	//
	// Default: - "LOW".
	//
	LatencyMode *string `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-multitrackinputconfiguration
	//
	MultitrackInputConfiguration interface{} `field:"optional" json:"multitrackInputConfiguration" yaml:"multitrackInputConfiguration"`
	// Channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-name
	//
	// Default: - "-".
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional transcode preset for the channel.
	//
	// This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string ("").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-preset
	//
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Recording Configuration ARN.
	//
	// A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-recordingconfigurationarn
	//
	// Default: - "".
	//
	RecordingConfigurationArn *string `field:"optional" json:"recordingConfigurationArn" yaml:"recordingConfigurationArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-type
	//
	// Default: - "STANDARD".
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

