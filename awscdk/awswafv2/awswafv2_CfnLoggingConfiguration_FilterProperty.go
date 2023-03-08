package awswafv2


// A single logging filter, used in `LoggingFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	behavior: jsii.String("behavior"),
//   	conditions: []interface{}{
//   		&conditionProperty{
//   			actionCondition: &actionConditionProperty{
//   				action: jsii.String("action"),
//   			},
//   			labelNameCondition: &labelNameConditionProperty{
//   				labelName: jsii.String("labelName"),
//   			},
//   		},
//   	},
//   	requirement: jsii.String("requirement"),
//   }
//
type CfnLoggingConfiguration_FilterProperty struct {
	// How to handle logs that satisfy the filter's conditions and requirement.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Match conditions for the filter.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
}

