package awscdkgluealpha


// Encryption configuration for a Glue Data Catalog.
//
// Encryption is fixed at construction: a catalog either carries encryption
// settings or it does not, which keeps its configuration easy to reason about
// and avoids order-dependent mutation after the catalog is created.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var key Key
//   var role IRole
//
//   glue.Catalog_EncryptAccount(this, &CatalogEncryptionOptions{
//   	EncryptionAtRest: glue.DataCatalogEncryptionAtRest_KmsWithServiceRole(role, key),
//   })
//
// Experimental.
type CatalogEncryptionOptions struct {
	// Connection-password encryption configuration for the catalog.
	// Default: - connection-password encryption is not managed by CDK.
	//
	// Experimental.
	ConnectionPasswordEncryption *ConnectionPasswordEncryption `field:"optional" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// Encryption-at-rest configuration for the catalog.
	// Default: - encryption at rest is not managed by CDK (the catalog default applies).
	//
	// Experimental.
	EncryptionAtRest DataCatalogEncryptionAtRest `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

