package awsroute53resolver


// A reference to a ResolverQueryLoggingConfigAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverQueryLoggingConfigAssociationReference := &ResolverQueryLoggingConfigAssociationReference{
//   	ResolverQueryLoggingConfigAssociationId: jsii.String("resolverQueryLoggingConfigAssociationId"),
//   }
//
type ResolverQueryLoggingConfigAssociationReference struct {
	// The Id of the ResolverQueryLoggingConfigAssociation resource.
	ResolverQueryLoggingConfigAssociationId *string `field:"required" json:"resolverQueryLoggingConfigAssociationId" yaml:"resolverQueryLoggingConfigAssociationId"`
}

