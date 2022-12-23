package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rateIncreaseCriteria interface{}
//
//   ioTJobExecutionsRolloutConfigProperty := &ioTJobExecutionsRolloutConfigProperty{
//   	exponentialRate: &ioTJobExponentialRolloutRateProperty{
//   		baseRatePerMinute: jsii.Number(123),
//   		incrementFactor: jsii.Number(123),
//   		rateIncreaseCriteria: rateIncreaseCriteria,
//   	},
//   	maximumPerMinute: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobExecutionsRolloutConfigProperty struct {
	// `CfnDeployment.IoTJobExecutionsRolloutConfigProperty.ExponentialRate`.
	ExponentialRate interface{} `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// `CfnDeployment.IoTJobExecutionsRolloutConfigProperty.MaximumPerMinute`.
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

