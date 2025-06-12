package awsdatasync


// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmkSecretConfigProperty := &CmkSecretConfigProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-cmksecretconfig.html
//
type CfnLocationObjectStorage_CmkSecretConfigProperty struct {
	// Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn.
	//
	// DataSync provides this key to AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-cmksecretconfig.html#cfn-datasync-locationobjectstorage-cmksecretconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Specifies the ARN for an AWS Secrets Manager secret, managed by DataSync.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-cmksecretconfig.html#cfn-datasync-locationobjectstorage-cmksecretconfig-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

