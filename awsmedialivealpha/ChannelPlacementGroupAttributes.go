package awsmedialivealpha


// Attributes for importing an existing Channel Placement Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   channelPlacementGroupAttributes := &ChannelPlacementGroupAttributes{
//   	ChannelPlacementGroupArn: jsii.String("channelPlacementGroupArn"),
//   	ChannelPlacementGroupId: jsii.String("channelPlacementGroupId"),
//   	ClusterId: jsii.String("clusterId"),
//   }
//
// Experimental.
type ChannelPlacementGroupAttributes struct {
	// The ARN of the channel placement group.
	// Experimental.
	ChannelPlacementGroupArn *string `field:"required" json:"channelPlacementGroupArn" yaml:"channelPlacementGroupArn"`
	// The ID of the channel placement group.
	// Experimental.
	ChannelPlacementGroupId *string `field:"required" json:"channelPlacementGroupId" yaml:"channelPlacementGroupId"`
	// The ID of the cluster this group belongs to.
	// Experimental.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
}

