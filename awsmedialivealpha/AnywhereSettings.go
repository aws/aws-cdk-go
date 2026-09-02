package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Anywhere settings for running the channel on AWS Elemental Anywhere.
//
// Example:
//   var stack Stack
//   var cluster ICluster
//   var input IInput
//   var video EncodeConfiguration
//   var bucket IBucket
//
//
//   cpg := medialive.NewChannelPlacementGroup(stack, jsii.String("CPG"), &ChannelPlacementGroupProps{
//   	ChannelPlacementGroupName: jsii.String("my-cpg"),
//   	Cluster: Cluster,
//   })
//
//   medialive.NewChannel(stack, jsii.String("AnywhereChannel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	AnywhereSettings: &AnywhereSettings{
//   		Cluster: *Cluster,
//   		ChannelPlacementGroup: cpg,
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type AnywhereSettings struct {
	// The cluster for this channel.
	// Experimental.
	Cluster interfacesawsmedialive.IClusterRef `field:"required" json:"cluster" yaml:"cluster"`
	// The channel placement group for this channel.
	// Default: - no placement group.
	//
	// Experimental.
	ChannelPlacementGroup interfacesawsmedialive.IChannelPlacementGroupRef `field:"optional" json:"channelPlacementGroup" yaml:"channelPlacementGroup"`
}

