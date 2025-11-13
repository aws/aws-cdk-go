package interfacesawsroute53resolver


// A reference to a ResolverRuleAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverRuleAssociationReference := &ResolverRuleAssociationReference{
//   	ResolverRuleAssociationId: jsii.String("resolverRuleAssociationId"),
//   }
//
type ResolverRuleAssociationReference struct {
	// The ResolverRuleAssociationId of the ResolverRuleAssociation resource.
	ResolverRuleAssociationId *string `field:"required" json:"resolverRuleAssociationId" yaml:"resolverRuleAssociationId"`
}

