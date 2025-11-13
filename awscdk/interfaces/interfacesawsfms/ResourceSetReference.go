package interfacesawsfms


// A reference to a ResourceSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSetReference := &ResourceSetReference{
//   	ResourceSetId: jsii.String("resourceSetId"),
//   }
//
type ResourceSetReference struct {
	// The Id of the ResourceSet resource.
	ResourceSetId *string `field:"required" json:"resourceSetId" yaml:"resourceSetId"`
}

