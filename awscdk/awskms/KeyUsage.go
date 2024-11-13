package awskms


// The key usage, represents the cryptographic operations of keys.
//
// Example:
//   key := kms.NewKey(this, jsii.String("MyKey"), &KeyProps{
//   	KeySpec: kms.KeySpec_ECC_SECG_P256K1,
//   	 // Default to SYMMETRIC_DEFAULT
//   	KeyUsage: kms.KeyUsage_SIGN_VERIFY,
//   })
//
type KeyUsage string

const (
	// Encryption and decryption.
	KeyUsage_ENCRYPT_DECRYPT KeyUsage = "ENCRYPT_DECRYPT"
	// Signing and verification.
	KeyUsage_SIGN_VERIFY KeyUsage = "SIGN_VERIFY"
	// Generating and verifying MACs.
	KeyUsage_GENERATE_VERIFY_MAC KeyUsage = "GENERATE_VERIFY_MAC"
	// Deriving shared secrets.
	KeyUsage_KEY_AGREEMENT KeyUsage = "KEY_AGREEMENT"
)

