package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsJobExecutionsRolloutConfigProperty := &AwsJobExecutionsRolloutConfigProperty{
//   	ExponentialRate: &AwsJobExponentialRolloutRateProperty{
//   		BaseRatePerMinute: jsii.Number(123),
//   		IncrementFactor: jsii.Number(123),
//   		RateIncreaseCriteria: &AwsJobRateIncreaseCriteriaProperty{
//   			NumberOfNotifiedThings: jsii.Number(123),
//   			NumberOfSucceededThings: jsii.Number(123),
//   		},
//   	},
//   	MaximumPerMinute: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexecutionsrolloutconfig.html
//
type CfnOTAUpdate_AwsJobExecutionsRolloutConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexecutionsrolloutconfig.html#cfn-iot-otaupdate-awsjobexecutionsrolloutconfig-exponentialrate
	//
	ExponentialRate interface{} `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// The maximum number of OTA update job executions started per minute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobexecutionsrolloutconfig.html#cfn-iot-otaupdate-awsjobexecutionsrolloutconfig-maximumperminute
	//
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

