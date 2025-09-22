package awspinpoint


// A reference to a EmailChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailChannelReference := &EmailChannelReference{
//   	EmailChannelId: jsii.String("emailChannelId"),
//   }
//
type EmailChannelReference struct {
	// The Id of the EmailChannel resource.
	EmailChannelId *string `field:"required" json:"emailChannelId" yaml:"emailChannelId"`
}

