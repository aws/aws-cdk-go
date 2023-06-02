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
type CfnTopic_NamedEntityDefinitionMetricProperty struct {
	// The aggregation of a named entity.
	//
	// Valid values for this structure are `SUM` , `MIN` , `MAX` , `COUNT` , `AVERAGE` , `DISTINCT_COUNT` , `STDEV` , `STDEVP` , `VAR` , `VARP` , `PERCENTILE` , `MEDIAN` , and `CUSTOM` .
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The additional parameters for an aggregation function.
	AggregationFunctionParameters interface{} `field:"optional" json:"aggregationFunctionParameters" yaml:"aggregationFunctionParameters"`
}

