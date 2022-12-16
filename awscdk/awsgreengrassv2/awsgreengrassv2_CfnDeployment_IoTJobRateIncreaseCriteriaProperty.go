package awsgreengrassv2


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
	// `CfnDeployment.IoTJobRateIncreaseCriteriaProperty.NumberOfNotifiedThings`.
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// `CfnDeployment.IoTJobRateIncreaseCriteriaProperty.NumberOfSucceededThings`.
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

