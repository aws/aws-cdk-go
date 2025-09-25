package awsevidently


// Use this structure to start and stop the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runningStatusObjectProperty := &RunningStatusObjectProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	AnalysisCompleteTime: jsii.String("analysisCompleteTime"),
//   	DesiredState: jsii.String("desiredState"),
//   	Reason: jsii.String("reason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html
//
type CfnExperiment_RunningStatusObjectProperty struct {
	// To start the experiment now, specify `START` for this parameter.
	//
	// If this experiment is currently running and you want to stop it now, specify `STOP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// If you are using AWS CloudFormation to start the experiment, use this field to specify when the experiment is to end.
	//
	// The format is as a UNIX timestamp. For more information about this format, see [The Current Epoch Unix Timestamp](https://docs.aws.amazon.com/https://www.unixtimestamp.com/index.php) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-analysiscompletetime
	//
	AnalysisCompleteTime *string `field:"optional" json:"analysisCompleteTime" yaml:"analysisCompleteTime"`
	// If you are using AWS CloudFormation to stop this experiment, specify either `COMPLETED` or `CANCELLED` here to indicate how to classify this experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-desiredstate
	//
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// If you are using AWS CloudFormation to stop this experiment, this is an optional field that you can use to record why the experiment is being stopped or cancelled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-runningstatusobject.html#cfn-evidently-experiment-runningstatusobject-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

