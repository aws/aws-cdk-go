package awspinpoint


// A reference to a ADMChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aDMChannelReference := &ADMChannelReference{
//   	AdmChannelId: jsii.String("admChannelId"),
//   }
//
type ADMChannelReference struct {
	// The Id of the ADMChannel resource.
	AdmChannelId *string `field:"required" json:"admChannelId" yaml:"admChannelId"`
}

