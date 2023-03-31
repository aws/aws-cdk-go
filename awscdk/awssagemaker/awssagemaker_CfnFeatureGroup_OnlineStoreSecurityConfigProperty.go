package awssagemaker


// The security configuration for `OnlineStore` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineStoreSecurityConfigProperty := &onlineStoreSecurityConfigProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFeatureGroup_OnlineStoreSecurityConfigProperty struct {
	// The AWS Key Management Service (KMS) key ARN that SageMaker Feature Store uses to encrypt the Amazon S3 objects at rest using Amazon S3 server-side encryption.
	//
	// The caller (either user or IAM role) of `CreateFeatureGroup` must have below permissions to the `OnlineStore` `KmsKeyId` :
	//
	// - `"kms:Encrypt"`
	// - `"kms:Decrypt"`
	// - `"kms:DescribeKey"`
	// - `"kms:CreateGrant"`
	// - `"kms:RetireGrant"`
	// - `"kms:ReEncryptFrom"`
	// - `"kms:ReEncryptTo"`
	// - `"kms:GenerateDataKey"`
	// - `"kms:ListAliases"`
	// - `"kms:ListGrants"`
	// - `"kms:RevokeGrant"`
	//
	// The caller (either user or IAM role) to all DataPlane operations ( `PutRecord` , `GetRecord` , `DeleteRecord` ) must have the following permissions to the `KmsKeyId` :
	//
	// - `"kms:Decrypt"`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

