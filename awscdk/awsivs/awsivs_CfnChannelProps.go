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
//   cfnChannelProps := &cfnChannelProps{
//   	authorized: jsii.Boolean(false),
//   	latencyMode: jsii.String("latencyMode"),
//   	name: jsii.String("name"),
//   	recordingConfigurationArn: jsii.String("recordingConfigurationArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnChannelProps struct {
	// Whether the channel is authorized.
	//
	// *Default* : `false`.
	Authorized interface{} `field:"optional" json:"authorized" yaml:"authorized"`
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
	// - `STANDARD` : Multiple qualities are generated from the original input, to automatically give viewers the best experience for their devices and network conditions. Resolution can be up to 1080p and bitrate can be up to 8.5 Mbps. Audio is transcoded only for renditions 360p and below; above that, audio is passed through.
	// - `BASIC` : delivers the original input to viewers. The viewer’s video-quality choice is limited to the original input. Resolution can be up to 480p and bitrate can be up to 1.5 Mbps.
	//
	// *Default* : `STANDARD`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

