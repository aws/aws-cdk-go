package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Connection-password encryption configuration for a Glue Data Catalog.
//
// When enabled, the Data Catalog encrypts the password as part of
// `CreateConnection` or `UpdateConnection` and stores it in the
// `ENCRYPTED_PASSWORD` field of the connection properties. This is independent
// from catalog encryption at rest, and may use a different KMS key.
//
// Example:
//   var key Key
//
//   glue.Catalog_EncryptAccount(this, &CatalogEncryptionOptions{
//   	ConnectionPasswordEncryption: &ConnectionPasswordEncryption{
//   		KmsKey: key,
//   		// Whether GetConnection/GetConnections return the password encrypted (default: true)
//   		ReturnConnectionPasswordEncrypted: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_ConnectionPasswordEncryption.html
//
// Experimental.
type ConnectionPasswordEncryption struct {
	// The KMS key used to encrypt connection passwords.
	// Default: - an AWS-managed key is used and the key is not exposed as a grantable resource.
	//
	// Experimental.
	KmsKey interfacesawskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Whether passwords remain encrypted in the responses of `GetConnection` and `GetConnections`.
	//
	// This takes effect independently from catalog encryption.
	// Default: true.
	//
	// Experimental.
	ReturnConnectionPasswordEncrypted *bool `field:"optional" json:"returnConnectionPasswordEncrypted" yaml:"returnConnectionPasswordEncrypted"`
}

