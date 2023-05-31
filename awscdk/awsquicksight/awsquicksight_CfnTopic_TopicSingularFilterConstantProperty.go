package awsquicksight


// A structure that represents a singular filter constant, used in filters to specify a single value to match against.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicSingularFilterConstantProperty := &TopicSingularFilterConstantProperty{
//   	ConstantType: jsii.String("constantType"),
//   	SingularConstant: jsii.String("singularConstant"),
//   }
//
type CfnTopic_TopicSingularFilterConstantProperty struct {
	// The type of the singular filter constant.
	//
	// Valid values for this structure are `SINGULAR` .
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// The value of the singular filter constant.
	SingularConstant *string `field:"optional" json:"singularConstant" yaml:"singularConstant"`
}

