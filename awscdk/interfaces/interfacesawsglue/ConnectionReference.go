package interfacesawsglue


// A reference to a Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionReference := &ConnectionReference{
//   	CatalogId: jsii.String("catalogId"),
//   	ConnectionName: jsii.String("connectionName"),
//   }
//
type ConnectionReference struct {
	// The CatalogId of the Connection resource.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The Name of the Connection resource.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
}

