package previewawsdatabrewmixins


// Represents the sample size and sampling type for DataBrew to use for interactive data analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sampleProperty := &SampleProperty{
//   	Size: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html
//
type CfnProjectPropsMixin_SampleProperty struct {
	// The number of rows in the sample.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The way in which DataBrew obtains rows from a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

