package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterControlProperty := &FilterControlProperty{
//   	DateTimePicker: &FilterDateTimePickerControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   		Type: jsii.String("type"),
//   	},
//   	Dropdown: &FilterDropDownControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
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
//   		DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   			SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	List: &FilterListControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
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
//   					FontSize: &FontSizeProperty{
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
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	RelativeDateTime: &FilterRelativeDateTimeControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   	},
//   	Slider: &FilterSliderControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		MaximumValue: jsii.Number(123),
//   		MinimumValue: jsii.Number(123),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		StepSize: jsii.Number(123),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &SliderControlDisplayOptionsProperty{
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   		Type: jsii.String("type"),
//   	},
//   	TextArea: &FilterTextAreaControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Delimiter: jsii.String("delimiter"),
//   		DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   	},
//   	TextField: &FilterTextFieldControlProperty{
//   		FilterControlId: jsii.String("filterControlId"),
//   		SourceFilterId: jsii.String("sourceFilterId"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
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
//   	},
//   }
//
type CfnDashboard_FilterControlProperty struct {
	// `CfnDashboard.FilterControlProperty.DateTimePicker`.
	DateTimePicker interface{} `field:"optional" json:"dateTimePicker" yaml:"dateTimePicker"`
	// `CfnDashboard.FilterControlProperty.Dropdown`.
	Dropdown interface{} `field:"optional" json:"dropdown" yaml:"dropdown"`
	// `CfnDashboard.FilterControlProperty.List`.
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// `CfnDashboard.FilterControlProperty.RelativeDateTime`.
	RelativeDateTime interface{} `field:"optional" json:"relativeDateTime" yaml:"relativeDateTime"`
	// `CfnDashboard.FilterControlProperty.Slider`.
	Slider interface{} `field:"optional" json:"slider" yaml:"slider"`
	// `CfnDashboard.FilterControlProperty.TextArea`.
	TextArea interface{} `field:"optional" json:"textArea" yaml:"textArea"`
	// `CfnDashboard.FilterControlProperty.TextField`.
	TextField interface{} `field:"optional" json:"textField" yaml:"textField"`
}

