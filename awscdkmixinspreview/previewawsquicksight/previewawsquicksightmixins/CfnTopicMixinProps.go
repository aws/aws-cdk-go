package previewawsquicksightmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTopicPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTopicMixinProps := &CfnTopicMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ConfigOptions: &TopicConfigOptionsProperty{
//   		QBusinessInsightsEnabled: jsii.Boolean(false),
//   	},
//   	CustomInstructions: &CustomInstructionsProperty{
//   		CustomInstructionsString: jsii.String("customInstructionsString"),
//   	},
//   	DataSets: []interface{}{
//   		&DatasetMetadataProperty{
//   			CalculatedFields: []interface{}{
//   				&TopicCalculatedFieldProperty{
//   					Aggregation: jsii.String("aggregation"),
//   					AllowedAggregations: []*string{
//   						jsii.String("allowedAggregations"),
//   					},
//   					CalculatedFieldDescription: jsii.String("calculatedFieldDescription"),
//   					CalculatedFieldName: jsii.String("calculatedFieldName"),
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
//   					DisableIndexing: jsii.Boolean(false),
//   					Expression: jsii.String("expression"),
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
//   					ColumnName: jsii.String("columnName"),
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
//   					DisableIndexing: jsii.Boolean(false),
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
//   			DatasetArn: jsii.String("datasetArn"),
//   			DatasetDescription: jsii.String("datasetDescription"),
//   			DatasetName: jsii.String("datasetName"),
//   			Filters: []interface{}{
//   				&TopicFilterProperty{
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
//   					FilterName: jsii.String("filterName"),
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
//   					OperandFieldName: jsii.String("operandFieldName"),
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
//   					EntityName: jsii.String("entityName"),
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
//   	FolderArns: []*string{
//   		jsii.String("folderArns"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicId: jsii.String("topicId"),
//   	UserExperienceVersion: jsii.String("userExperienceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html
//
type CfnTopicMixinProps struct {
	// The ID of the AWS account that you want to create a topic in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Configuration options for a `Topic` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-configoptions
	//
	ConfigOptions interface{} `field:"optional" json:"configOptions" yaml:"configOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-custominstructions
	//
	CustomInstructions interface{} `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// The data sets that the topic is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-datasets
	//
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
	// The description of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// The name of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID for the topic.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-topicid
	//
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
	// The user experience version of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html#cfn-quicksight-topic-userexperienceversion
	//
	UserExperienceVersion *string `field:"optional" json:"userExperienceVersion" yaml:"userExperienceVersion"`
}

