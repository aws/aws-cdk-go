package awslex


// Identifies the Amazon Polly voice used for audio interaction with the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   voiceSettingsProperty := &voiceSettingsProperty{
//   	voiceId: jsii.String("voiceId"),
//   }
//
type CfnBot_VoiceSettingsProperty struct {
	// The Amazon Polly voice used for voice interaction with the user.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
}

