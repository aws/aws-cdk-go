package previewawsquicksightmixins


// A structure that represents a filter used to select items for a topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicFilterProperty := &TopicFilterProperty{
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
//   	FilterName: jsii.String("filterName"),
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
//   	OperandFieldName: jsii.String("operandFieldName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html
//
type CfnTopicPropsMixin_TopicFilterProperty struct {
	// The category filter that is associated with this filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-categoryfilter
	//
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// The date range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-daterangefilter
	//
	DateRangeFilter interface{} `field:"optional" json:"dateRangeFilter" yaml:"dateRangeFilter"`
	// The class of the filter.
	//
	// Valid values for this structure are `ENFORCED_VALUE_FILTER` , `CONDITIONAL_VALUE_FILTER` , and `NAMED_VALUE_FILTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-filterclass
	//
	FilterClass *string `field:"optional" json:"filterClass" yaml:"filterClass"`
	// A description of the filter used to select items for a topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-filterdescription
	//
	FilterDescription *string `field:"optional" json:"filterDescription" yaml:"filterDescription"`
	// The name of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-filtername
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The other names or aliases for the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-filtersynonyms
	//
	FilterSynonyms *[]*string `field:"optional" json:"filterSynonyms" yaml:"filterSynonyms"`
	// The type of the filter.
	//
	// Valid values for this structure are `CATEGORY_FILTER` , `NUMERIC_EQUALITY_FILTER` , `NUMERIC_RANGE_FILTER` , `DATE_RANGE_FILTER` , and `RELATIVE_DATE_FILTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-filtertype
	//
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// The numeric equality filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-numericequalityfilter
	//
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// The numeric range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-numericrangefilter
	//
	NumericRangeFilter interface{} `field:"optional" json:"numericRangeFilter" yaml:"numericRangeFilter"`
	// The name of the field that the filter operates on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-operandfieldname
	//
	OperandFieldName *string `field:"optional" json:"operandFieldName" yaml:"operandFieldName"`
	// The relative date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicfilter.html#cfn-quicksight-topic-topicfilter-relativedatefilter
	//
	RelativeDateFilter interface{} `field:"optional" json:"relativeDateFilter" yaml:"relativeDateFilter"`
}

