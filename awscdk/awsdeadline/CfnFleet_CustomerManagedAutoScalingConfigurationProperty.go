package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedAutoScalingConfigurationProperty := &CustomerManagedAutoScalingConfigurationProperty{
//   	ScaleOutWorkersPerMinute: jsii.Number(123),
//   	StandbyWorkerCount: jsii.Number(123),
//   	WorkerIdleDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedautoscalingconfiguration.html
//
type CfnFleet_CustomerManagedAutoScalingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedautoscalingconfiguration.html#cfn-deadline-fleet-customermanagedautoscalingconfiguration-scaleoutworkersperminute
	//
	ScaleOutWorkersPerMinute *float64 `field:"optional" json:"scaleOutWorkersPerMinute" yaml:"scaleOutWorkersPerMinute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedautoscalingconfiguration.html#cfn-deadline-fleet-customermanagedautoscalingconfiguration-standbyworkercount
	//
	StandbyWorkerCount *float64 `field:"optional" json:"standbyWorkerCount" yaml:"standbyWorkerCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedautoscalingconfiguration.html#cfn-deadline-fleet-customermanagedautoscalingconfiguration-workeridledurationseconds
	//
	WorkerIdleDurationSeconds *float64 `field:"optional" json:"workerIdleDurationSeconds" yaml:"workerIdleDurationSeconds"`
}

