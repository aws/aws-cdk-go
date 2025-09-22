package awspinpoint


// A reference to a APNSChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPNSChannelReference := &APNSChannelReference{
//   	ApnsChannelId: jsii.String("apnsChannelId"),
//   }
//
type APNSChannelReference struct {
	// The Id of the APNSChannel resource.
	ApnsChannelId *string `field:"required" json:"apnsChannelId" yaml:"apnsChannelId"`
}

