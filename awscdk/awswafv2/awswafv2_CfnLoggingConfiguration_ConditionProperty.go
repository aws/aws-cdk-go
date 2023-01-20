package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &conditionProperty{
//   	actionCondition: &actionConditionProperty{
//   		action: jsii.String("action"),
//   	},
//   	labelNameCondition: &labelNameConditionProperty{
//   		labelName: jsii.String("labelName"),
//   	},
//   }
//
type CfnLoggingConfiguration_ConditionProperty struct {
	// `CfnLoggingConfiguration.ConditionProperty.ActionCondition`.
	ActionCondition interface{} `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// `CfnLoggingConfiguration.ConditionProperty.LabelNameCondition`.
	LabelNameCondition interface{} `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

