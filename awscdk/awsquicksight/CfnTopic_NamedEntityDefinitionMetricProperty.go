package awsquicksight


// A structure that represents a metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namedEntityDefinitionMetricProperty := &NamedEntityDefinitionMetricProperty{
//   	Aggregation: jsii.String("aggregation"),
//   	AggregationFunctionParameters: map[string]*string{
//   		"aggregationFunctionParametersKey": jsii.String("aggregationFunctionParameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinitionmetric.html
//
type CfnTopic_NamedEntityDefinitionMetricProperty struct {
	// The aggregation of a named entity.
	//
	// Valid values for this structure are `SUM` , `MIN` , `MAX` , `COUNT` , `AVERAGE` , `DISTINCT_COUNT` , `STDEV` , `STDEVP` , `VAR` , `VARP` , `PERCENTILE` , `MEDIAN` , and `CUSTOM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinitionmetric.html#cfn-quicksight-topic-namedentitydefinitionmetric-aggregation
	//
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The additional parameters for an aggregation function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinitionmetric.html#cfn-quicksight-topic-namedentitydefinitionmetric-aggregationfunctionparameters
	//
	AggregationFunctionParameters interface{} `field:"optional" json:"aggregationFunctionParameters" yaml:"aggregationFunctionParameters"`
}

