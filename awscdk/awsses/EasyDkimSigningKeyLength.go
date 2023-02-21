package awsses


// The signing key length for Easy DKIM.
type EasyDkimSigningKeyLength string

const (
	// RSA 1024-bit.
	EasyDkimSigningKeyLength_RSA_1024_BIT EasyDkimSigningKeyLength = "RSA_1024_BIT"
	// RSA 2048-bit.
	EasyDkimSigningKeyLength_RSA_2048_BIT EasyDkimSigningKeyLength = "RSA_2048_BIT"
)

