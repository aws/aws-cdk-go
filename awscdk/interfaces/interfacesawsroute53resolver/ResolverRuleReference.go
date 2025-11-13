package interfacesawsroute53resolver


// A reference to a ResolverRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverRuleReference := &ResolverRuleReference{
//   	ResolverRuleArn: jsii.String("resolverRuleArn"),
//   	ResolverRuleId: jsii.String("resolverRuleId"),
//   }
//
type ResolverRuleReference struct {
	// The ARN of the ResolverRule resource.
	ResolverRuleArn *string `field:"required" json:"resolverRuleArn" yaml:"resolverRuleArn"`
	// The ResolverRuleId of the ResolverRule resource.
	ResolverRuleId *string `field:"required" json:"resolverRuleId" yaml:"resolverRuleId"`
}

