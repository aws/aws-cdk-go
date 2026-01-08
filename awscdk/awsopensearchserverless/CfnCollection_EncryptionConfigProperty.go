package awsopensearchserverless


// The configuration to encrypt the collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	AwsOwnedKey: jsii.Boolean(false),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html
//
type CfnCollection_EncryptionConfigProperty struct {
	// The configuration to encrypt the collection with AWS owned key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html#cfn-opensearchserverless-collection-encryptionconfig-awsownedkey
	//
	AwsOwnedKey interface{} `field:"optional" json:"awsOwnedKey" yaml:"awsOwnedKey"`
	// The ARN of the KMS key to encrypt the collection with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-encryptionconfig.html#cfn-opensearchserverless-collection-encryptionconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

