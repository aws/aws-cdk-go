package awscustomerprofiles


// A group of intent filter conditions combined by the group's logical operator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentFilterGroupProperty := &IntentFilterGroupProperty{
//   	Conditions: []interface{}{
//   		&IntentFilterConditionProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			Intent: &IntentProperty{
//   				IntentDomain: jsii.String("intentDomain"),
//   				Type: jsii.String("type"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	LogicalOperator: jsii.String("logicalOperator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltergroup.html
//
type CfnEventTrigger_IntentFilterGroupProperty struct {
	// A list of intent filter conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltergroup.html#cfn-customerprofiles-eventtrigger-intentfiltergroup-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The operator used to combine multiple dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-intentfiltergroup.html#cfn-customerprofiles-eventtrigger-intentfiltergroup-logicaloperator
	//
	LogicalOperator *string `field:"required" json:"logicalOperator" yaml:"logicalOperator"`
}

