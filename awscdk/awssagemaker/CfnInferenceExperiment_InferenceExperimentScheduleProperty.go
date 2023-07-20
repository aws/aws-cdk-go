package awssagemaker


// The start and end times of an inference experiment.
//
// The maximum duration that you can set for an inference experiment is 30 days.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceExperimentScheduleProperty := &InferenceExperimentScheduleProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html
//
type CfnInferenceExperiment_InferenceExperimentScheduleProperty struct {
	// The timestamp at which the inference experiment ended or will end.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html#cfn-sagemaker-inferenceexperiment-inferenceexperimentschedule-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The timestamp at which the inference experiment started or will start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html#cfn-sagemaker-inferenceexperiment-inferenceexperimentschedule-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

