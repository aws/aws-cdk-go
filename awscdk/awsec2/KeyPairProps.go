package awsec2


// The properties of a Key Pair.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   keyPair := ec2.NewKeyPair(this, jsii.String("KeyPair"), &KeyPairProps{
//   	Type: ec2.KeyPairType_ED25519,
//   	Format: ec2.KeyPairFormat_PEM,
//   })
//   instance := ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	// Use the custom key pair
//   	KeyPair: KeyPair,
//   })
//
type KeyPairProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The format of the key pair.
	// Default: PEM.
	//
	Format KeyPairFormat `field:"optional" json:"format" yaml:"format"`
	// A unique name for the key pair.
	// Default: A generated name.
	//
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// The public key material.
	//
	// If this is provided the key is considered "imported". For imported
	// keys, it is assumed that you already have the private key material
	// so the private key material will not be returned or stored in
	// AWS Systems Manager Parameter Store.
	// Default: a public and private key will be generated.
	//
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// The type of key pair.
	// Default: RSA (ignored if keyMaterial is provided).
	//
	Type KeyPairType `field:"optional" json:"type" yaml:"type"`
}

