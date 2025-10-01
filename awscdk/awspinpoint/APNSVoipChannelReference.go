package awspinpoint


// A reference to a APNSVoipChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPNSVoipChannelReference := &APNSVoipChannelReference{
//   	ApnsVoipChannelId: jsii.String("apnsVoipChannelId"),
//   }
//
type APNSVoipChannelReference struct {
	// The Id of the APNSVoipChannel resource.
	ApnsVoipChannelId *string `field:"required" json:"apnsVoipChannelId" yaml:"apnsVoipChannelId"`
}

