package awsroute53resolver


// Properties for defining a `CfnResolverRuleAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleAssociationProps := &cfnResolverRuleAssociationProps{
//   	resolverRuleId: jsii.String("resolverRuleId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnResolverRuleAssociationProps struct {
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	ResolverRuleId *string `field:"required" json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The name of an association between a Resolver rule and a VPC.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

