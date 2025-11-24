package mixinsawsmedialive


// Settings to enable specific features. You can't configure these features until you have enabled them in the channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   featureActivationsProperty := &FeatureActivationsProperty{
//   	InputPrepareScheduleActions: jsii.String("inputPrepareScheduleActions"),
//   	OutputStaticImageOverlayScheduleActions: jsii.String("outputStaticImageOverlayScheduleActions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-featureactivations.html
//
type CfnChannelPropsMixin_FeatureActivationsProperty struct {
	// Enables the Input Prepare feature.
	//
	// You can create Input Prepare actions in the schedule only if this feature is enabled.
	// If you disable the feature on an existing schedule, make sure that you first delete all input prepare actions from the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-featureactivations.html#cfn-medialive-channel-featureactivations-inputpreparescheduleactions
	//
	InputPrepareScheduleActions *string `field:"optional" json:"inputPrepareScheduleActions" yaml:"inputPrepareScheduleActions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-featureactivations.html#cfn-medialive-channel-featureactivations-outputstaticimageoverlayscheduleactions
	//
	OutputStaticImageOverlayScheduleActions *string `field:"optional" json:"outputStaticImageOverlayScheduleActions" yaml:"outputStaticImageOverlayScheduleActions"`
}

