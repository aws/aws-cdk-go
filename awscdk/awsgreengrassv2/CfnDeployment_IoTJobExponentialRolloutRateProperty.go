package awsgreengrassv2


// Contains information about an exponential rollout rate for a configuration deployment job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rateIncreaseCriteria interface{}
//
//   ioTJobExponentialRolloutRateProperty := &IoTJobExponentialRolloutRateProperty{
//   	BaseRatePerMinute: jsii.Number(123),
//   	IncrementFactor: jsii.Number(123),
//   	RateIncreaseCriteria: rateIncreaseCriteria,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobexponentialrolloutrate.html
//
type CfnDeployment_IoTJobExponentialRolloutRateProperty struct {
	// The minimum number of devices that receive a pending job notification, per minute, when the job starts.
	//
	// This parameter defines the initial rollout rate of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobexponentialrolloutrate.html#cfn-greengrassv2-deployment-iotjobexponentialrolloutrate-baserateperminute
	//
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rollout rate for the job.
	//
	// This parameter supports up to one digit after the decimal (for example, you can specify `1.5` , but not `1.55` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobexponentialrolloutrate.html#cfn-greengrassv2-deployment-iotjobexponentialrolloutrate-incrementfactor
	//
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// The criteria to increase the rollout rate for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobexponentialrolloutrate.html#cfn-greengrassv2-deployment-iotjobexponentialrolloutrate-rateincreasecriteria
	//
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

