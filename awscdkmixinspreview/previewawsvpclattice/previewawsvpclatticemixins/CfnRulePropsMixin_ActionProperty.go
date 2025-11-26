package previewawsvpclatticemixins


// Describes the action for a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-action.html
//
type CfnRulePropsMixin_ActionProperty struct {
	// The fixed response action.
	//
	// The rule returns a custom HTTP response.
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

