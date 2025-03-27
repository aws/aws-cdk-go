package awsdatabrew


// Represents the sample size and sampling type for DataBrew to use for interactive data analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleProperty := &SampleProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Size: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html
//
type CfnProject_SampleProperty struct {
	// The way in which DataBrew obtains rows from a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of rows in the sample.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

