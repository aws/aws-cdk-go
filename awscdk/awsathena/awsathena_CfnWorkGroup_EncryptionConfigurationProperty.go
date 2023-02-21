package awsathena


// If query results are encrypted in Amazon S3, indicates the encryption option used (for example, `SSE_KMS` or `CSE_KMS` ) and key information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	encryptionOption: jsii.String("encryptionOption"),
//
//   	// the properties below are optional
//   	kmsKey: jsii.String("kmsKey"),
//   }
//
type CfnWorkGroup_EncryptionConfigurationProperty struct {
	// Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys ( `SSE_S3` ), server-side encryption with KMS-managed keys ( `SSE_KMS` ), or client-side encryption with KMS-managed keys ( `CSE_KMS` ) is used.
	//
	// If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// For `SSE_KMS` and `CSE_KMS` , this is the KMS key ARN or ID.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

