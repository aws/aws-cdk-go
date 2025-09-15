package awsroute53resolver


// A reference to a ResolverQueryLoggingConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverQueryLoggingConfigReference := &ResolverQueryLoggingConfigReference{
//   	ResolverQueryLoggingConfigArn: jsii.String("resolverQueryLoggingConfigArn"),
//   	ResolverQueryLoggingConfigId: jsii.String("resolverQueryLoggingConfigId"),
//   }
//
type ResolverQueryLoggingConfigReference struct {
	// The ARN of the ResolverQueryLoggingConfig resource.
	ResolverQueryLoggingConfigArn *string `field:"required" json:"resolverQueryLoggingConfigArn" yaml:"resolverQueryLoggingConfigArn"`
	// The Id of the ResolverQueryLoggingConfig resource.
	ResolverQueryLoggingConfigId *string `field:"required" json:"resolverQueryLoggingConfigId" yaml:"resolverQueryLoggingConfigId"`
}

