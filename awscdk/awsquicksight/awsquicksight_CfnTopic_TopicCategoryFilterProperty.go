package awsquicksight


// A structure that represents a category filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicCategoryFilterProperty := &TopicCategoryFilterProperty{
//   	CategoryFilterFunction: jsii.String("categoryFilterFunction"),
//   	CategoryFilterType: jsii.String("categoryFilterType"),
//   	Constant: &TopicCategoryFilterConstantProperty{
//   		CollectiveConstant: &CollectiveConstantProperty{
//   			ValueList: []*string{
//   				jsii.String("valueList"),
//   			},
//   		},
//   		ConstantType: jsii.String("constantType"),
//   		SingularConstant: jsii.String("singularConstant"),
//   	},
//   	Inverse: jsii.Boolean(false),
//   }
//
type CfnTopic_TopicCategoryFilterProperty struct {
	// The category filter function.
	//
	// Valid values for this structure are `EXACT` and `CONTAINS` .
	CategoryFilterFunction *string `field:"optional" json:"categoryFilterFunction" yaml:"categoryFilterFunction"`
	// The category filter type.
	//
	// This element is used to specify whether a filter is a simple category filter or an inverse category filter.
	CategoryFilterType *string `field:"optional" json:"categoryFilterType" yaml:"categoryFilterType"`
	// The constant used in a category filter.
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
	// A Boolean value that indicates if the filter is inverse.
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
}

