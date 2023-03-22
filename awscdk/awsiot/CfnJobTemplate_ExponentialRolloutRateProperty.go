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
type CfnJobTemplate_ExponentialRolloutRateProperty struct {
	// The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.
	//
	// This parameter allows you to define the initial rate of rollout.
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rate of rollout for a job.
	//
	// AWS IoT Core supports up to one digit after the decimal (for example, 1.5, but not 1.55).
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// The criteria to initiate the increase in rate of rollout for a job.
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

