package awsecr


// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
//
// By default, when no encryption configuration is set or the `AES256` encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES256 encryption algorithm. This does not require any action on your part.
//
// For more control over the encryption of the contents of your repository, you can use server-side encryption with AWS Key Management Service key stored in AWS Key Management Service ( AWS  ) to encrypt your images. For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html
//
type CfnRepositoryCreationTemplate_EncryptionConfigurationProperty struct {
	// The encryption type to use.
	//
	// If you use the `KMS` encryption type, the contents of the repository will be encrypted using server-side encryption with AWS Key Management Service key stored in AWS  . When you use AWS  to encrypt your data, you can either use the default AWS managed AWS  key for Amazon ECR, or specify your own AWS  key, which you already created.
	//
	// If you use the `KMS_DSSE` encryption type, the contents of the repository will be encrypted with two layers of encryption using server-side encryption with the AWS  Management Service key stored in AWS  . Similar to the `KMS` encryption type, you can either use the default AWS managed AWS  key for Amazon ECR, or specify your own AWS  key, which you've already created.
	//
	// If you use the `AES256` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES256 encryption algorithm.
	//
	// For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// If you use the `KMS` encryption type, specify the AWS  key to use for encryption.
	//
	// The alias, key ID, or full ARN of the AWS  key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed AWS  key for Amazon ECR will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

