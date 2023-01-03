package awswafv2


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
	// `CfnLoggingConfiguration.FilterProperty.Behavior`.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// `CfnLoggingConfiguration.FilterProperty.Conditions`.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// `CfnLoggingConfiguration.FilterProperty.Requirement`.
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
}

