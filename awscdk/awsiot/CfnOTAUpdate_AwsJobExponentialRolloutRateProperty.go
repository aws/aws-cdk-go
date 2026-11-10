package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsJobExponentialRolloutRateProperty := &AwsJobExponentialRolloutRateProperty{
//   	BaseRatePerMinute: jsii.Number(123),
//   	IncrementFactor: jsii.Number(123),
//   	RateIncreaseCriteria: &AwsJobRateIncreaseCriteriaProperty{
//   		NumberOfNotifiedThings: jsii.Number(123),
//   		NumberOfSucceededThings: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexponentialrolloutrate.html
//
type CfnOTAUpdate_AwsJobExponentialRolloutRateProperty struct {
	// The minimum number of things that will be notified of a pending job, per minute, at the start of the job rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexponentialrolloutrate.html#cfn-iot-otaupdate-awsjobexponentialrolloutrate-baserateperminute
	//
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The rate of increase for a job rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexponentialrolloutrate.html#cfn-iot-otaupdate-awsjobexponentialrolloutrate-incrementfactor
	//
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexponentialrolloutrate.html#cfn-iot-otaupdate-awsjobexponentialrolloutrate-rateincreasecriteria
	//
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

