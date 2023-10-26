package awsquicksight


// A structure that represents a calculated field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicCalculatedFieldProperty := &TopicCalculatedFieldProperty{
//   	CalculatedFieldName: jsii.String("calculatedFieldName"),
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	Aggregation: jsii.String("aggregation"),
//   	AllowedAggregations: []*string{
//   		jsii.String("allowedAggregations"),
//   	},
//   	CalculatedFieldDescription: jsii.String("calculatedFieldDescription"),
//   	CalculatedFieldSynonyms: []*string{
//   		jsii.String("calculatedFieldSynonyms"),
//   	},
//   	CellValueSynonyms: []interface{}{
//   		&CellValueSynonymProperty{
//   			CellValue: jsii.String("cellValue"),
//   			Synonyms: []*string{
//   				jsii.String("synonyms"),
//   			},
//   		},
//   	},
//   	ColumnDataRole: jsii.String("columnDataRole"),
//   	ComparativeOrder: &ComparativeOrderProperty{
//   		SpecifedOrder: []*string{
//   			jsii.String("specifedOrder"),
//   		},
//   		TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   		UseOrdering: jsii.String("useOrdering"),
//   	},
//   	DefaultFormatting: &DefaultFormattingProperty{
//   		DisplayFormat: jsii.String("displayFormat"),
//   		DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   			BlankCellFormat: jsii.String("blankCellFormat"),
//   			CurrencySymbol: jsii.String("currencySymbol"),
//   			DateFormat: jsii.String("dateFormat"),
//   			DecimalSeparator: jsii.String("decimalSeparator"),
//   			FractionDigits: jsii.Number(123),
//   			GroupingSeparator: jsii.String("groupingSeparator"),
//   			NegativeFormat: &NegativeFormatProperty{
//   				Prefix: jsii.String("prefix"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   			Prefix: jsii.String("prefix"),
//   			Suffix: jsii.String("suffix"),
//   			UnitScaler: jsii.String("unitScaler"),
//   			UseBlankCellFormat: jsii.Boolean(false),
//   			UseGrouping: jsii.Boolean(false),
//   		},
//   	},
//   	IsIncludedInTopic: jsii.Boolean(false),
//   	NeverAggregateInFilter: jsii.Boolean(false),
//   	NonAdditive: jsii.Boolean(false),
//   	NotAllowedAggregations: []*string{
//   		jsii.String("notAllowedAggregations"),
//   	},
//   	SemanticType: &SemanticTypeProperty{
//   		FalseyCellValue: jsii.String("falseyCellValue"),
//   		FalseyCellValueSynonyms: []*string{
//   			jsii.String("falseyCellValueSynonyms"),
//   		},
//   		SubTypeName: jsii.String("subTypeName"),
//   		TruthyCellValue: jsii.String("truthyCellValue"),
//   		TruthyCellValueSynonyms: []*string{
//   			jsii.String("truthyCellValueSynonyms"),
//   		},
//   		TypeName: jsii.String("typeName"),
//   		TypeParameters: map[string]*string{
//   			"typeParametersKey": jsii.String("typeParameters"),
//   		},
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html
//
type CfnTopic_TopicCalculatedFieldProperty struct {
	// The calculated field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfieldname
	//
	CalculatedFieldName *string `field:"required" json:"calculatedFieldName" yaml:"calculatedFieldName"`
	// The calculated field expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The default aggregation.
	//
	// Valid values for this structure are `SUM` , `MAX` , `MIN` , `COUNT` , `DISTINCT_COUNT` , and `AVERAGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-aggregation
	//
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The list of aggregation types that are allowed for the calculated field.
	//
	// Valid values for this structure are `COUNT` , `DISTINCT_COUNT` , `MIN` , `MAX` , `MEDIAN` , `SUM` , `AVERAGE` , `STDEV` , `STDEVP` , `VAR` , `VARP` , and `PERCENTILE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-allowedaggregations
	//
	AllowedAggregations *[]*string `field:"optional" json:"allowedAggregations" yaml:"allowedAggregations"`
	// The calculated field description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfielddescription
	//
	CalculatedFieldDescription *string `field:"optional" json:"calculatedFieldDescription" yaml:"calculatedFieldDescription"`
	// The other names or aliases for the calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfieldsynonyms
	//
	CalculatedFieldSynonyms *[]*string `field:"optional" json:"calculatedFieldSynonyms" yaml:"calculatedFieldSynonyms"`
	// The other names or aliases for the calculated field cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-cellvaluesynonyms
	//
	CellValueSynonyms interface{} `field:"optional" json:"cellValueSynonyms" yaml:"cellValueSynonyms"`
	// The column data role for a calculated field.
	//
	// Valid values for this structure are `DIMENSION` and `MEASURE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-columndatarole
	//
	ColumnDataRole *string `field:"optional" json:"columnDataRole" yaml:"columnDataRole"`
	// The order in which data is displayed for the calculated field when it's used in a comparative context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-comparativeorder
	//
	ComparativeOrder interface{} `field:"optional" json:"comparativeOrder" yaml:"comparativeOrder"`
	// The default formatting definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-defaultformatting
	//
	DefaultFormatting interface{} `field:"optional" json:"defaultFormatting" yaml:"defaultFormatting"`
	// A boolean value that indicates if a calculated field is included in the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-isincludedintopic
	//
	// Default: - false.
	//
	IsIncludedInTopic interface{} `field:"optional" json:"isIncludedInTopic" yaml:"isIncludedInTopic"`
	// A Boolean value that indicates whether to never aggregate calculated field in filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-neveraggregateinfilter
	//
	// Default: - false.
	//
	NeverAggregateInFilter interface{} `field:"optional" json:"neverAggregateInFilter" yaml:"neverAggregateInFilter"`
	// The non additive for the table style target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-nonadditive
	//
	// Default: - false.
	//
	NonAdditive interface{} `field:"optional" json:"nonAdditive" yaml:"nonAdditive"`
	// The list of aggregation types that are not allowed for the calculated field.
	//
	// Valid values for this structure are `COUNT` , `DISTINCT_COUNT` , `MIN` , `MAX` , `MEDIAN` , `SUM` , `AVERAGE` , `STDEV` , `STDEVP` , `VAR` , `VARP` , and `PERCENTILE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-notallowedaggregations
	//
	NotAllowedAggregations *[]*string `field:"optional" json:"notAllowedAggregations" yaml:"notAllowedAggregations"`
	// The semantic type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-semantictype
	//
	SemanticType interface{} `field:"optional" json:"semanticType" yaml:"semanticType"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

