package mixinsawswafv2


// Filtering that specifies which web requests are kept in the logs and which are dropped, defined for a web ACL's `LoggingConfiguration` .
//
// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingFilterProperty := &LoggingFilterProperty{
//   	DefaultBehavior: jsii.String("defaultBehavior"),
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Behavior: jsii.String("behavior"),
//   			Conditions: []interface{}{
//   				&ConditionProperty{
//   					ActionCondition: &ActionConditionProperty{
//   						Action: jsii.String("action"),
//   					},
//   					LabelNameCondition: &LabelNameConditionProperty{
//   						LabelName: jsii.String("labelName"),
//   					},
//   				},
//   			},
//   			Requirement: jsii.String("requirement"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-loggingfilter.html
//
type CfnLoggingConfigurationPropsMixin_LoggingFilterProperty struct {
	// Default handling for logs that don't match any of the specified filtering conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-loggingfilter.html#cfn-wafv2-loggingconfiguration-loggingfilter-defaultbehavior
	//
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// The filters that you want to apply to the logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-loggingfilter.html#cfn-wafv2-loggingconfiguration-loggingfilter-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

