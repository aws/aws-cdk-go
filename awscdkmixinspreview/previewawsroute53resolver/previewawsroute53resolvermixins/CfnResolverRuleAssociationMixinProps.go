package previewawsroute53resolvermixins


// Properties for CfnResolverRuleAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResolverRuleAssociationMixinProps := &CfnResolverRuleAssociationMixinProps{
//   	Name: jsii.String("name"),
//   	ResolverRuleId: jsii.String("resolverRuleId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html
//
type CfnResolverRuleAssociationMixinProps struct {
	// The name of an association between a Resolver rule and a VPC.
	//
	// The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and spaces. The name cannot consist of only numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-resolverruleid
	//
	ResolverRuleId *string `field:"optional" json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

