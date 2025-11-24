package mixinsawsdatabrew


// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
//
// If a `JobSample` value isn't provided, the default is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobSampleProperty := &JobSampleProperty{
//   	Mode: jsii.String("mode"),
//   	Size: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html
//
type CfnJobPropsMixin_JobSampleProperty struct {
	// A value that determines whether the profile job is run on the entire dataset or a specified number of rows.
	//
	// This value must be one of the following:
	//
	// - FULL_DATASET - The profile job is run on the entire dataset.
	// - CUSTOM_ROWS - The profile job is run on the number of rows specified in the `Size` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html#cfn-databrew-job-jobsample-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The `Size` parameter is only required when the mode is CUSTOM_ROWS.
	//
	// The profile job is run on the specified number of rows. The maximum value for size is Long.MAX_VALUE.
	//
	// Long.MAX_VALUE = 9223372036854775807
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html#cfn-databrew-job-jobsample-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

