package awswisdom


// The KMS key used for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionConfigurationProperty := &serverSideEncryptionConfigurationProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnAssistant_ServerSideEncryptionConfigurationProperty struct {
	// The KMS key .
	//
	// For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

