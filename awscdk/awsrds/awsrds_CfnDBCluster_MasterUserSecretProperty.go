package awsrds


// Contains the secret managed by RDS in AWS Secrets Manager for the master user password.
//
// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
//
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
type CfnDBCluster_MasterUserSecretProperty struct {
	// The AWS KMS key identifier that is used to encrypt the secret.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Amazon Resource Name (ARN) of the secret.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

