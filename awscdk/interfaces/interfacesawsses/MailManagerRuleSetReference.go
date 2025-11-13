package interfacesawsses


// A reference to a MailManagerRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerRuleSetReference := &MailManagerRuleSetReference{
//   	RuleSetId: jsii.String("ruleSetId"),
//   }
//
type MailManagerRuleSetReference struct {
	// The RuleSetId of the MailManagerRuleSet resource.
	RuleSetId *string `field:"required" json:"ruleSetId" yaml:"ruleSetId"`
}

