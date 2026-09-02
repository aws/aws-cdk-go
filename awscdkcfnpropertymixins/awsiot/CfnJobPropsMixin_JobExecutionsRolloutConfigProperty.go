package awsiot


// Allows you to create a staged rollout of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   jobExecutionsRolloutConfigProperty := &JobExecutionsRolloutConfigProperty{
//   	ExponentialRate: &ExponentialRolloutRateProperty{
//   		BaseRatePerMinute: jsii.Number(123),
//   		IncrementFactor: jsii.Number(123),
//   		RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   			NumberOfNotifiedThings: jsii.Number(123),
//   			NumberOfSucceededThings: jsii.Number(123),
//   		},
//   	},
//   	MaximumPerMinute: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-jobexecutionsrolloutconfig.html
//
type CfnJobPropsMixin_JobExecutionsRolloutConfigProperty struct {
	// Allows you to create an exponential rate of rollout for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-jobexecutionsrolloutconfig.html#cfn-iot-job-jobexecutionsrolloutconfig-exponentialrate
	//
	ExponentialRate interface{} `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-jobexecutionsrolloutconfig.html#cfn-iot-job-jobexecutionsrolloutconfig-maximumperminute
	//
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

