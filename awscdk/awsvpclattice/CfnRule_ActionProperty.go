package awsvpclattice


// Describes the action for a rule.
//
// Each rule must include exactly one of the following types of actions: `forward` or `fixed-response` , and it must be the last action to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	FixedResponse: &FixedResponseProperty{
//   		StatusCode: jsii.Number(123),
//   	},
//   	Forward: &ForwardProperty{
//   		TargetGroups: []interface{}{
//   			&WeightedTargetGroupProperty{
//   				TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   				// the properties below are optional
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-action.html
//
type CfnRule_ActionProperty struct {
	// Describes the rule action that returns a custom HTTP response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-action.html#cfn-vpclattice-rule-action-fixedresponse
	//
	FixedResponse interface{} `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// The forward action.
	//
	// Traffic that matches the rule is forwarded to the specified target groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-action.html#cfn-vpclattice-rule-action-forward
	//
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
}

