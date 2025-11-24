package mixinsawsquicksight


// A `NestedFilter` filters data with a subset of data that is defined by the nested inner filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dateIconVisibility interface{}
//   var helperTextVisibility interface{}
//
//   nestedFilterProperty := &NestedFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	IncludeInnerSet: jsii.Boolean(false),
//   	InnerFilter: &InnerFilterProperty{
//   		CategoryInnerFilter: &CategoryInnerFilterProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			Configuration: &CategoryFilterConfigurationProperty{
//   				CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   					CategoryValue: jsii.String("categoryValue"),
//   					MatchOperator: jsii.String("matchOperator"),
//   					NullOption: jsii.String("nullOption"),
//   					ParameterName: jsii.String("parameterName"),
//   					SelectAllOptions: jsii.String("selectAllOptions"),
//   				},
//   				CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   					CategoryValues: []*string{
//   						jsii.String("categoryValues"),
//   					},
//   					MatchOperator: jsii.String("matchOperator"),
//   					NullOption: jsii.String("nullOption"),
//   					SelectAllOptions: jsii.String("selectAllOptions"),
//   				},
//   				FilterListConfiguration: &FilterListConfigurationProperty{
//   					CategoryValues: []*string{
//   						jsii.String("categoryValues"),
//   					},
//   					MatchOperator: jsii.String("matchOperator"),
//   					NullOption: jsii.String("nullOption"),
//   					SelectAllOptions: jsii.String("selectAllOptions"),
//   				},
//   			},
//   			DefaultFilterControlConfiguration: &DefaultFilterControlConfigurationProperty{
//   				ControlOptions: &DefaultFilterControlOptionsProperty{
//   					DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
//   						CommitMode: jsii.String("commitMode"),
//   						DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   							DateIconVisibility: dateIconVisibility,
//   							DateTimeFormat: jsii.String("dateTimeFormat"),
//   							HelperTextVisibility: helperTextVisibility,
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					DefaultDropdownOptions: &DefaultFilterDropDownControlOptionsProperty{
//   						CommitMode: jsii.String("commitMode"),
//   						DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						SelectableValues: &FilterSelectableValuesProperty{
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					DefaultListOptions: &DefaultFilterListControlOptionsProperty{
//   						DisplayOptions: &ListControlDisplayOptionsProperty{
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							SearchOptions: &ListControlSearchOptionsProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						SelectableValues: &FilterSelectableValuesProperty{
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					DefaultRelativeDateTimeOptions: &DefaultRelativeDateTimeControlOptionsProperty{
//   						CommitMode: jsii.String("commitMode"),
//   						DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   							DateTimeFormat: jsii.String("dateTimeFormat"),
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   					},
//   					DefaultSliderOptions: &DefaultSliderControlOptionsProperty{
//   						DisplayOptions: &SliderControlDisplayOptionsProperty{
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						MaximumValue: jsii.Number(123),
//   						MinimumValue: jsii.Number(123),
//   						StepSize: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   					DefaultTextAreaOptions: &DefaultTextAreaControlOptionsProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   					},
//   					DefaultTextFieldOptions: &DefaultTextFieldControlOptionsProperty{
//   						DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   							InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   								InfoIconText: jsii.String("infoIconText"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							TitleOptions: &LabelOptionsProperty{
//   								CustomLabel: jsii.String("customLabel"),
//   								FontConfiguration: &FontConfigurationProperty{
//   									FontColor: jsii.String("fontColor"),
//   									FontDecoration: jsii.String("fontDecoration"),
//   									FontFamily: jsii.String("fontFamily"),
//   									FontSize: &FontSizeProperty{
//   										Absolute: jsii.String("absolute"),
//   										Relative: jsii.String("relative"),
//   									},
//   									FontStyle: jsii.String("fontStyle"),
//   									FontWeight: &FontWeightProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   					},
//   				},
//   				Title: jsii.String("title"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html
//
type CfnTemplatePropsMixin_NestedFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-filterid
	//
	FilterId *string `field:"optional" json:"filterId" yaml:"filterId"`
	// A boolean condition to include or exclude the subset that is defined by the values of the nested inner filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-includeinnerset
	//
	// Default: - false.
	//
	IncludeInnerSet interface{} `field:"optional" json:"includeInnerSet" yaml:"includeInnerSet"`
	// The `InnerFilter` defines the subset of data to be used with the `NestedFilter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-innerfilter
	//
	InnerFilter interface{} `field:"optional" json:"innerFilter" yaml:"innerFilter"`
}

