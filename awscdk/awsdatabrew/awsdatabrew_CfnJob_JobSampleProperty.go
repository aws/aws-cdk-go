package awsdatabrew


// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
//
// If a `JobSample` value isn't provided, the default is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobSampleProperty := &jobSampleProperty{
//   	mode: jsii.String("mode"),
//   	size: jsii.Number(123),
//   }
//
type CfnJob_JobSampleProperty struct {
	// A value that determines whether the profile job is run on the entire dataset or a specified number of rows.
	//
	// This value must be one of the following:
	//
	// - FULL_DATASET - The profile job is run on the entire dataset.
	// - CUSTOM_ROWS - The profile job is run on the number of rows specified in the `Size` parameter.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The `Size` parameter is only required when the mode is CUSTOM_ROWS.
	//
	// The profile job is run on the specified number of rows. The maximum value for size is Long.MAX_VALUE.
	//
	// Long.MAX_VALUE = 9223372036854775807
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

