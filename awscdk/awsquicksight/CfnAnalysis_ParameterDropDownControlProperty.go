package awsquicksight


// A control to display a dropdown list with buttons that are used to select a single value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterDropDownControlProperty := &ParameterDropDownControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
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
//   	CommitMode: jsii.String("commitMode"),
//   	DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   		InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   			InfoIconText: jsii.String("infoIconText"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	SelectableValues: &ParameterSelectableValuesProperty{
//   		LinkToDataSetColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html
//
type CfnAnalysis_ParameterDropDownControlProperty struct {
	// The ID of the `ParameterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-parametercontrolid
	//
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// The source parameter name of the `ParameterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-sourceparametername
	//
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-cascadingcontrolconfiguration
	//
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// The visibility configuration of the Apply button on a `ParameterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// A list of selectable values that are used in a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-selectablevalues
	//
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// The type parameter name of the `ParameterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

