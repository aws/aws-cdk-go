package awspinpoint


// A reference to a VoiceChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   voiceChannelReference := &VoiceChannelReference{
//   	VoiceChannelId: jsii.String("voiceChannelId"),
//   }
//
type VoiceChannelReference struct {
	// The Id of the VoiceChannel resource.
	VoiceChannelId *string `field:"required" json:"voiceChannelId" yaml:"voiceChannelId"`
}

