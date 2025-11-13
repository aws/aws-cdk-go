package interfacesawsathena


// A reference to a DataCatalog resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogReference := &DataCatalogReference{
//   	DataCatalogName: jsii.String("dataCatalogName"),
//   }
//
type DataCatalogReference struct {
	// The Name of the DataCatalog resource.
	DataCatalogName *string `field:"required" json:"dataCatalogName" yaml:"dataCatalogName"`
}

