package awsquicksight


// A structure that represents a filter used to select items for a topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicFilterProperty := &TopicFilterProperty{
//   	FilterName: jsii.String("filterName"),
//   	OperandFieldName: jsii.String("operandFieldName"),
//
//   	// the properties below are optional
//   	CategoryFilter: &TopicCategoryFilterProperty{
//   		CategoryFilterFunction: jsii.String("categoryFilterFunction"),
//   		CategoryFilterType: jsii.String("categoryFilterType"),
//   		Constant: &TopicCategoryFilterConstantProperty{
//   			CollectiveConstant: &CollectiveConstantProperty{
//   				ValueList: []*string{
//   					jsii.String("valueList"),
//   				},
//   			},
//   			ConstantType: jsii.String("constantType"),
//   			SingularConstant: jsii.String("singularConstant"),
//   		},
//   		Inverse: jsii.Boolean(false),
//   	},
//   	DateRangeFilter: &TopicDateRangeFilterProperty{
//   		Constant: &TopicRangeFilterConstantProperty{
//   			ConstantType: jsii.String("constantType"),
//   			RangeConstant: &RangeConstantProperty{
//   				Maximum: jsii.String("maximum"),
//   				Minimum: jsii.String("minimum"),
//   			},
//   		},
//   		Inclusive: jsii.Boolean(false),
//   	},
//   	FilterClass: jsii.String("filterClass"),
//   	FilterDescription: jsii.String("filterDescription"),
//   	FilterSynonyms: []*string{
//   		jsii.String("filterSynonyms"),
//   	},
//   	FilterType: jsii.String("filterType"),
//   	NumericEqualityFilter: &TopicNumericEqualityFilterProperty{
//   		Aggregation: jsii.String("aggregation"),
//   		Constant: &TopicSingularFilterConstantProperty{
//   			ConstantType: jsii.String("constantType"),
//   			SingularConstant: jsii.String("singularConstant"),
//   		},
//   	},
//   	NumericRangeFilter: &TopicNumericRangeFilterProperty{
//   		Aggregation: jsii.String("aggregation"),
//   		Constant: &TopicRangeFilterConstantProperty{
//   			ConstantType: jsii.String("constantType"),
//   			RangeConstant: &RangeConstantProperty{
//   				Maximum: jsii.String("maximum"),
//   				Minimum: jsii.String("minimum"),
//   			},
//   		},
//   		Inclusive: jsii.Boolean(false),
//   	},
//   	RelativeDateFilter: &TopicRelativeDateFilterProperty{
//   		Constant: &TopicSingularFilterConstantProperty{
//   			ConstantType: jsii.String("constantType"),
//   			SingularConstant: jsii.String("singularConstant"),
//   		},
//   		RelativeDateFilterFunction: jsii.String("relativeDateFilterFunction"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   }
//
type CfnTopic_TopicFilterProperty struct {
	// The name of the filter.
	FilterName *string `field:"required" json:"filterName" yaml:"filterName"`
	// The name of the field that the filter operates on.
	OperandFieldName *string `field:"required" json:"operandFieldName" yaml:"operandFieldName"`
	// The category filter that is associated with this filter.
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// The date range filter.
	DateRangeFilter interface{} `field:"optional" json:"dateRangeFilter" yaml:"dateRangeFilter"`
	// The class of the filter.
	//
	// Valid values for this structure are `ENFORCED_VALUE_FILTER` , `CONDITIONAL_VALUE_FILTER` , and `NAMED_VALUE_FILTER` .
	FilterClass *string `field:"optional" json:"filterClass" yaml:"filterClass"`
	// A description of the filter used to select items for a topic.
	FilterDescription *string `field:"optional" json:"filterDescription" yaml:"filterDescription"`
	// The other names or aliases for the filter.
	FilterSynonyms *[]*string `field:"optional" json:"filterSynonyms" yaml:"filterSynonyms"`
	// The type of the filter.
	//
	// Valid values for this structure are `CATEGORY_FILTER` , `NUMERIC_EQUALITY_FILTER` , `NUMERIC_RANGE_FILTER` , `DATE_RANGE_FILTER` , and `RELATIVE_DATE_FILTER` .
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// The numeric equality filter.
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// The numeric range filter.
	NumericRangeFilter interface{} `field:"optional" json:"numericRangeFilter" yaml:"numericRangeFilter"`
	// The relative date filter.
	RelativeDateFilter interface{} `field:"optional" json:"relativeDateFilter" yaml:"relativeDateFilter"`
}

