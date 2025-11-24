package mixinsawsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsquicksight/mixinsawsquicksight/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new Q topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTopicPropsMixin := awscdkmixinspreview.Mixins.NewCfnTopicPropsMixin(&CfnTopicMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html
//
type CfnTopicPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTopicMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTopicPropsMixin
type jsiiProxy_CfnTopicPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTopicPropsMixin) Props() *CfnTopicMixinProps {
	var returns *CfnTopicMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::Topic`.
func NewCfnTopicPropsMixin(props *CfnTopicMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTopicPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTopicPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnTopicPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::Topic`.
func NewCfnTopicPropsMixin_Override(c CfnTopicPropsMixin, props *CfnTopicMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnTopicPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTopicPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnTopicPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnTopicPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

