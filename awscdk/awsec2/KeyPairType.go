package awsec2


// The type of the key pair.
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
type KeyPairType string

const (
	// An RSA key.
	KeyPairType_RSA KeyPairType = "RSA"
	// An ED25519 key.
	//
	// Note that ED25519 keys are not supported for Windows instances.
	KeyPairType_ED25519 KeyPairType = "ED25519"
)

