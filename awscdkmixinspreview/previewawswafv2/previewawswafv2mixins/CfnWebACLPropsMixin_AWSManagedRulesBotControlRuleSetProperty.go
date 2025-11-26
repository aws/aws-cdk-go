package previewawswafv2mixins


// Details for your use of the Bot Control managed rule group, `AWSManagedRulesBotControlRuleSet` .
//
// This configuration is used in `ManagedRuleGroupConfig` .
//
// For additional information about this and the other intelligent threat mitigation rule groups, see [Intelligent threat mitigation in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-managed-protections) and [AWS Managed Rules rule groups list](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-list) in the *AWS WAF Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSManagedRulesBotControlRuleSetProperty := &AWSManagedRulesBotControlRuleSetProperty{
//   	EnableMachineLearning: jsii.Boolean(false),
//   	InspectionLevel: jsii.String("inspectionLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html
//
type CfnWebACLPropsMixin_AWSManagedRulesBotControlRuleSetProperty struct {
	// Applies only to the targeted inspection level.
	//
	// Determines whether to use machine learning (ML) to analyze your web traffic for bot-related activity. Machine learning is required for the Bot Control rules `TGT_ML_CoordinatedActivityLow` and `TGT_ML_CoordinatedActivityMedium` , which
	// inspect for anomalous behavior that might indicate distributed, coordinated bot activity.
	//
	// For more information about this choice, see the listing for these rules in the table at [Bot Control rules listing](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html#aws-managed-rule-groups-bot-rules) in the *AWS WAF Developer Guide* .
	//
	// Default: `TRUE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html#cfn-wafv2-webacl-awsmanagedrulesbotcontrolruleset-enablemachinelearning
	//
	EnableMachineLearning interface{} `field:"optional" json:"enableMachineLearning" yaml:"enableMachineLearning"`
	// The inspection level to use for the Bot Control rule group.
	//
	// The common level is the least expensive. The targeted level includes all common level rules and adds rules with more advanced inspection criteria. For details, see [AWS WAF Bot Control rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html#cfn-wafv2-webacl-awsmanagedrulesbotcontrolruleset-inspectionlevel
	//
	InspectionLevel *string `field:"optional" json:"inspectionLevel" yaml:"inspectionLevel"`
}

