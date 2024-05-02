package awsquicksight


// The default configuration for all dependent controls of the filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultFilterControlConfigurationProperty := &DefaultFilterControlConfigurationProperty{
//   	ControlOptions: &DefaultFilterControlOptionsProperty{
//   		DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
//   			DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		DefaultDropdownOptions: &DefaultFilterDropDownControlOptionsProperty{
//   			DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			SelectableValues: &FilterSelectableValuesProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		DefaultListOptions: &DefaultFilterListControlOptionsProperty{
//   			DisplayOptions: &ListControlDisplayOptionsProperty{
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				SearchOptions: &ListControlSearchOptionsProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			SelectableValues: &FilterSelectableValuesProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		DefaultRelativeDateTimeOptions: &DefaultRelativeDateTimeControlOptionsProperty{
//   			DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   		DefaultSliderOptions: &DefaultSliderControlOptionsProperty{
//   			MaximumValue: jsii.Number(123),
//   			MinimumValue: jsii.Number(123),
//   			StepSize: jsii.Number(123),
//
//   			// the properties below are optional
//   			DisplayOptions: &SliderControlDisplayOptionsProperty{
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		DefaultTextAreaOptions: &DefaultTextAreaControlOptionsProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   		DefaultTextFieldOptions: &DefaultTextFieldControlOptionsProperty{
//   			DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   				InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   					InfoIconText: jsii.String("infoIconText"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				TitleOptions: &LabelOptionsProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontSize: &FontSizeProperty{
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   	},
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html
//
type CfnAnalysis_DefaultFilterControlConfigurationProperty struct {
	// The control option for the `DefaultFilterControlConfiguration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html#cfn-quicksight-analysis-defaultfiltercontrolconfiguration-controloptions
	//
	ControlOptions interface{} `field:"required" json:"controlOptions" yaml:"controlOptions"`
	// The title of the `DefaultFilterControlConfiguration` .
	//
	// This title is shared by all controls that are tied to this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html#cfn-quicksight-analysis-defaultfiltercontrolconfiguration-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
}

