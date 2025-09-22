package awsroute53resolver


// A reference to a ResolverConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverConfigReference := &ResolverConfigReference{
//   	ResourceId: jsii.String("resourceId"),
//   }
//
type ResolverConfigReference struct {
	// The ResourceId of the ResolverConfig resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

