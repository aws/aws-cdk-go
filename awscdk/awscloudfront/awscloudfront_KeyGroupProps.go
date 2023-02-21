package awscloudfront


// Properties for creating a Public Key.
//
// Example:
//   // Validating signed URLs or signed cookies with Trusted Key Groups
//
//   // public key in PEM format
//   var publicKey string
//
//   pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &publicKeyProps{
//   	encodedKey: publicKey,
//   })
//
//   keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &keyGroupProps{
//   	items: []iPublicKey{
//   		pubKey,
//   	},
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Dist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   		trustedKeyGroups: []iKeyGroup{
//   			keyGroup,
//   		},
//   	},
//   })
//
// Experimental.
type KeyGroupProps struct {
	// A list of public keys to add to the key group.
	// Experimental.
	Items *[]IPublicKey `field:"required" json:"items" yaml:"items"`
	// A comment to describe the key group.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the key group.
	// Experimental.
	KeyGroupName *string `field:"optional" json:"keyGroupName" yaml:"keyGroupName"`
}

