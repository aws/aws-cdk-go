package interfacesawsmedialive


// A reference to a ChannelPlacementGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelPlacementGroupReference := &ChannelPlacementGroupReference{
//   	ChannelPlacementGroupArn: jsii.String("channelPlacementGroupArn"),
//   	ChannelPlacementGroupId: jsii.String("channelPlacementGroupId"),
//   	ClusterId: jsii.String("clusterId"),
//   }
//
type ChannelPlacementGroupReference struct {
	// The ARN of the ChannelPlacementGroup resource.
	ChannelPlacementGroupArn *string `field:"required" json:"channelPlacementGroupArn" yaml:"channelPlacementGroupArn"`
	// The Id of the ChannelPlacementGroup resource.
	ChannelPlacementGroupId *string `field:"required" json:"channelPlacementGroupId" yaml:"channelPlacementGroupId"`
	// The ClusterId of the ChannelPlacementGroup resource.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
}

