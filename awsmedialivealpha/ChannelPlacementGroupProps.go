package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Properties for creating a MediaLive Channel Placement Group.
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
type ChannelPlacementGroupProps struct {
	// The cluster this channel placement group belongs to.
	// Experimental.
	Cluster interfacesawsmedialive.IClusterRef `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the channel placement group.
	// Default: - auto-generated name.
	//
	// Experimental.
	ChannelPlacementGroupName *string `field:"optional" json:"channelPlacementGroupName" yaml:"channelPlacementGroupName"`
	// List of node IDs for the channel placement group.
	// Default: - no nodes.
	//
	// Experimental.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// Tags to add to the channel placement group.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

