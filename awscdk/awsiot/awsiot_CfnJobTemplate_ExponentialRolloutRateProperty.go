package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exponentialRolloutRateProperty := &exponentialRolloutRateProperty{
//   	baseRatePerMinute: jsii.Number(123),
//   	incrementFactor: jsii.Number(123),
//   	rateIncreaseCriteria: &rateIncreaseCriteriaProperty{
//   		numberOfNotifiedThings: jsii.Number(123),
//   		numberOfSucceededThings: jsii.Number(123),
//   	},
//   }
//
type CfnJobTemplate_ExponentialRolloutRateProperty struct {
	// `CfnJobTemplate.ExponentialRolloutRateProperty.BaseRatePerMinute`.
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// `CfnJobTemplate.ExponentialRolloutRateProperty.IncrementFactor`.
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// `CfnJobTemplate.ExponentialRolloutRateProperty.RateIncreaseCriteria`.
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

