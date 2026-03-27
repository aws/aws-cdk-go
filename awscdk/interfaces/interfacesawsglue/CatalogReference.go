package interfacesawsglue


// A reference to a Catalog resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogReference := &CatalogReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type CatalogReference struct {
	// The ResourceArn of the Catalog resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

