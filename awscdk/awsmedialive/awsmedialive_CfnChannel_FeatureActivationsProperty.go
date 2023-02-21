package awsmedialive


// Settings to enable specific features. You can't configure these features until you have enabled them in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureActivationsProperty := &featureActivationsProperty{
//   	inputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   }
//
type CfnChannel_FeatureActivationsProperty struct {
	// Enables the Input Prepare feature.
	//
	// You can create Input Prepare actions in the schedule only if this feature is enabled.
	// If you disable the feature on an existing schedule, make sure that you first delete all input prepare actions from the schedule.
	InputPrepareScheduleActions *string `field:"optional" json:"inputPrepareScheduleActions" yaml:"inputPrepareScheduleActions"`
}

