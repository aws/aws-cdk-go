package awscustomerprofiles


// Filters the destination to profiles whose predicted intents meet the specified conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentFilterDefinitionProperty := &IntentFilterDefinitionProperty{
//   	Groups: []interface{}{
//   		&IntentFilterGroupProperty{
//   			Conditions: []interface{}{
//   				&IntentFilterConditionProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					Intent: &IntentProperty{
//   						IntentDomain: jsii.String("intentDomain"),
//   						Type: jsii.String("type"),
//   					},
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			LogicalOperator: jsii.String("logicalOperator"),
//   		},
//   	},
//   	LogicalOperator: jsii.String("logicalOperator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfilterdefinition.html
//
type CfnEventTrigger_IntentFilterDefinitionProperty struct {
	// A list of intent filter groups combined by the definition's logical operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfilterdefinition.html#cfn-customerprofiles-eventtrigger-intentfilterdefinition-groups
	//
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// The operator used to combine multiple dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfilterdefinition.html#cfn-customerprofiles-eventtrigger-intentfilterdefinition-logicaloperator
	//
	LogicalOperator *string `field:"required" json:"logicalOperator" yaml:"logicalOperator"`
}

