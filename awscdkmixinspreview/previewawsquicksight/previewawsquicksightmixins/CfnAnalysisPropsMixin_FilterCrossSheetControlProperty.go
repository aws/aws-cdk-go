package previewawsquicksightmixins


// A control from a filter that is scoped across more than one sheet.
//
// This represents your filter control on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterCrossSheetControlProperty := &FilterCrossSheetControlProperty{
//   	CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   		SourceControls: []interface{}{
//   			&CascadingControlSourceProperty{
//   				ColumnToMatch: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   			},
//   		},
//   	},
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercrosssheetcontrol.html
//
type CfnAnalysisPropsMixin_FilterCrossSheetControlProperty struct {
	// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercrosssheetcontrol.html#cfn-quicksight-analysis-filtercrosssheetcontrol-cascadingcontrolconfiguration
	//
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// The ID of the `FilterCrossSheetControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercrosssheetcontrol.html#cfn-quicksight-analysis-filtercrosssheetcontrol-filtercontrolid
	//
	FilterControlId *string `field:"optional" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterCrossSheetControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercrosssheetcontrol.html#cfn-quicksight-analysis-filtercrosssheetcontrol-sourcefilterid
	//
	SourceFilterId *string `field:"optional" json:"sourceFilterId" yaml:"sourceFilterId"`
}

