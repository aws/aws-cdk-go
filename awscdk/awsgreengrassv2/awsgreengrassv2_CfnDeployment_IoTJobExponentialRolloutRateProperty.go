package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rateIncreaseCriteria interface{}
//
//   ioTJobExponentialRolloutRateProperty := &ioTJobExponentialRolloutRateProperty{
//   	baseRatePerMinute: jsii.Number(123),
//   	incrementFactor: jsii.Number(123),
//   	rateIncreaseCriteria: rateIncreaseCriteria,
//   }
//
type CfnDeployment_IoTJobExponentialRolloutRateProperty struct {
	// `CfnDeployment.IoTJobExponentialRolloutRateProperty.BaseRatePerMinute`.
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// `CfnDeployment.IoTJobExponentialRolloutRateProperty.IncrementFactor`.
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// `CfnDeployment.IoTJobExponentialRolloutRateProperty.RateIncreaseCriteria`.
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

