package awsroute53resolver


// Properties for defining a `CfnResolverQueryLoggingConfigAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigAssociationProps := &CfnResolverQueryLoggingConfigAssociationProps{
//   	ResolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html
//
type CfnResolverQueryLoggingConfigAssociationProps struct {
	// The ID of the query logging configuration that a VPC is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resolverquerylogconfigid
	//
	ResolverQueryLogConfigId *string `field:"optional" json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

