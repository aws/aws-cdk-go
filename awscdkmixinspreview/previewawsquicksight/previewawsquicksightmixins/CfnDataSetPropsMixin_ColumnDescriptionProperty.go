package previewawsquicksightmixins


// Metadata that contains a description for a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnDescriptionProperty := &ColumnDescriptionProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columndescription.html
//
type CfnDataSetPropsMixin_ColumnDescriptionProperty struct {
	// The text of a description for a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columndescription.html#cfn-quicksight-dataset-columndescription-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

