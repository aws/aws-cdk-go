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
type CfnRule_ActionProperty struct {
	// The forward action.
	//
	// Traffic that matches the rule is forwarded to the specified target groups.
	Forward interface{} `field:"required" json:"forward" yaml:"forward"`
}

