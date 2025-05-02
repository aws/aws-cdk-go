package awswisdom


// An object that can be used to specify tag conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &TagFilterProperty{
//   	AndConditions: []interface{}{
//   		&TagConditionProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OrConditions: []interface{}{
//   		&OrConditionProperty{
//   			AndConditions: []interface{}{
//   				&TagConditionProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TagCondition: &TagConditionProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	TagCondition: &TagConditionProperty{
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagfilter.html
//
type CfnAIAgent_TagFilterProperty struct {
	// A list of conditions which would be applied together with an `AND` condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagfilter.html#cfn-wisdom-aiagent-tagfilter-andconditions
	//
	AndConditions interface{} `field:"optional" json:"andConditions" yaml:"andConditions"`
	// A list of conditions which would be applied together with an `OR` condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagfilter.html#cfn-wisdom-aiagent-tagfilter-orconditions
	//
	OrConditions interface{} `field:"optional" json:"orConditions" yaml:"orConditions"`
	// A leaf node condition which can be used to specify a tag condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagfilter.html#cfn-wisdom-aiagent-tagfilter-tagcondition
	//
	TagCondition interface{} `field:"optional" json:"tagCondition" yaml:"tagCondition"`
}

