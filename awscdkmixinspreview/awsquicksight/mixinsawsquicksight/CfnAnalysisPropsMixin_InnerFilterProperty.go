package mixinsawsquicksight


// The `InnerFilter` defines the subset of data to be used with the `NestedFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   innerFilterProperty := &InnerFilterProperty{
//   	CategoryInnerFilter: &CategoryInnerFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Configuration: &CategoryFilterConfigurationProperty{
//   			CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   				CategoryValue: jsii.String("categoryValue"),
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//   				ParameterName: jsii.String("parameterName"),
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   			CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   				CategoryValues: []*string{
//   					jsii.String("categoryValues"),
//   				},
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   			FilterListConfiguration: &FilterListConfigurationProperty{
//   				CategoryValues: []*string{
//   					jsii.String("categoryValues"),
//   				},
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   		},
//   		DefaultFilterControlConfiguration: &DefaultFilterControlConfigurationProperty{
//   			ControlOptions: &DefaultFilterControlOptionsProperty{
//   				DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
//   					CommitMode: jsii.String("commitMode"),
//   					DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   						DateIconVisibility: jsii.String("dateIconVisibility"),
//   						DateTimeFormat: jsii.String("dateTimeFormat"),
//   						HelperTextVisibility: jsii.String("helperTextVisibility"),
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				DefaultDropdownOptions: &DefaultFilterDropDownControlOptionsProperty{
//   					CommitMode: jsii.String("commitMode"),
//   					DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					SelectableValues: &FilterSelectableValuesProperty{
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				DefaultListOptions: &DefaultFilterListControlOptionsProperty{
//   					DisplayOptions: &ListControlDisplayOptionsProperty{
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						SearchOptions: &ListControlSearchOptionsProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					SelectableValues: &FilterSelectableValuesProperty{
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				DefaultRelativeDateTimeOptions: &DefaultRelativeDateTimeControlOptionsProperty{
//   					CommitMode: jsii.String("commitMode"),
//   					DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   						DateTimeFormat: jsii.String("dateTimeFormat"),
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   				DefaultSliderOptions: &DefaultSliderControlOptionsProperty{
//   					DisplayOptions: &SliderControlDisplayOptionsProperty{
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					MaximumValue: jsii.Number(123),
//   					MinimumValue: jsii.Number(123),
//   					StepSize: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				DefaultTextAreaOptions: &DefaultTextAreaControlOptionsProperty{
//   					Delimiter: jsii.String("delimiter"),
//   					DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   				DefaultTextFieldOptions: &DefaultTextFieldControlOptionsProperty{
//   					DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   						InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   							InfoIconText: jsii.String("infoIconText"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						TitleOptions: &LabelOptionsProperty{
//   							CustomLabel: jsii.String("customLabel"),
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   			},
//   			Title: jsii.String("title"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-innerfilter.html
//
type CfnAnalysisPropsMixin_InnerFilterProperty struct {
	// A `CategoryInnerFilter` filters text values for the `NestedFilter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-innerfilter.html#cfn-quicksight-analysis-innerfilter-categoryinnerfilter
	//
	CategoryInnerFilter interface{} `field:"optional" json:"categoryInnerFilter" yaml:"categoryInnerFilter"`
}

