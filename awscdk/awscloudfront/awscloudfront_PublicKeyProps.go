package awscloudfront


// Properties for creating a Public Key.
//
// Example:
//   // Validating signed URLs or signed cookies with Trusted Key Groups
//
//   // public key in PEM format
//   var publicKey string
//
//   pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &PublicKeyProps{
//   	EncodedKey: publicKey,
//   })
//
//   keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &KeyGroupProps{
//   	Items: []iPublicKey{
//   		pubKey,
//   	},
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Dist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   		TrustedKeyGroups: []iKeyGroup{
//   			keyGroup,
//   		},
//   	},
//   })
//
// Experimental.
type PublicKeyProps struct {
	// The public key that you can use with signed URLs and signed cookies, or with field-level encryption.
	//
	// The `encodedKey` parameter must include `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----` lines.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html
	//
	// Experimental.
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// A comment to describe the public key.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the public key.
	// Experimental.
	PublicKeyName *string `field:"optional" json:"publicKeyName" yaml:"publicKeyName"`
}

