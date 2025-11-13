package interfacesawspinpoint


// A reference to a APNSVoipSandboxChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPNSVoipSandboxChannelReference := &APNSVoipSandboxChannelReference{
//   	ApnsVoipSandboxChannelId: jsii.String("apnsVoipSandboxChannelId"),
//   }
//
type APNSVoipSandboxChannelReference struct {
	// The Id of the APNSVoipSandboxChannel resource.
	ApnsVoipSandboxChannelId *string `field:"required" json:"apnsVoipSandboxChannelId" yaml:"apnsVoipSandboxChannelId"`
}

