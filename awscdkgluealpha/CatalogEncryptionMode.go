package awscdkgluealpha


// The encryption-at-rest mode for a Glue Data Catalog.
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_EncryptionAtRest.html#Glue-Type-EncryptionAtRest-CatalogEncryptionMode
//
// Experimental.
type CatalogEncryptionMode string

const (
	// Encryption at rest is disabled.
	// Experimental.
	CatalogEncryptionMode_DISABLED CatalogEncryptionMode = "DISABLED"
	// Server-side encryption (SSE) with an AWS KMS key.
	// Experimental.
	CatalogEncryptionMode_SSE_KMS CatalogEncryptionMode = "SSE_KMS"
	// Server-side encryption (SSE) with an AWS KMS key, using a service role that AWS Glue assumes to access the key on your behalf.
	// Experimental.
	CatalogEncryptionMode_SSE_KMS_WITH_SERVICE_ROLE CatalogEncryptionMode = "SSE_KMS_WITH_SERVICE_ROLE"
)

