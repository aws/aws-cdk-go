package mixinsawsiot


// Allows you to create an exponential rate of rollout for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exponentialRolloutRateProperty := &ExponentialRolloutRateProperty{
//   	BaseRatePerMinute: jsii.Number(123),
//   	IncrementFactor: jsii.Number(123),
//   	RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   		NumberOfNotifiedThings: jsii.Number(123),
//   		NumberOfSucceededThings: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-exponentialrolloutrate.html
//
type CfnJobTemplatePropsMixin_ExponentialRolloutRateProperty struct {
	// The minimum number of things that will be notified of a pending job, per minute at the start of job rollout.
	//
	// This parameter allows you to define the initial rate of rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-exponentialrolloutrate.html#cfn-iot-jobtemplate-exponentialrolloutrate-baserateperminute
	//
	BaseRatePerMinute *float64 `field:"optional" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rate of rollout for a job.
	//
	// AWS IoT Core supports up to one digit after the decimal (for example, 1.5, but not 1.55).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-exponentialrolloutrate.html#cfn-iot-jobtemplate-exponentialrolloutrate-incrementfactor
	//
	IncrementFactor *float64 `field:"optional" json:"incrementFactor" yaml:"incrementFactor"`
	// The criteria to initiate the increase in rate of rollout for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-exponentialrolloutrate.html#cfn-iot-jobtemplate-exponentialrolloutrate-rateincreasecriteria
	//
	RateIncreaseCriteria interface{} `field:"optional" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

