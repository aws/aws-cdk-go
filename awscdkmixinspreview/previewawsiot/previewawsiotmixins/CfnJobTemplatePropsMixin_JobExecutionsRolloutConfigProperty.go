package previewawsiotmixins


// Allows you to create a staged rollout of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobExecutionsRolloutConfigProperty := &JobExecutionsRolloutConfigProperty{
//   	ExponentialRolloutRate: &ExponentialRolloutRateProperty{
//   		BaseRatePerMinute: jsii.Number(123),
//   		IncrementFactor: jsii.Number(123),
//   		RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   			NumberOfNotifiedThings: jsii.Number(123),
//   			NumberOfSucceededThings: jsii.Number(123),
//   		},
//   	},
//   	MaximumPerMinute: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html
//
type CfnJobTemplatePropsMixin_JobExecutionsRolloutConfigProperty struct {
	// The rate of increase for a job rollout.
	//
	// This parameter allows you to define an exponential rate for a job rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html#cfn-iot-jobtemplate-jobexecutionsrolloutconfig-exponentialrolloutrate
	//
	ExponentialRolloutRate interface{} `field:"optional" json:"exponentialRolloutRate" yaml:"exponentialRolloutRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html#cfn-iot-jobtemplate-jobexecutionsrolloutconfig-maximumperminute
	//
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

