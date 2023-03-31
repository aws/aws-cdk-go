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
//   inferenceExperimentScheduleProperty := &inferenceExperimentScheduleProperty{
//   	endTime: jsii.String("endTime"),
//   	startTime: jsii.String("startTime"),
//   }
//
type CfnInferenceExperiment_InferenceExperimentScheduleProperty struct {
	// The timestamp at which the inference experiment ended or will end.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The timestamp at which the inference experiment started or will start.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

