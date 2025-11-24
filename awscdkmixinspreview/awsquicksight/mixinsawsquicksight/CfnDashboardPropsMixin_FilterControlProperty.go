package mixinsawsquicksight


// The control of a filter that is used to interact with a dashboard or an analysis.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterControlProperty := &FilterControlProperty{
//   	CrossSheet: &FilterCrossSheetControlProperty{
//   		CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   			SourceControls: []interface{}{
//   				&CascadingControlSourceProperty{
//   					ColumnToMatch: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   				},
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   	},
//   	DateTimePicker: &FilterDateTimePickerControlProperty{
//   		CommitMode: jsii.String("commitMode"),
//   		DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   			DateIconVisibility: jsii.String("dateIconVisibility"),
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			HelperTextVisibility: jsii.String("helperTextVisibility"),
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	Dropdown: &FilterDropDownControlProperty{
//   		CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   			SourceControls: []interface{}{
//   				&CascadingControlSourceProperty{
//   					ColumnToMatch: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   				},
//   			},
//   		},
//   		CommitMode: jsii.String("commitMode"),
//   		DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	List: &FilterListControlProperty{
//   		CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   			SourceControls: []interface{}{
//   				&CascadingControlSourceProperty{
//   					ColumnToMatch: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   				},
//   			},
//   		},
//   		DisplayOptions: &ListControlDisplayOptionsProperty{
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			SearchOptions: &ListControlSearchOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	RelativeDateTime: &FilterRelativeDateTimeControlProperty{
//   		CommitMode: jsii.String("commitMode"),
//   		DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   	},
//   	Slider: &FilterSliderControlProperty{
//   		DisplayOptions: &SliderControlDisplayOptionsProperty{
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		MaximumValue: jsii.Number(123),
//   		MinimumValue: jsii.Number(123),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		StepSize: jsii.Number(123),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	TextArea: &FilterTextAreaControlProperty{
//   		Delimiter: jsii.String("delimiter"),
//   		DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   	},
//   	TextField: &FilterTextFieldControlProperty{
//   		DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   			InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   				InfoIconText: jsii.String("infoIconText"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html
//
type CfnDashboardPropsMixin_FilterControlProperty struct {
	// A control from a filter that is scoped across more than one sheet.
	//
	// This represents your filter control on a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-crosssheet
	//
	CrossSheet interface{} `field:"optional" json:"crossSheet" yaml:"crossSheet"`
	// A control from a date filter that is used to specify date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-datetimepicker
	//
	DateTimePicker interface{} `field:"optional" json:"dateTimePicker" yaml:"dateTimePicker"`
	// A control to display a dropdown list with buttons that are used to select a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-dropdown
	//
	Dropdown interface{} `field:"optional" json:"dropdown" yaml:"dropdown"`
	// A control to display a list of buttons or boxes.
	//
	// This is used to select either a single value or multiple values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// A control from a date filter that is used to specify the relative date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-relativedatetime
	//
	RelativeDateTime interface{} `field:"optional" json:"relativeDateTime" yaml:"relativeDateTime"`
	// A control to display a horizontal toggle bar.
	//
	// This is used to change a value by sliding the toggle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-slider
	//
	Slider interface{} `field:"optional" json:"slider" yaml:"slider"`
	// A control to display a text box that is used to enter multiple entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-textarea
	//
	TextArea interface{} `field:"optional" json:"textArea" yaml:"textArea"`
	// A control to display a text box that is used to enter a single entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtercontrol.html#cfn-quicksight-dashboard-filtercontrol-textfield
	//
	TextField interface{} `field:"optional" json:"textField" yaml:"textField"`
}

