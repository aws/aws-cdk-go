package awsquicksight


// A `TimeEqualityFilter` filters values that are equal to a given value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dateIconVisibility interface{}
//   var helperTextVisibility interface{}
//
//   timeEqualityFilterProperty := &TimeEqualityFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//
//   	// the properties below are optional
//   	DefaultFilterControlConfiguration: &DefaultFilterControlConfigurationProperty{
//   		ControlOptions: &DefaultFilterControlOptionsProperty{
//   			DefaultDateTimePickerOptions: &DefaultDateTimePickerControlOptionsProperty{
//   				CommitMode: jsii.String("commitMode"),
//   				DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   					DateIconVisibility: dateIconVisibility,
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					HelperTextVisibility: helperTextVisibility,
//   					InfoIconLabelOptions: &SheetControlInfoIconLabelOptionsProperty{
//   						InfoIconText: jsii.String("infoIconText"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					TitleOptions: &LabelOptionsProperty{
//   						CustomLabel: jsii.String("customLabel"),
//   						FontConfiguration: &FontConfigurationProperty{
//   							FontColor: jsii.String("fontColor"),
//   							FontDecoration: jsii.String("fontDecoration"),
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   							FontFamily: jsii.String("fontFamily"),
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
//   	ParameterName: jsii.String("parameterName"),
//   	RollingDate: &RollingDateConfigurationProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html
//
type CfnTemplate_TimeEqualityFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The default configurations for the associated controls.
	//
	// This applies only for filters that are scoped to multiple sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-defaultfiltercontrolconfiguration
	//
	DefaultFilterControlConfiguration interface{} `field:"optional" json:"defaultFilterControlConfiguration" yaml:"defaultFilterControlConfiguration"`
	// The parameter whose value should be used for the filter value.
	//
	// This field is mutually exclusive to `Value` and `RollingDate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The rolling date input for the `TimeEquality` filter.
	//
	// This field is mutually exclusive to `Value` and `ParameterName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-rollingdate
	//
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// The value of a `TimeEquality` filter.
	//
	// This field is mutually exclusive to `RollingDate` and `ParameterName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

