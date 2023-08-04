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
type KeyGroupProps struct {
	// A list of public keys to add to the key group.
	Items *[]IPublicKey `field:"required" json:"items" yaml:"items"`
	// A comment to describe the key group.
	// Default: - no comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the key group.
	// Default: - generated from the `id`.
	//
	KeyGroupName *string `field:"optional" json:"keyGroupName" yaml:"keyGroupName"`
}

