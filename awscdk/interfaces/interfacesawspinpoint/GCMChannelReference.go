package interfacesawspinpoint


// A reference to a GCMChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gCMChannelReference := &GCMChannelReference{
//   	GcmChannelId: jsii.String("gcmChannelId"),
//   }
//
type GCMChannelReference struct {
	// The Id of the GCMChannel resource.
	GcmChannelId *string `field:"required" json:"gcmChannelId" yaml:"gcmChannelId"`
}

