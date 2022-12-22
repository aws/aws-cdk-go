package awsgreengrassv2


// Contains information about the rollout configuration for a job.
//
// This configuration defines the rate at which the job deploys a configuration to a fleet of target devices.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobExecutionsRolloutConfigProperty := &ioTJobExecutionsRolloutConfigProperty{
//   	exponentialRate: &ioTJobExponentialRolloutRateProperty{
//   		baseRatePerMinute: jsii.Number(123),
//   		incrementFactor: jsii.Number(123),
//   		rateIncreaseCriteria: &ioTJobRateIncreaseCriteriaProperty{
//   			numberOfNotifiedThings: jsii.Number(123),
//   			numberOfSucceededThings: jsii.Number(123),
//   		},
//   	},
//   	maximumPerMinute: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobExecutionsRolloutConfigProperty struct {
	// The exponential rate to increase the job rollout rate.
	ExponentialRate interface{} `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// The maximum number of devices that receive a pending job notification, per minute.
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

