package previewawsquicksightmixins


// The control of a parameter that users can interact with in a dashboard or an analysis.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dateIconVisibility interface{}
//   var helperTextVisibility interface{}
//
//   parameterControlProperty := &ParameterControlProperty{
//   	DateTimePicker: &ParameterDateTimePickerControlProperty{
//   		DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   			DateIconVisibility: dateIconVisibility,
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			HelperTextVisibility: helperTextVisibility,
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
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//   	},
//   	Dropdown: &ParameterDropDownControlProperty{
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
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SelectableValues: &ParameterSelectableValuesProperty{
//   			LinkToDataSetColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	List: &ParameterListControlProperty{
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
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SelectableValues: &ParameterSelectableValuesProperty{
//   			LinkToDataSetColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//   		Type: jsii.String("type"),
//   	},
//   	Slider: &ParameterSliderControlProperty{
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
//   		MaximumValue: jsii.Number(123),
//   		MinimumValue: jsii.Number(123),
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		StepSize: jsii.Number(123),
//   		Title: jsii.String("title"),
//   	},
//   	TextArea: &ParameterTextAreaControlProperty{
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
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//   	},
//   	TextField: &ParameterTextFieldControlProperty{
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
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html
//
type CfnTemplatePropsMixin_ParameterControlProperty struct {
	// A control from a date parameter that specifies date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-datetimepicker
	//
	DateTimePicker interface{} `field:"optional" json:"dateTimePicker" yaml:"dateTimePicker"`
	// A control to display a dropdown list with buttons that are used to select a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-dropdown
	//
	Dropdown interface{} `field:"optional" json:"dropdown" yaml:"dropdown"`
	// A control to display a list with buttons or boxes that are used to select either a single value or multiple values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// A control to display a horizontal toggle bar.
	//
	// This is used to change a value by sliding the toggle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-slider
	//
	Slider interface{} `field:"optional" json:"slider" yaml:"slider"`
	// A control to display a text box that is used to enter multiple entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-textarea
	//
	TextArea interface{} `field:"optional" json:"textArea" yaml:"textArea"`
	// A control to display a text box that is used to enter a single entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-textfield
	//
	TextField interface{} `field:"optional" json:"textField" yaml:"textField"`
}

