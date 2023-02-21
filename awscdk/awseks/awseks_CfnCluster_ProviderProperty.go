package awseks


// Identifies the AWS Key Management Service ( AWS KMS ) key used to encrypt the secrets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &providerProperty{
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnCluster_ProviderProperty struct {
	// Amazon Resource Name (ARN) or alias of the KMS key.
	//
	// The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key. For more information, see [Allowing Users in Other Accounts to Use a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html) in the *AWS Key Management Service Developer Guide* .
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

