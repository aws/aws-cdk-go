package awsiot


// Allows you to create an exponential rate of rollout for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exponentialRolloutRateProperty := &ExponentialRolloutRateProperty{
//   	BaseRatePerMinute: jsii.Number(123),
//   	IncrementFactor: jsii.Number(123),
//   	RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   		NumberOfNotifiedThings: jsii.Number(123),
//   		NumberOfSucceededThings: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-exponentialrolloutrate.html
//
type CfnJob_ExponentialRolloutRateProperty struct {
	// The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-exponentialrolloutrate.html#cfn-iot-job-exponentialrolloutrate-baserateperminute
	//
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rate of rollout for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-exponentialrolloutrate.html#cfn-iot-job-exponentialrolloutrate-incrementfactor
	//
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// Allows you to define a criteria to initiate the increase in rate of rollout for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-exponentialrolloutrate.html#cfn-iot-job-exponentialrolloutrate-rateincreasecriteria
	//
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

