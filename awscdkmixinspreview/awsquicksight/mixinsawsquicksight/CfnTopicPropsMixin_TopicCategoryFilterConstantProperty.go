package mixinsawsquicksight


// A constant used in a category filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicCategoryFilterConstantProperty := &TopicCategoryFilterConstantProperty{
//   	CollectiveConstant: &CollectiveConstantProperty{
//   		ValueList: []*string{
//   			jsii.String("valueList"),
//   		},
//   	},
//   	ConstantType: jsii.String("constantType"),
//   	SingularConstant: jsii.String("singularConstant"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html
//
type CfnTopicPropsMixin_TopicCategoryFilterConstantProperty struct {
	// A collective constant used in a category filter.
	//
	// This element is used to specify a list of values for the constant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-collectiveconstant
	//
	CollectiveConstant interface{} `field:"optional" json:"collectiveConstant" yaml:"collectiveConstant"`
	// The type of category filter constant.
	//
	// This element is used to specify whether a constant is a singular or collective. Valid values are `SINGULAR` and `COLLECTIVE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-constanttype
	//
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// A singular constant used in a category filter.
	//
	// This element is used to specify a single value for the constant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-singularconstant
	//
	SingularConstant *string `field:"optional" json:"singularConstant" yaml:"singularConstant"`
}

