package awsiot


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
	// `CfnJobTemplate.RateIncreaseCriteriaProperty.NumberOfNotifiedThings`.
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// `CfnJobTemplate.RateIncreaseCriteriaProperty.NumberOfSucceededThings`.
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

