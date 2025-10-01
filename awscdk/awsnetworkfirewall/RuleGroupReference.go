package awsnetworkfirewall


// A reference to a RuleGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReference := &RuleGroupReference{
//   	RuleGroupArn: jsii.String("ruleGroupArn"),
//   }
//
type RuleGroupReference struct {
	// The RuleGroupArn of the RuleGroup resource.
	RuleGroupArn *string `field:"required" json:"ruleGroupArn" yaml:"ruleGroupArn"`
}

