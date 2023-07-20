package awsvpclattice


// The forward action.
//
// Traffic that matches the rule is forwarded to the specified target groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardProperty := &ForwardProperty{
//   	TargetGroups: []interface{}{
//   		&WeightedTargetGroupProperty{
//   			TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   			// the properties below are optional
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-forward.html
//
type CfnRule_ForwardProperty struct {
	// The target groups.
	//
	// Traffic matching the rule is forwarded to the specified target groups. With forward actions, you can assign a weight that controls the prioritization and selection of each target group. This means that requests are distributed to individual target groups based on their weights. For example, if two target groups have the same weight, each target group receives half of the traffic.
	//
	// The default value is 1. This means that if only one target group is provided, there is no need to set the weight; 100% of traffic will go to that target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-forward.html#cfn-vpclattice-rule-forward-targetgroups
	//
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

