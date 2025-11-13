package interfacesawsroute53resolver


// A reference to a OutpostResolver resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outpostResolverReference := &OutpostResolverReference{
//   	OutpostResolverArn: jsii.String("outpostResolverArn"),
//   	OutpostResolverId: jsii.String("outpostResolverId"),
//   }
//
type OutpostResolverReference struct {
	// The ARN of the OutpostResolver resource.
	OutpostResolverArn *string `field:"required" json:"outpostResolverArn" yaml:"outpostResolverArn"`
	// The Id of the OutpostResolver resource.
	OutpostResolverId *string `field:"required" json:"outpostResolverId" yaml:"outpostResolverId"`
}

