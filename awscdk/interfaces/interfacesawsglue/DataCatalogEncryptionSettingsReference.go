package interfacesawsglue


// A reference to a DataCatalogEncryptionSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogEncryptionSettingsReference := &DataCatalogEncryptionSettingsReference{
//   	CatalogId: jsii.String("catalogId"),
//   }
//
type DataCatalogEncryptionSettingsReference struct {
	// The CatalogId of the DataCatalogEncryptionSettings resource.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
}

