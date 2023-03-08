package awsiot


// Allows you to create a staged rollout of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobExecutionsRolloutConfigProperty := &JobExecutionsRolloutConfigProperty{
//   	ExponentialRolloutRate: &ExponentialRolloutRateProperty{
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
type CfnJobTemplate_JobExecutionsRolloutConfigProperty struct {
	// Allows you to create an exponential rate of rollout for a job.
	ExponentialRolloutRate interface{} `field:"optional" json:"exponentialRolloutRate" yaml:"exponentialRolloutRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

