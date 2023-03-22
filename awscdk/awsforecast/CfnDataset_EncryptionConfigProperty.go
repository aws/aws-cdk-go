package awsforecast


// An AWS Key Management Service (KMS) key and an AWS Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.
//
// You can specify this optional object in the `CreateDataset` and `CreatePredictor` requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
type CfnDataset_EncryptionConfigProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.
	//
	// Passing a role across AWS accounts is not allowed. If you pass a role that isn't in your account, you get an `InvalidInputException` error.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

