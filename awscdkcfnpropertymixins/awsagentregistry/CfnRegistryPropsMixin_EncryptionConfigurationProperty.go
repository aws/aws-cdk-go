package awsagentregistry


// The server-side encryption configuration for a registry.
//
// Specifies a customer managed key used to encrypt the registry's content. When omitted, the registry's content is encrypted with an AWS owned key. You cannot change the encryption configuration after registry creation. Specifying a different KMS key, adding this property to an existing registry, or removing it replaces the registry: CloudFormation creates a new registry with a new Amazon Resource Name (ARN) and then deletes the original, including all registry records it contains. Registry records that are not managed by the stack are not re-created in the new registry, and if any remain in the original registry its deletion fails and it is left behind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-encryptionconfiguration.html
//
type CfnRegistryPropsMixin_EncryptionConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the customer-managed AWS KMS key used to encrypt the registry's content.
	//
	// The key must be a symmetric encryption key in the same AWS account and Region as the registry. Multi-Region keys are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-encryptionconfiguration.html#cfn-agentregistry-registry-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

