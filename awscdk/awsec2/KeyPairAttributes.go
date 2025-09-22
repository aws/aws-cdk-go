package awsec2


// Attributes of a Key Pair.
//
// Example:
//   keyPair := ec2.KeyPair_FromKeyPairAttributes(this, jsii.String("KeyPair"), &KeyPairAttributes{
//   	KeyPairName: jsii.String("the-keypair-name"),
//   	Type: ec2.KeyPairType_RSA,
//   })
//
type KeyPairAttributes struct {
	// The unique name of the key pair.
	KeyPairName *string `field:"required" json:"keyPairName" yaml:"keyPairName"`
	// The type of the key pair.
	// Default: no type specified.
	//
	Type KeyPairType `field:"optional" json:"type" yaml:"type"`
}

