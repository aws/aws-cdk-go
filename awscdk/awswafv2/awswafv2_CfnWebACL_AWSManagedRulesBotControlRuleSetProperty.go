package awswafv2


// Details for your use of the Bot Control managed rule group, used in `ManagedRuleGroupConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesBotControlRuleSetProperty := &aWSManagedRulesBotControlRuleSetProperty{
//   	inspectionLevel: jsii.String("inspectionLevel"),
//   }
//
type CfnWebACL_AWSManagedRulesBotControlRuleSetProperty struct {
	// The inspection level to use for the Bot Control rule group.
	//
	// The common level is the least expensive. The targeted level includes all common level rules and adds rules with more advanced inspection criteria. For details, see [AWS WAF Bot Control rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html) .
	InspectionLevel *string `field:"required" json:"inspectionLevel" yaml:"inspectionLevel"`
}

