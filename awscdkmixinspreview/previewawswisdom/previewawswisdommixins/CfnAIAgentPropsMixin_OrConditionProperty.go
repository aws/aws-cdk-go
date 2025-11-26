package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   orConditionProperty := &OrConditionProperty{
//   	AndConditions: []interface{}{
//   		&TagConditionProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TagCondition: &TagConditionProperty{
//   		Key: jsii.String("key"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orcondition.html
//
type CfnAIAgentPropsMixin_OrConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orcondition.html#cfn-wisdom-aiagent-orcondition-andconditions
	//
	AndConditions interface{} `field:"optional" json:"andConditions" yaml:"andConditions"`
	// A leaf node condition which can be used to specify a tag condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orcondition.html#cfn-wisdom-aiagent-orcondition-tagcondition
	//
	TagCondition interface{} `field:"optional" json:"tagCondition" yaml:"tagCondition"`
}

