package awsforecast


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &encryptionConfigProperty{
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDataset_EncryptionConfigProperty struct {
	// `CfnDataset.EncryptionConfigProperty.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `CfnDataset.EncryptionConfigProperty.RoleArn`.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

