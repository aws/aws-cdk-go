package awsmedialivealpha


// Feature activations for the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var featureActivationState FeatureActivationState
//
//   featureActivations := &FeatureActivations{
//   	InputPrepareScheduleActions: featureActivationState,
//   	OutputStaticImageOverlayScheduleActions: featureActivationState,
//   }
//
// Experimental.
type FeatureActivations struct {
	// Enable Input Prepare schedule actions.
	// Default: - DISABLED, applied by MediaLive.
	//
	// Experimental.
	InputPrepareScheduleActions FeatureActivationState `field:"optional" json:"inputPrepareScheduleActions" yaml:"inputPrepareScheduleActions"`
	// Enable output static image overlay schedule actions.
	// Default: - DISABLED, applied by MediaLive.
	//
	// Experimental.
	OutputStaticImageOverlayScheduleActions FeatureActivationState `field:"optional" json:"outputStaticImageOverlayScheduleActions" yaml:"outputStaticImageOverlayScheduleActions"`
}

