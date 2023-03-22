package awsdatabrew


// Represents the sample size and sampling type for DataBrew to use for interactive data analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleProperty := &sampleProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	size: jsii.Number(123),
//   }
//
type CfnProject_SampleProperty struct {
	// The way in which DataBrew obtains rows from a dataset.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of rows in the sample.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

