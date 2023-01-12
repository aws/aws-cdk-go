package awsopensearchservice


// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestOptionsProperty := &encryptionAtRestOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDomain_EncryptionAtRestOptionsProperty struct {
	// Specify `true` to enable encryption at rest.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ID.
	//
	// Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a` . Required if you enable encryption at rest.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

