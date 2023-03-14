package awsiot


// Allows you to create a staged rollout of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobExecutionsRolloutConfigProperty := &jobExecutionsRolloutConfigProperty{
//   	exponentialRolloutRate: &exponentialRolloutRateProperty{
//   		baseRatePerMinute: jsii.Number(123),
//   		incrementFactor: jsii.Number(123),
//   		rateIncreaseCriteria: &rateIncreaseCriteriaProperty{
//   			numberOfNotifiedThings: jsii.Number(123),
//   			numberOfSucceededThings: jsii.Number(123),
//   		},
//   	},
//   	maximumPerMinute: jsii.Number(123),
//   }
//
type CfnJobTemplate_JobExecutionsRolloutConfigProperty struct {
	// `CfnJobTemplate.JobExecutionsRolloutConfigProperty.ExponentialRolloutRate`.
	ExponentialRolloutRate interface{} `field:"optional" json:"exponentialRolloutRate" yaml:"exponentialRolloutRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

