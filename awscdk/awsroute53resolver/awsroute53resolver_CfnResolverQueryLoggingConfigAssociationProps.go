package awsroute53resolver


// Properties for defining a `CfnResolverQueryLoggingConfigAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigAssociationProps := &cfnResolverQueryLoggingConfigAssociationProps{
//   	resolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverQueryLoggingConfigAssociationProps struct {
	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId *string `field:"optional" json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

