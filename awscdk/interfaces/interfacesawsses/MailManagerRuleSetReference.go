package interfacesawsses


// A reference to a MailManagerRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerRuleSetReference := &MailManagerRuleSetReference{
//   	RuleSetArn: jsii.String("ruleSetArn"),
//   	RuleSetId: jsii.String("ruleSetId"),
//   }
//
type MailManagerRuleSetReference struct {
	// The ARN of the MailManagerRuleSet resource.
	RuleSetArn *string `field:"required" json:"ruleSetArn" yaml:"ruleSetArn"`
	// The RuleSetId of the MailManagerRuleSet resource.
	RuleSetId *string `field:"required" json:"ruleSetId" yaml:"ruleSetId"`
}

