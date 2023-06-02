package awsquicksight


// A filter that filters topics based on the value of a numeric field.
//
// The filter includes only topics whose numeric field value matches the specified value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicNumericEqualityFilterProperty := &TopicNumericEqualityFilterProperty{
//   	Aggregation: jsii.String("aggregation"),
//   	Constant: &TopicSingularFilterConstantProperty{
//   		ConstantType: jsii.String("constantType"),
//   		SingularConstant: jsii.String("singularConstant"),
//   	},
//   }
//
type CfnTopic_TopicNumericEqualityFilterProperty struct {
	// An aggregation function that specifies how to calculate the value of a numeric field for a topic.
	//
	// Valid values for this structure are `NO_AGGREGATION` , `SUM` , `AVERAGE` , `COUNT` , `DISTINCT_COUNT` , `MAX` , `MEDIAN` , `MIN` , `STDEV` , `STDEVP` , `VAR` , and `VARP` .
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The constant used in a numeric equality filter.
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
}

