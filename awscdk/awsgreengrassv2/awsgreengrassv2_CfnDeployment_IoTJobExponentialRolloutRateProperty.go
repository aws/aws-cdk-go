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
//   ioTJobExponentialRolloutRateProperty := &ioTJobExponentialRolloutRateProperty{
//   	baseRatePerMinute: jsii.Number(123),
//   	incrementFactor: jsii.Number(123),
//   	rateIncreaseCriteria: rateIncreaseCriteria,
//   }
//
type CfnDeployment_IoTJobExponentialRolloutRateProperty struct {
	// The minimum number of devices that receive a pending job notification, per minute, when the job starts.
	//
	// This parameter defines the initial rollout rate of the job.
	BaseRatePerMinute *float64 `field:"required" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// The exponential factor to increase the rollout rate for the job.
	//
	// This parameter supports up to one digit after the decimal (for example, you can specify `1.5` , but not `1.55` ).
	IncrementFactor *float64 `field:"required" json:"incrementFactor" yaml:"incrementFactor"`
	// The criteria to increase the rollout rate for the job.
	RateIncreaseCriteria interface{} `field:"required" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

