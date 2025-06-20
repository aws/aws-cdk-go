package awscloudfront


// Properties for creating a Public Key.
//
// Example:
//   // Create a key group to use with CloudFront signed URLs and signed cookies.
//   // Create a key group to use with CloudFront signed URLs and signed cookies.
//   cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &KeyGroupProps{
//   	Items: []iPublicKey{
//   		cloudfront.NewPublicKey(this, jsii.String("MyPublicKey"), &PublicKeyProps{
//   			EncodedKey: jsii.String("..."),
//   		}),
//   	},
//   })
//
type PublicKeyProps struct {
	// The public key that you can use with signed URLs and signed cookies, or with field-level encryption.
	//
	// The `encodedKey` parameter must include `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----` lines.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html
	//
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// A comment to describe the public key.
	// Default: - no comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the public key.
	// Default: - generated from the `id`.
	//
	PublicKeyName *string `field:"optional" json:"publicKeyName" yaml:"publicKeyName"`
}

