package awsnetworkfirewall


// The setting that allows the policy owner to change the behavior of the rule group within a policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulRuleGroupOverrideProperty := &statefulRuleGroupOverrideProperty{
//   	action: jsii.String("action"),
//   }
//
type CfnFirewallPolicy_StatefulRuleGroupOverrideProperty struct {
	// The action that changes the rule group from `DROP` to `ALERT` .
	//
	// This only applies to managed rule groups.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

