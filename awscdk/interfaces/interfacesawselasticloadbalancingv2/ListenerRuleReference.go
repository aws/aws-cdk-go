package interfacesawselasticloadbalancingv2


// A reference to a ListenerRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerRuleReference := &ListenerRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type ListenerRuleReference struct {
	// The RuleArn of the ListenerRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

