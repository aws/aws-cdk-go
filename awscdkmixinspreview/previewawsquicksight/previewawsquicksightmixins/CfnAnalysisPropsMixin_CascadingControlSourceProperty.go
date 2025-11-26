package previewawsquicksightmixins


// The source controls that are used in a `CascadingControlConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cascadingControlSourceProperty := &CascadingControlSourceProperty{
//   	ColumnToMatch: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-cascadingcontrolsource.html
//
type CfnAnalysisPropsMixin_CascadingControlSourceProperty struct {
	// The column identifier that determines which column to look up for the source sheet control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-cascadingcontrolsource.html#cfn-quicksight-analysis-cascadingcontrolsource-columntomatch
	//
	ColumnToMatch interface{} `field:"optional" json:"columnToMatch" yaml:"columnToMatch"`
	// The source sheet control ID of a `CascadingControlSource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-cascadingcontrolsource.html#cfn-quicksight-analysis-cascadingcontrolsource-sourcesheetcontrolid
	//
	SourceSheetControlId *string `field:"optional" json:"sourceSheetControlId" yaml:"sourceSheetControlId"`
}

