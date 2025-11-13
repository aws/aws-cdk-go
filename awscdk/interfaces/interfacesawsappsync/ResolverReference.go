package interfacesawsappsync


// A reference to a Resolver resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverReference := &ResolverReference{
//   	ResolverArn: jsii.String("resolverArn"),
//   }
//
type ResolverReference struct {
	// The ResolverArn of the Resolver resource.
	ResolverArn *string `field:"required" json:"resolverArn" yaml:"resolverArn"`
}

