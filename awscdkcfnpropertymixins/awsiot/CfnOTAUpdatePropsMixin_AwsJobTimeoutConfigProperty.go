package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   awsJobTimeoutConfigProperty := &AwsJobTimeoutConfigProperty{
//   	InProgressTimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobtimeoutconfig.html
//
type CfnOTAUpdatePropsMixin_AwsJobTimeoutConfigProperty struct {
	// Specifies the amount of time, in minutes, this device has to finish execution of this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobtimeoutconfig.html#cfn-iot-otaupdate-awsjobtimeoutconfig-inprogresstimeoutinminutes
	//
	InProgressTimeoutInMinutes *float64 `field:"optional" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

