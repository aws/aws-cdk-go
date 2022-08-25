package awskms


// The key spec, represents the cryptographic configuration of keys.
//
// Example:
//   key := kms.NewKey(this, jsii.String("MyKey"), &keyProps{
//   	keySpec: kms.keySpec_ECC_SECG_P256K1,
//   	 // Default to SYMMETRIC_DEFAULT
//   	keyUsage: kms.keyUsage_SIGN_VERIFY,
//   })
//
type KeySpec string

const (
	// The default key spec.
	//
	// Valid usage: ENCRYPT_DECRYPT.
	KeySpec_SYMMETRIC_DEFAULT KeySpec = "SYMMETRIC_DEFAULT"
	// RSA with 2048 bits of key.
	//
	// Valid usage: ENCRYPT_DECRYPT and SIGN_VERIFY.
	KeySpec_RSA_2048 KeySpec = "RSA_2048"
	// RSA with 3072 bits of key.
	//
	// Valid usage: ENCRYPT_DECRYPT and SIGN_VERIFY.
	KeySpec_RSA_3072 KeySpec = "RSA_3072"
	// RSA with 4096 bits of key.
	//
	// Valid usage: ENCRYPT_DECRYPT and SIGN_VERIFY.
	KeySpec_RSA_4096 KeySpec = "RSA_4096"
	// NIST FIPS 186-4, Section 6.4, ECDSA signature using the curve specified by the key and SHA-256 for the message digest.
	//
	// Valid usage: SIGN_VERIFY.
	KeySpec_ECC_NIST_P256 KeySpec = "ECC_NIST_P256"
	// NIST FIPS 186-4, Section 6.4, ECDSA signature using the curve specified by the key and SHA-384 for the message digest.
	//
	// Valid usage: SIGN_VERIFY.
	KeySpec_ECC_NIST_P384 KeySpec = "ECC_NIST_P384"
	// NIST FIPS 186-4, Section 6.4, ECDSA signature using the curve specified by the key and SHA-512 for the message digest.
	//
	// Valid usage: SIGN_VERIFY.
	KeySpec_ECC_NIST_P521 KeySpec = "ECC_NIST_P521"
	// Standards for Efficient Cryptography 2, Section 2.4.1, ECDSA signature on the Koblitz curve.
	//
	// Valid usage: SIGN_VERIFY.
	KeySpec_ECC_SECG_P256K1 KeySpec = "ECC_SECG_P256K1"
)

