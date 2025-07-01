package awsroute53resolver


// Properties for defining a `CfnResolverRuleAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleAssociationProps := &CfnResolverRuleAssociationProps{
//   	ResolverRuleId: jsii.String("resolverRuleId"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html
//
type CfnResolverRuleAssociationProps struct {
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-resolverruleid
	//
	ResolverRuleId *string `field:"required" json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The name of an association between a Resolver rule and a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

