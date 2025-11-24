package mixinsawsvpclattice


// The action for the default rule.
//
// Each listener has a default rule. The default rule is used if no other rules match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultActionProperty := &DefaultActionProperty{
//   	FixedResponse: &FixedResponseProperty{
//   		StatusCode: jsii.Number(123),
//   	},
//   	Forward: &ForwardProperty{
//   		TargetGroups: []interface{}{
//   			&WeightedTargetGroupProperty{
//   				TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html
//
type CfnListenerPropsMixin_DefaultActionProperty struct {
	// Describes an action that returns a custom HTTP response.
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

