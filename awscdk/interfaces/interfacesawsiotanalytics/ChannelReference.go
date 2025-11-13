package interfacesawsiotanalytics


// A reference to a Channel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelReference := &ChannelReference{
//   	ChannelName: jsii.String("channelName"),
//   }
//
type ChannelReference struct {
	// The ChannelName of the Channel resource.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
}

