package awsquicksight


// Represents a column in a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicColumnProperty := &TopicColumnProperty{
//   	ColumnName: jsii.String("columnName"),
//
//   	// the properties below are optional
//   	Aggregation: jsii.String("aggregation"),
//   	AllowedAggregations: []*string{
//   		jsii.String("allowedAggregations"),
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
//   	ColumnDescription: jsii.String("columnDescription"),
//   	ColumnFriendlyName: jsii.String("columnFriendlyName"),
//   	ColumnSynonyms: []*string{
//   		jsii.String("columnSynonyms"),
//   	},
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
type CfnTopic_TopicColumnProperty struct {
	// The name of the column.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The type of aggregation that is performed on the column data when it's queried.
	//
	// Valid values for this structure are `SUM` , `MAX` , `MIN` , `COUNT` , `DISTINCT_COUNT` , and `AVERAGE` .
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The list of aggregation types that are allowed for the column.
	//
	// Valid values for this structure are `COUNT` , `DISTINCT_COUNT` , `MIN` , `MAX` , `MEDIAN` , `SUM` , `AVERAGE` , `STDEV` , `STDEVP` , `VAR` , `VARP` , and `PERCENTILE` .
	AllowedAggregations *[]*string `field:"optional" json:"allowedAggregations" yaml:"allowedAggregations"`
	// The other names or aliases for the column cell value.
	CellValueSynonyms interface{} `field:"optional" json:"cellValueSynonyms" yaml:"cellValueSynonyms"`
	// The role of the column in the data.
	//
	// Valid values are `DIMENSION` and `MEASURE` .
	ColumnDataRole *string `field:"optional" json:"columnDataRole" yaml:"columnDataRole"`
	// A description of the column and its contents.
	ColumnDescription *string `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// A user-friendly name for the column.
	ColumnFriendlyName *string `field:"optional" json:"columnFriendlyName" yaml:"columnFriendlyName"`
	// The other names or aliases for the column.
	ColumnSynonyms *[]*string `field:"optional" json:"columnSynonyms" yaml:"columnSynonyms"`
	// The order in which data is displayed for the column when it's used in a comparative context.
	ComparativeOrder interface{} `field:"optional" json:"comparativeOrder" yaml:"comparativeOrder"`
	// The default formatting used for values in the column.
	DefaultFormatting interface{} `field:"optional" json:"defaultFormatting" yaml:"defaultFormatting"`
	// A Boolean value that indicates whether the column is included in the query results.
	IsIncludedInTopic interface{} `field:"optional" json:"isIncludedInTopic" yaml:"isIncludedInTopic"`
	// A Boolean value that indicates whether to aggregate the column data when it's used in a filter context.
	NeverAggregateInFilter interface{} `field:"optional" json:"neverAggregateInFilter" yaml:"neverAggregateInFilter"`
	// The list of aggregation types that are not allowed for the column.
	//
	// Valid values for this structure are `COUNT` , `DISTINCT_COUNT` , `MIN` , `MAX` , `MEDIAN` , `SUM` , `AVERAGE` , `STDEV` , `STDEVP` , `VAR` , `VARP` , and `PERCENTILE` .
	NotAllowedAggregations *[]*string `field:"optional" json:"notAllowedAggregations" yaml:"notAllowedAggregations"`
	// The semantic type of data contained in the column.
	SemanticType interface{} `field:"optional" json:"semanticType" yaml:"semanticType"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

