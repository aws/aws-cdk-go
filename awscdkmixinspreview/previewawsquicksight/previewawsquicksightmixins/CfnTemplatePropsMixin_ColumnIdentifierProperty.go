package previewawsquicksightmixins


// A column of a data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnIdentifierProperty := &ColumnIdentifierProperty{
//   	ColumnName: jsii.String("columnName"),
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnidentifier.html
//
type CfnTemplatePropsMixin_ColumnIdentifierProperty struct {
	// The name of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnidentifier.html#cfn-quicksight-template-columnidentifier-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The data set that the column belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnidentifier.html#cfn-quicksight-template-columnidentifier-datasetidentifier
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}

