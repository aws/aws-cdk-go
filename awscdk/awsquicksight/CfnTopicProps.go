package awsquicksight


// Properties for defining a `CfnTopic`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicProps := &CfnTopicProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSets: []interface{}{
//   		&DatasetMetadataProperty{
//   			DatasetArn: jsii.String("datasetArn"),
//
//   			// the properties below are optional
//   			CalculatedFields: []interface{}{
//   				&TopicCalculatedFieldProperty{
//   					CalculatedFieldName: jsii.String("calculatedFieldName"),
//   					Expression: jsii.String("expression"),
//
//   					// the properties below are optional
//   					Aggregation: jsii.String("aggregation"),
//   					AllowedAggregations: []*string{
//   						jsii.String("allowedAggregations"),
//   					},
//   					CalculatedFieldDescription: jsii.String("calculatedFieldDescription"),
//   					CalculatedFieldSynonyms: []*string{
//   						jsii.String("calculatedFieldSynonyms"),
//   					},
//   					CellValueSynonyms: []interface{}{
//   						&CellValueSynonymProperty{
//   							CellValue: jsii.String("cellValue"),
//   							Synonyms: []*string{
//   								jsii.String("synonyms"),
//   							},
//   						},
//   					},
//   					ColumnDataRole: jsii.String("columnDataRole"),
//   					ComparativeOrder: &ComparativeOrderProperty{
//   						SpecifedOrder: []*string{
//   							jsii.String("specifedOrder"),
//   						},
//   						TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   						UseOrdering: jsii.String("useOrdering"),
//   					},
//   					DefaultFormatting: &DefaultFormattingProperty{
//   						DisplayFormat: jsii.String("displayFormat"),
//   						DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   							BlankCellFormat: jsii.String("blankCellFormat"),
//   							CurrencySymbol: jsii.String("currencySymbol"),
//   							DateFormat: jsii.String("dateFormat"),
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							FractionDigits: jsii.Number(123),
//   							GroupingSeparator: jsii.String("groupingSeparator"),
//   							NegativeFormat: &NegativeFormatProperty{
//   								Prefix: jsii.String("prefix"),
//   								Suffix: jsii.String("suffix"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							Suffix: jsii.String("suffix"),
//   							UnitScaler: jsii.String("unitScaler"),
//   							UseBlankCellFormat: jsii.Boolean(false),
//   							UseGrouping: jsii.Boolean(false),
//   						},
//   					},
//   					IsIncludedInTopic: jsii.Boolean(false),
//   					NeverAggregateInFilter: jsii.Boolean(false),
//   					NonAdditive: jsii.Boolean(false),
//   					NotAllowedAggregations: []*string{
//   						jsii.String("notAllowedAggregations"),
//   					},
//   					SemanticType: &SemanticTypeProperty{
//   						FalseyCellValue: jsii.String("falseyCellValue"),
//   						FalseyCellValueSynonyms: []*string{
//   							jsii.String("falseyCellValueSynonyms"),
//   						},
//   						SubTypeName: jsii.String("subTypeName"),
//   						TruthyCellValue: jsii.String("truthyCellValue"),
//   						TruthyCellValueSynonyms: []*string{
//   							jsii.String("truthyCellValueSynonyms"),
//   						},
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   			Columns: []interface{}{
//   				&TopicColumnProperty{
//   					ColumnName: jsii.String("columnName"),
//
//   					// the properties below are optional
//   					Aggregation: jsii.String("aggregation"),
//   					AllowedAggregations: []*string{
//   						jsii.String("allowedAggregations"),
//   					},
//   					CellValueSynonyms: []interface{}{
//   						&CellValueSynonymProperty{
//   							CellValue: jsii.String("cellValue"),
//   							Synonyms: []*string{
//   								jsii.String("synonyms"),
//   							},
//   						},
//   					},
//   					ColumnDataRole: jsii.String("columnDataRole"),
//   					ColumnDescription: jsii.String("columnDescription"),
//   					ColumnFriendlyName: jsii.String("columnFriendlyName"),
//   					ColumnSynonyms: []*string{
//   						jsii.String("columnSynonyms"),
//   					},
//   					ComparativeOrder: &ComparativeOrderProperty{
//   						SpecifedOrder: []*string{
//   							jsii.String("specifedOrder"),
//   						},
//   						TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   						UseOrdering: jsii.String("useOrdering"),
//   					},
//   					DefaultFormatting: &DefaultFormattingProperty{
//   						DisplayFormat: jsii.String("displayFormat"),
//   						DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   							BlankCellFormat: jsii.String("blankCellFormat"),
//   							CurrencySymbol: jsii.String("currencySymbol"),
//   							DateFormat: jsii.String("dateFormat"),
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							FractionDigits: jsii.Number(123),
//   							GroupingSeparator: jsii.String("groupingSeparator"),
//   							NegativeFormat: &NegativeFormatProperty{
//   								Prefix: jsii.String("prefix"),
//   								Suffix: jsii.String("suffix"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							Suffix: jsii.String("suffix"),
//   							UnitScaler: jsii.String("unitScaler"),
//   							UseBlankCellFormat: jsii.Boolean(false),
//   							UseGrouping: jsii.Boolean(false),
//   						},
//   					},
//   					IsIncludedInTopic: jsii.Boolean(false),
//   					NeverAggregateInFilter: jsii.Boolean(false),
//   					NonAdditive: jsii.Boolean(false),
//   					NotAllowedAggregations: []*string{
//   						jsii.String("notAllowedAggregations"),
//   					},
//   					SemanticType: &SemanticTypeProperty{
//   						FalseyCellValue: jsii.String("falseyCellValue"),
//   						FalseyCellValueSynonyms: []*string{
//   							jsii.String("falseyCellValueSynonyms"),
//   						},
//   						SubTypeName: jsii.String("subTypeName"),
//   						TruthyCellValue: jsii.String("truthyCellValue"),
//   						TruthyCellValueSynonyms: []*string{
//   							jsii.String("truthyCellValueSynonyms"),
//   						},
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   			DataAggregation: &DataAggregationProperty{
//   				DatasetRowDateGranularity: jsii.String("datasetRowDateGranularity"),
//   				DefaultDateColumnName: jsii.String("defaultDateColumnName"),
//   			},
//   			DatasetDescription: jsii.String("datasetDescription"),
//   			DatasetName: jsii.String("datasetName"),
//   			Filters: []interface{}{
//   				&TopicFilterProperty{
//   					FilterName: jsii.String("filterName"),
//   					OperandFieldName: jsii.String("operandFieldName"),
//
//   					// the properties below are optional
//   					CategoryFilter: &TopicCategoryFilterProperty{
//   						CategoryFilterFunction: jsii.String("categoryFilterFunction"),
//   						CategoryFilterType: jsii.String("categoryFilterType"),
//   						Constant: &TopicCategoryFilterConstantProperty{
//   							CollectiveConstant: &CollectiveConstantProperty{
//   								ValueList: []*string{
//   									jsii.String("valueList"),
//   								},
//   							},
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   						Inverse: jsii.Boolean(false),
//   					},
//   					DateRangeFilter: &TopicDateRangeFilterProperty{
//   						Constant: &TopicRangeFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							RangeConstant: &RangeConstantProperty{
//   								Maximum: jsii.String("maximum"),
//   								Minimum: jsii.String("minimum"),
//   							},
//   						},
//   						Inclusive: jsii.Boolean(false),
//   					},
//   					FilterClass: jsii.String("filterClass"),
//   					FilterDescription: jsii.String("filterDescription"),
//   					FilterSynonyms: []*string{
//   						jsii.String("filterSynonyms"),
//   					},
//   					FilterType: jsii.String("filterType"),
//   					NumericEqualityFilter: &TopicNumericEqualityFilterProperty{
//   						Aggregation: jsii.String("aggregation"),
//   						Constant: &TopicSingularFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   					},
//   					NumericRangeFilter: &TopicNumericRangeFilterProperty{
//   						Aggregation: jsii.String("aggregation"),
//   						Constant: &TopicRangeFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							RangeConstant: &RangeConstantProperty{
//   								Maximum: jsii.String("maximum"),
//   								Minimum: jsii.String("minimum"),
//   							},
//   						},
//   						Inclusive: jsii.Boolean(false),
//   					},
//   					RelativeDateFilter: &TopicRelativeDateFilterProperty{
//   						Constant: &TopicSingularFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   						RelativeDateFilterFunction: jsii.String("relativeDateFilterFunction"),
//   						TimeGranularity: jsii.String("timeGranularity"),
//   					},
//   				},
//   			},
//   			NamedEntities: []interface{}{
//   				&TopicNamedEntityProperty{
//   					EntityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					Definition: []interface{}{
//   						&NamedEntityDefinitionProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Metric: &NamedEntityDefinitionMetricProperty{
//   								Aggregation: jsii.String("aggregation"),
//   								AggregationFunctionParameters: map[string]*string{
//   									"aggregationFunctionParametersKey": jsii.String("aggregationFunctionParameters"),
//   								},
//   							},
//   							PropertyName: jsii.String("propertyName"),
//   							PropertyRole: jsii.String("propertyRole"),
//   							PropertyUsage: jsii.String("propertyUsage"),
//   						},
//   					},
//   					EntityDescription: jsii.String("entityDescription"),
//   					EntitySynonyms: []*string{
//   						jsii.String("entitySynonyms"),
//   					},
//   					SemanticEntityType: &SemanticEntityTypeProperty{
//   						SubTypeName: jsii.String("subTypeName"),
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	TopicId: jsii.String("topicId"),
//   	UserExperienceVersion: jsii.String("userExperienceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html
//
type CfnTopicProps struct {
	// The ID of the AWS account that you want to create a topic in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The data sets that the topic is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-datasets
	//
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
	// The description of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID for the topic.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-topicid
	//
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
	// The user experience version of a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-userexperienceversion
	//
	UserExperienceVersion *string `field:"optional" json:"userExperienceVersion" yaml:"userExperienceVersion"`
}

