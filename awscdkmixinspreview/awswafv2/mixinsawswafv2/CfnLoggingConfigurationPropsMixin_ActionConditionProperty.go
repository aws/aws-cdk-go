package mixinsawswafv2


// A single action condition for a condition in a logging filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionConditionProperty := &ActionConditionProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-actioncondition.html
//
type CfnLoggingConfigurationPropsMixin_ActionConditionProperty struct {
	// The action setting that a log record must contain in order to meet the condition.
	//
	// This is the action that AWS WAF applied to the web request.
	//
	// For rule groups, this is either the configured rule action setting, or if you've applied a rule action override to the rule, it's the override action. The value `EXCLUDED_AS_COUNT` matches on excluded rules and also on rules that have a rule action override of Count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-actioncondition.html#cfn-wafv2-loggingconfiguration-actioncondition-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
}

