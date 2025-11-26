package previewawsquicksightmixins


// A control to display a list of buttons or boxes.
//
// This is used to select either a single value or multiple values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterListControlProperty := &FilterListControlProperty{
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
//   	FilterControlId: jsii.String("filterControlId"),
//   	SelectableValues: &FilterSelectableValuesProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html
//
type CfnTemplatePropsMixin_FilterListControlProperty struct {
	// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-cascadingcontrolconfiguration
	//
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// The display options of a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The ID of the `FilterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-filtercontrolid
	//
	FilterControlId *string `field:"optional" json:"filterControlId" yaml:"filterControlId"`
	// A list of selectable values that are used in a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-selectablevalues
	//
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// The source filter ID of the `FilterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-sourcefilterid
	//
	SourceFilterId *string `field:"optional" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterListControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of the `FilterListControl` . Choose one of the following options:.
	//
	// - `MULTI_SELECT` : The user can select multiple entries from the list.
	// - `SINGLE_SELECT` : The user can select a single entry from the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterlistcontrol.html#cfn-quicksight-template-filterlistcontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

