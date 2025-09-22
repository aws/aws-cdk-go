package awspinpoint


// A reference to a APNSSandboxChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPNSSandboxChannelReference := &APNSSandboxChannelReference{
//   	ApnsSandboxChannelId: jsii.String("apnsSandboxChannelId"),
//   }
//
type APNSSandboxChannelReference struct {
	// The Id of the APNSSandboxChannel resource.
	ApnsSandboxChannelId *string `field:"required" json:"apnsSandboxChannelId" yaml:"apnsSandboxChannelId"`
}

