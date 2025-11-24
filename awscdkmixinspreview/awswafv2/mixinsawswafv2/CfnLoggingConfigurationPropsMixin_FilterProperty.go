package mixinsawswafv2


// A single logging filter, used in `LoggingFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterProperty := &FilterProperty{
//   	Behavior: jsii.String("behavior"),
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			ActionCondition: &ActionConditionProperty{
//   				Action: jsii.String("action"),
//   			},
//   			LabelNameCondition: &LabelNameConditionProperty{
//   				LabelName: jsii.String("labelName"),
//   			},
//   		},
//   	},
//   	Requirement: jsii.String("requirement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-filter.html
//
type CfnLoggingConfigurationPropsMixin_FilterProperty struct {
	// How to handle logs that satisfy the filter's conditions and requirement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-filter.html#cfn-wafv2-loggingconfiguration-filter-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// Match conditions for the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-filter.html#cfn-wafv2-loggingconfiguration-filter-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-filter.html#cfn-wafv2-loggingconfiguration-filter-requirement
	//
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

