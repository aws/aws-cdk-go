package awswafv2


// Configures the use of the anti-DDoS managed rule group, `AWSManagedRulesAntiDDoSRuleSet` . This configuration is used in `ManagedRuleGroupConfig` .
//
// The configuration that you provide here determines whether and how the rules in the rule group are used.
//
// For additional information about this and the other intelligent threat mitigation rule groups, see [Intelligent threat mitigation in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-managed-protections) and [AWS Managed Rules rule groups list](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-list) in the *AWS WAF Developer Guide* .
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
	// Configures the request handling that's applied by the managed rule group rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` during a distributed denial of service (DDoS) attack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.html#cfn-wafv2-webacl-awsmanagedrulesantiddosruleset-clientsideactionconfig
	//
	ClientSideActionConfig interface{} `field:"required" json:"clientSideActionConfig" yaml:"clientSideActionConfig"`
	// The sensitivity that the rule group rule `DDoSRequests` uses when matching against the DDoS suspicion labeling on a request.
	//
	// The managed rule group adds the labeling during DDoS events, before the `DDoSRequests` rule runs.
	//
	// The higher the sensitivity, the more levels of labeling that the rule matches:
	//
	// - Low sensitivity is less sensitive, causing the rule to match only on the most likely participants in an attack, which are the requests with the high suspicion label `awswaf:managed:aws:anti-ddos:high-suspicion-ddos-request` .
	// - Medium sensitivity causes the rule to match on the medium and high suspicion labels.
	// - High sensitivity causes the rule to match on all of the suspicion labels: low, medium, and high.
	//
	// Default: `LOW`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.html#cfn-wafv2-webacl-awsmanagedrulesantiddosruleset-sensitivitytoblock
	//
	SensitivityToBlock *string `field:"optional" json:"sensitivityToBlock" yaml:"sensitivityToBlock"`
}

