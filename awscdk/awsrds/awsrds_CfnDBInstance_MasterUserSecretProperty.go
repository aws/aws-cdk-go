package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   masterUserSecretProperty := &masterUserSecretProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnDBInstance_MasterUserSecretProperty struct {
	// `CfnDBInstance.MasterUserSecretProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `CfnDBInstance.MasterUserSecretProperty.SecretArn`.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

