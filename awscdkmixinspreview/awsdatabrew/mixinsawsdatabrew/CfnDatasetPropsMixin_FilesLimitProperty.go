package mixinsawsdatabrew


// Represents a limit imposed on number of Amazon S3 files that should be selected for a dataset from a connected Amazon S3 path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filesLimitProperty := &FilesLimitProperty{
//   	MaxFiles: jsii.Number(123),
//   	Order: jsii.String("order"),
//   	OrderedBy: jsii.String("orderedBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-fileslimit.html
//
type CfnDatasetPropsMixin_FilesLimitProperty struct {
	// The number of Amazon S3 files to select.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-fileslimit.html#cfn-databrew-dataset-fileslimit-maxfiles
	//
	MaxFiles *float64 `field:"optional" json:"maxFiles" yaml:"maxFiles"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses DESCENDING order, i.e. most recent files are selected first. Anotherpossible value is ASCENDING.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-fileslimit.html#cfn-databrew-dataset-fileslimit-order
	//
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses LAST_MODIFIED_DATE as a sorting criteria. Currently it's the only allowed value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-fileslimit.html#cfn-databrew-dataset-fileslimit-orderedby
	//
	OrderedBy *string `field:"optional" json:"orderedBy" yaml:"orderedBy"`
}

