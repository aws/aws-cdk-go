package mixinsawswafv2


// This is part of the configuration for the managed rules `AWSManagedRulesAntiDDoSRuleSet` in `ManagedRuleGroupConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientSideActionConfigProperty := &ClientSideActionConfigProperty{
//   	Challenge: &ClientSideActionProperty{
//   		ExemptUriRegularExpressions: []interface{}{
//   			&RegexProperty{
//   				RegexString: jsii.String("regexString"),
//   			},
//   		},
//   		Sensitivity: jsii.String("sensitivity"),
//   		UsageOfAction: jsii.String("usageOfAction"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideactionconfig.html
//
type CfnWebACLPropsMixin_ClientSideActionConfigProperty struct {
	// Configuration for the use of the `AWSManagedRulesAntiDDoSRuleSet` rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` .
	//
	// > This setting isn't related to the configuration of the `Challenge` action itself. It only configures the use of the two anti-DDoS rules named here.
	//
	// You can enable or disable the use of these rules, and you can configure how to use them when they are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideactionconfig.html#cfn-wafv2-webacl-clientsideactionconfig-challenge
	//
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
}

