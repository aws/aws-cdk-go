package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   	InsecureIngest: jsii.Boolean(false),
//   	LatencyMode: jsii.String("latencyMode"),
//   	Name: jsii.String("name"),
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
type CfnChannelProps struct {
	// Whether the channel is authorized.
	//
	// *Default* : `false`.
	Authorized interface{} `field:"optional" json:"authorized" yaml:"authorized"`
	// Whether the channel allows insecure RTMP ingest.
	//
	// *Default* : `false`.
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Channel latency mode. Valid values:.
	//
	// - `NORMAL` : Use NORMAL to broadcast and deliver live video up to Full HD.
	// - `LOW` : Use LOW for near real-time interactions with viewers.
	//
	// > In the  console, `LOW` and `NORMAL` correspond to `Ultra-low` and `Standard` , respectively.
	//
	// *Default* : `LOW`.
	LatencyMode *string `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// Channel name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of a RecordingConfiguration resource.
	//
	// An empty string indicates that recording is disabled for the channel. A RecordingConfiguration ARN indicates that recording is enabled using the specified recording configuration. See the [RecordingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html) resource for more information and an example.
	//
	// *Default* : "" (empty string, recording is disabled).
	RecordingConfigurationArn *string `field:"optional" json:"recordingConfigurationArn" yaml:"recordingConfigurationArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// *If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.* Valid values:
	//
	// - `STANDARD` : Video is transcoded: multiple qualities are generated from the original input to automatically give viewers the best experience for their devices and network conditions. Transcoding allows higher playback quality across a range of download speeds. Resolution can be up to 1080p and bitrate can be up to 8.5 Mbps. Audio is transcoded only for renditions 360p and below; above that, audio is passed through.
	// - `BASIC` : Video is transmuxed: Amazon IVS delivers the original input to viewers. The viewer’s video-quality choice is limited to the original input. Resolution can be up to 1080p and bitrate can be up to 1.5 Mbps for 480p and up to 3.5 Mbps for resolutions between 480p and 1080p.
	//
	// *Default* : `STANDARD`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

