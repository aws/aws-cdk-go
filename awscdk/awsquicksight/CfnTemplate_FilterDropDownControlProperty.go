package awsquicksight


// A control to display a dropdown list with buttons that are used to select a single value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDropDownControlProperty := &FilterDropDownControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
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
//   	SelectableValues: &FilterSelectableValuesProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html
//
type CfnTemplate_FilterDropDownControlProperty struct {
	// The ID of the `FilterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-filtercontrolid
	//
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-sourcefilterid
	//
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-cascadingcontrolconfiguration
	//
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// The visibility configuration of the Apply button on a `FilterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-commitmode
	//
	CommitMode *string `field:"optional" json:"commitMode" yaml:"commitMode"`
	// The display options of the `FilterDropDownControl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-displayoptions
	//
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// A list of selectable values that are used in a control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-selectablevalues
	//
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// The type of the `FilterDropDownControl` . Choose one of the following options:.
	//
	// - `MULTI_SELECT` : The user can select multiple entries from a dropdown menu.
	// - `SINGLE_SELECT` : The user can select a single entry from a dropdown menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

