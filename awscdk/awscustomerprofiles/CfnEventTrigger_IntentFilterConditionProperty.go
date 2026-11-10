package awscustomerprofiles


// The criteria that a predicted intent must meet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentFilterConditionProperty := &IntentFilterConditionProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	Intent: &IntentProperty{
//   		IntentDomain: jsii.String("intentDomain"),
//   		Type: jsii.String("type"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltercondition.html
//
type CfnEventTrigger_IntentFilterConditionProperty struct {
	// The operator used to compare the predicted intent against a list of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltercondition.html#cfn-customerprofiles-eventtrigger-intentfiltercondition-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Identifies a predicted intent by domain and type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltercondition.html#cfn-customerprofiles-eventtrigger-intentfiltercondition-intent
	//
	Intent interface{} `field:"required" json:"intent" yaml:"intent"`
	// A list of intent values used for comparison.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltercondition.html#cfn-customerprofiles-eventtrigger-intentfiltercondition-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

