package awscodepipeline


// Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
//
// `EncryptionKey` is a property of the [ArtifactStore](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionKeyProperty := &encryptionKeyProperty{
//   	id: jsii.String("id"),
//   	type: jsii.String("type"),
//   }
//
type CfnPipeline_EncryptionKeyProperty struct {
	// The ID used to identify the key.
	//
	// For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.
	//
	// > Aliases are recognized only in the account that created the AWS KMS key. For cross-account actions, you can only use the key ID or key ARN to identify the key.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of encryption key, such as an AWS KMS key.
	//
	// When creating or updating a pipeline, the value must be set to 'KMS'.
	Type *string `field:"required" json:"type" yaml:"type"`
}

