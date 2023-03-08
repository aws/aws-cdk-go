package awsmediaconvert


// Accelerated transcoding can significantly speed up jobs with long, visually complex content.
//
// Outputs that use this feature incur pro-tier pricing. For information about feature limitations, For more information, see [Job Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](https://docs.aws.amazon.com/mediaconvert/latest/ug/job-requirements.html) in the *AWS Elemental MediaConvert User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accelerationSettingsProperty := &accelerationSettingsProperty{
//   	mode: jsii.String("mode"),
//   }
//
type CfnJobTemplate_AccelerationSettingsProperty struct {
	// Specify the conditions when the service will run your job with accelerated transcoding.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

