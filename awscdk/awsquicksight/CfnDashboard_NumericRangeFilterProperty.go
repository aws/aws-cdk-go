package awsquicksight


// A `NumericRangeFilter` filters values that are within the value range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericRangeFilterProperty := &NumericRangeFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	AggregationFunction: &AggregationFunctionProperty{
//   		AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   			SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   			ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   		},
//   		CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   		DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   		NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   	},
//   	DefaultFilterControlConfiguration: &DefaultFilterControlConfigurationProperty{
//   		ControlOptions: &DefaultFilterControlOptionsProperty{
//   			DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
//   				CommitMode: jsii.String("commitMode"),
//   				DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			DefaultDropdownOptions: &DefaultFilterDropDownControlOptionsProperty{
//   				CommitMode: jsii.String("commitMode"),
//   				DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				SelectableValues: &FilterSelectableValuesProperty{
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			DefaultListOptions: &DefaultFilterListControlOptionsProperty{
//   				DisplayOptions: &ListControlDisplayOptionsProperty{
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					SearchOptions: &ListControlSearchOptionsProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				SelectableValues: &FilterSelectableValuesProperty{
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			DefaultRelativeDateTimeOptions: &DefaultRelativeDateTimeControlOptionsProperty{
//   				CommitMode: jsii.String("commitMode"),
//   				DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   			},
//   			DefaultSliderOptions: &DefaultSliderControlOptionsProperty{
//   				MaximumValue: jsii.Number(123),
//   				MinimumValue: jsii.Number(123),
//   				StepSize: jsii.Number(123),
//
//   				// the properties below are optional
//   				DisplayOptions: &SliderControlDisplayOptionsProperty{
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			DefaultTextAreaOptions: &DefaultTextAreaControlOptionsProperty{
//   				Delimiter: jsii.String("delimiter"),
//   				DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   			},
//   			DefaultTextFieldOptions: &DefaultTextFieldControlOptionsProperty{
//   				DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontSize: &FontSizeProperty{
//   								Absolute: jsii.String("absolute"),
//   								Relative: jsii.String("relative"),
//   							},
//   							FontStyle: jsii.String("fontStyle"),
//   							FontWeight: &FontWeightProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   			},
//   		},
//   		Title: jsii.String("title"),
//   	},
//   	IncludeMaximum: jsii.Boolean(false),
//   	IncludeMinimum: jsii.Boolean(false),
//   	RangeMaximum: &NumericRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		StaticValue: jsii.Number(123),
//   	},
//   	RangeMinimum: &NumericRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		StaticValue: jsii.Number(123),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html
//
type CfnDashboard_NumericRangeFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-nulloption
	//
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The aggregation function of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-aggregationfunction
	//
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// The default configurations for the associated controls.
	//
	// This applies only for filters that are scoped to multiple sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-defaultfiltercontrolconfiguration
	//
	DefaultFilterControlConfiguration interface{} `field:"optional" json:"defaultFilterControlConfiguration" yaml:"defaultFilterControlConfiguration"`
	// Determines whether the maximum value in the filter value range should be included in the filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-includemaximum
	//
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// Determines whether the minimum value in the filter value range should be included in the filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-includeminimum
	//
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// The maximum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-rangemaximum
	//
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-rangeminimum
	//
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-selectalloptions
	//
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

