package previewawsapptestmixins


// Defines a data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetProperty := &DataSetProperty{
//   	Ccsid: jsii.String("ccsid"),
//   	Format: jsii.String("format"),
//   	Length: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html
//
type CfnTestCasePropsMixin_DataSetProperty struct {
	// The CCSID of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html#cfn-apptest-testcase-dataset-ccsid
	//
	Ccsid *string `field:"optional" json:"ccsid" yaml:"ccsid"`
	// The format of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html#cfn-apptest-testcase-dataset-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The length of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html#cfn-apptest-testcase-dataset-length
	//
	Length *float64 `field:"optional" json:"length" yaml:"length"`
	// The name of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html#cfn-apptest-testcase-dataset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-dataset.html#cfn-apptest-testcase-dataset-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

