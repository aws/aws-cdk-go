package interfacesawsglue


// A reference to a DataCatalogEncryptionSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogEncryptionSettingsReference := &DataCatalogEncryptionSettingsReference{
//   	DataCatalogEncryptionSettingsId: jsii.String("dataCatalogEncryptionSettingsId"),
//   }
//
type DataCatalogEncryptionSettingsReference struct {
	// The Id of the DataCatalogEncryptionSettings resource.
	DataCatalogEncryptionSettingsId *string `field:"required" json:"dataCatalogEncryptionSettingsId" yaml:"dataCatalogEncryptionSettingsId"`
}

