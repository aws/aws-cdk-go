package awsgreengrassv2


// Contains information about criteria to meet before a job increases its rollout rate.
//
// Specify either `numberOfNotifiedThings` or `numberOfSucceededThings` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobRateIncreaseCriteriaProperty := &ioTJobRateIncreaseCriteriaProperty{
//   	numberOfNotifiedThings: jsii.Number(123),
//   	numberOfSucceededThings: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobRateIncreaseCriteriaProperty struct {
	// The number of devices to receive the job notification before the rollout rate increases.
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// The number of devices to successfully run the configuration job before the rollout rate increases.
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

