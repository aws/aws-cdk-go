package awsmedialive


// Destination settings for a MediaPackage output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageOutputDestinationSettingsProperty := &mediaPackageOutputDestinationSettingsProperty{
//   	channelId: jsii.String("channelId"),
//   }
//
type CfnChannel_MediaPackageOutputDestinationSettingsProperty struct {
	// The ID of the channel in MediaPackage that is the destination for this output group.
	//
	// You don't need to specify the individual inputs in MediaPackage; MediaLive handles the connection of the two MediaLive pipelines to the two MediaPackage inputs. The MediaPackage channel and MediaLive channel must be in the same Region.
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
}

