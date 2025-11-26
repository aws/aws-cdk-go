package previewawsquicksightmixins


// A control to display a list with buttons or boxes that are used to select either a single value or multiple values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterListControlProperty := &ParameterListControlProperty{
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
//   	DisplayOptions: &ListControlDisplayOptionsProperty{
//   		InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   			InfoIconText: jsii.String("infoIconText"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		SearchOptions: &ListControlSearchOptionsProperty{
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
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
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
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SelectableValues: &ParameterSelectableValuesProperty{
//   		LinkToDataSetColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html
//
type CfnDashboardPropsMixin_ParameterListControlProperty struct {
	// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-cascadingcontrolconfiguration
	//
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `ParameterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-parametercontrolid
	//
	ParameterControlId *string `field:"optional" json:"parameterControlId" yaml:"parameterControlId"`
	// A list of selectable values that are used in a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-selectablevalues
	//
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// The source parameter name of the `ParameterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-sourceparametername
	//
	SourceParameterName *string `field:"optional" json:"sourceParameterName" yaml:"sourceParameterName"`
	// The title of the `ParameterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of `ParameterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterlistcontrol.html#cfn-quicksight-dashboard-parameterlistcontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

