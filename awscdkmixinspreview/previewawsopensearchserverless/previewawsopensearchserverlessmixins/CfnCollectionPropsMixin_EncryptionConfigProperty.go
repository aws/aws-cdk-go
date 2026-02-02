package previewawsopensearchserverlessmixins


// Encryption settings for the collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	AwsOwnedKey: jsii.Boolean(false),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html
//
type CfnCollectionPropsMixin_EncryptionConfigProperty struct {
	// Indicates whether to use an AWS owned key for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html#cfn-opensearchserverless-collection-encryptionconfig-awsownedkey
	//
	AwsOwnedKey interface{} `field:"optional" json:"awsOwnedKey" yaml:"awsOwnedKey"`
	// Key Management Service key used to encrypt the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html#cfn-opensearchserverless-collection-encryptionconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

