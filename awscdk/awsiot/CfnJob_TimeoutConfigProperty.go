package awsiot


// Specifies the amount of time each device has to finish its execution of the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutConfigProperty := &TimeoutConfigProperty{
//   	InProgressTimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-timeoutconfig.html
//
type CfnJob_TimeoutConfigProperty struct {
	// Specifies the amount of time, in minutes, this device has to finish execution of this job.
	//
	// The timeout interval can be anywhere between 1 minute and 7 days (1 to 10080 minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-timeoutconfig.html#cfn-iot-job-timeoutconfig-inprogresstimeoutinminutes
	//
	InProgressTimeoutInMinutes *float64 `field:"optional" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

