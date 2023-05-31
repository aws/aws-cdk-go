package awsquicksight


// A filter used to restrict data based on a range of dates or times.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicDateRangeFilterProperty := &TopicDateRangeFilterProperty{
//   	Constant: &TopicRangeFilterConstantProperty{
//   		ConstantType: jsii.String("constantType"),
//   		RangeConstant: &RangeConstantProperty{
//   			Maximum: jsii.String("maximum"),
//   			Minimum: jsii.String("minimum"),
//   		},
//   	},
//   	Inclusive: jsii.Boolean(false),
//   }
//
type CfnTopic_TopicDateRangeFilterProperty struct {
	// The constant used in a date range filter.
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
	// A Boolean value that indicates whether the date range filter should include the boundary values.
	//
	// If set to true, the filter includes the start and end dates. If set to false, the filter excludes them.
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
}

