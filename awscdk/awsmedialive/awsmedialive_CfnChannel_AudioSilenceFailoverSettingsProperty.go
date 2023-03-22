package awsmedialive


// MediaLive will perform a failover if audio is not detected in this input for the specified period.
//
// The parent of this entity is FailoverConditionSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSilenceFailoverSettingsProperty := &audioSilenceFailoverSettingsProperty{
//   	audioSelectorName: jsii.String("audioSelectorName"),
//   	audioSilenceThresholdMsec: jsii.Number(123),
//   }
//
type CfnChannel_AudioSilenceFailoverSettingsProperty struct {
	// The name of the audio selector in the input that MediaLive should monitor to detect silence.
	//
	// Select your most important rendition. If you didn't create an audio selector in this input, leave blank.
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// The amount of time (in milliseconds) that the active input must be silent before automatic input failover occurs.
	//
	// Silence is defined as audio loss or audio quieter than -50 dBFS.
	AudioSilenceThresholdMsec *float64 `field:"optional" json:"audioSilenceThresholdMsec" yaml:"audioSilenceThresholdMsec"`
}

