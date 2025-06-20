package awswafv2


// Configures how to use the AntiDDOS AWS managed rule group in the web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesAntiDDoSRuleSetProperty := &AWSManagedRulesAntiDDoSRuleSetProperty{
//   	ClientSideActionConfig: &ClientSideActionConfigProperty{
//   		Challenge: &ClientSideActionProperty{
//   			UsageOfAction: jsii.String("usageOfAction"),
//
//   			// the properties below are optional
//   			ExemptUriRegularExpressions: []interface{}{
//   				&RegexProperty{
//   					RegexString: jsii.String("regexString"),
//   				},
//   			},
//   			Sensitivity: jsii.String("sensitivity"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	SensitivityToBlock: jsii.String("sensitivityToBlock"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.html
//
type CfnWebACL_AWSManagedRulesAntiDDoSRuleSetProperty struct {
	// Client side action config for AntiDDOS AMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.html#cfn-wafv2-webacl-awsmanagedrulesantiddosruleset-clientsideactionconfig
	//
	ClientSideActionConfig interface{} `field:"required" json:"clientSideActionConfig" yaml:"clientSideActionConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.html#cfn-wafv2-webacl-awsmanagedrulesantiddosruleset-sensitivitytoblock
	//
	SensitivityToBlock *string `field:"optional" json:"sensitivityToBlock" yaml:"sensitivityToBlock"`
}

