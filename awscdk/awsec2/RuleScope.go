package awsec2


// The scope and id in which a given SecurityGroup rule should be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   ruleScope := &RuleScope{
//   	Id: jsii.String("id"),
//   	Scope: securityGroup,
//   }
//
type RuleScope struct {
	// The construct ID to use for the rule.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The SecurityGroup in which a rule should be scoped.
	Scope ISecurityGroup `field:"required" json:"scope" yaml:"scope"`
}

