package awsiot


// Allows you to define a criteria to initiate the increase in rate of rollout for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rateIncreaseCriteriaProperty := &rateIncreaseCriteriaProperty{
//   	numberOfNotifiedThings: jsii.Number(123),
//   	numberOfSucceededThings: jsii.Number(123),
//   }
//
type CfnJobTemplate_RateIncreaseCriteriaProperty struct {
	// The threshold for number of notified things that will initiate the increase in rate of rollout.
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// The threshold for number of succeeded things that will initiate the increase in rate of rollout.
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

