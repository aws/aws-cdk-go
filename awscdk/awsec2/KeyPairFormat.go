package awsec2


// The format of the Key Pair.
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
type KeyPairFormat string

const (
	// A PPK file, typically used with PuTTY.
	KeyPairFormat_PPK KeyPairFormat = "PPK"
	// A PEM file.
	KeyPairFormat_PEM KeyPairFormat = "PEM"
)

