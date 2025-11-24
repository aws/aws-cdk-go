package mixinsawsquicksight


// The option that corresponds to the control type of the filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultFilterControlOptionsProperty := &DefaultFilterControlOptionsProperty{
//   	DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
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
//   		Type: jsii.String("type"),
//   	},
//   	DefaultDropdownOptions: &DefaultFilterDropDownControlOptionsProperty{
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
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	DefaultListOptions: &DefaultFilterListControlOptionsProperty{
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
//   		SelectableValues: &FilterSelectableValuesProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	DefaultRelativeDateTimeOptions: &DefaultRelativeDateTimeControlOptionsProperty{
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
//   	},
//   	DefaultSliderOptions: &DefaultSliderControlOptionsProperty{
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
//   		StepSize: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	DefaultTextAreaOptions: &DefaultTextAreaControlOptionsProperty{
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
//   	},
//   	DefaultTextFieldOptions: &DefaultTextFieldControlOptionsProperty{
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html
//
type CfnAnalysisPropsMixin_DefaultFilterControlOptionsProperty struct {
	// The default options that correspond to the filter control type of a `DateTimePicker` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaultdatetimepickeroptions
	//
	DefaultDateTimePickerOptions interface{} `field:"optional" json:"defaultDateTimePickerOptions" yaml:"defaultDateTimePickerOptions"`
	// The default options that correspond to the `Dropdown` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaultdropdownoptions
	//
	DefaultDropdownOptions interface{} `field:"optional" json:"defaultDropdownOptions" yaml:"defaultDropdownOptions"`
	// The default options that correspond to the `List` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaultlistoptions
	//
	DefaultListOptions interface{} `field:"optional" json:"defaultListOptions" yaml:"defaultListOptions"`
	// The default options that correspond to the `RelativeDateTime` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaultrelativedatetimeoptions
	//
	DefaultRelativeDateTimeOptions interface{} `field:"optional" json:"defaultRelativeDateTimeOptions" yaml:"defaultRelativeDateTimeOptions"`
	// The default options that correspond to the `Slider` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaultslideroptions
	//
	DefaultSliderOptions interface{} `field:"optional" json:"defaultSliderOptions" yaml:"defaultSliderOptions"`
	// The default options that correspond to the `TextArea` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaulttextareaoptions
	//
	DefaultTextAreaOptions interface{} `field:"optional" json:"defaultTextAreaOptions" yaml:"defaultTextAreaOptions"`
	// The default options that correspond to the `TextField` filter control type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontroloptions.html#cfn-quicksight-analysis-defaultfiltercontroloptions-defaulttextfieldoptions
	//
	DefaultTextFieldOptions interface{} `field:"optional" json:"defaultTextFieldOptions" yaml:"defaultTextFieldOptions"`
}

