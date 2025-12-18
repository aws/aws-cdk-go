package previewawsautoscalingmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionTriggersProperty := &RetentionTriggersProperty{
//   	TerminateHookAbandon: jsii.String("terminateHookAbandon"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-retentiontriggers.html
//
type CfnAutoScalingGroupPropsMixin_RetentionTriggersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-retentiontriggers.html#cfn-autoscaling-autoscalinggroup-retentiontriggers-terminatehookabandon
	//
	TerminateHookAbandon *string `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}

