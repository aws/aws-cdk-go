package interfacesawsroute53resolver


// A reference to a ResolverEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverEndpointReference := &ResolverEndpointReference{
//   	ResolverEndpointArn: jsii.String("resolverEndpointArn"),
//   	ResolverEndpointId: jsii.String("resolverEndpointId"),
//   }
//
type ResolverEndpointReference struct {
	// The ARN of the ResolverEndpoint resource.
	ResolverEndpointArn *string `field:"required" json:"resolverEndpointArn" yaml:"resolverEndpointArn"`
	// The ResolverEndpointId of the ResolverEndpoint resource.
	ResolverEndpointId *string `field:"required" json:"resolverEndpointId" yaml:"resolverEndpointId"`
}

