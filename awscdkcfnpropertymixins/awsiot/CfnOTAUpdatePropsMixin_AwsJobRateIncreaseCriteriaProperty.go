package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   awsJobRateIncreaseCriteriaProperty := &AwsJobRateIncreaseCriteriaProperty{
//   	NumberOfNotifiedThings: jsii.Number(123),
//   	NumberOfSucceededThings: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobrateincreasecriteria.html
//
type CfnOTAUpdatePropsMixin_AwsJobRateIncreaseCriteriaProperty struct {
	// When this number of things have been notified, it will initiate an increase in the rollout rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobrateincreasecriteria.html#cfn-iot-otaupdate-awsjobrateincreasecriteria-numberofnotifiedthings
	//
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// When this number of things have succeeded in their job execution, it will initiate an increase in the rollout rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobrateincreasecriteria.html#cfn-iot-otaupdate-awsjobrateincreasecriteria-numberofsucceededthings
	//
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

