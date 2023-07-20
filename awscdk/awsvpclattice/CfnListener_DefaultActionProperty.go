package awsvpclattice


// The action for the default rule.
//
// Each listener has a default rule. Each rule consists of a priority, one or more actions, and one or more conditions. The default rule is the rule that's used if no other rules match. Each rule must include exactly one of the following types of actions: `forward` or `fixed-response` , and it must be the last action to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultActionProperty := &DefaultActionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html
//
type CfnListener_DefaultActionProperty struct {
	// Information about an action that returns a custom HTTP response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html#cfn-vpclattice-listener-defaultaction-fixedresponse
	//
	FixedResponse interface{} `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Describes a forward action.
	//
	// You can use forward actions to route requests to one or more target groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html#cfn-vpclattice-listener-defaultaction-forward
	//
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
}

