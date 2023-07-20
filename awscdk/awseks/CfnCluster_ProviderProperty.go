package awseks


// Identifies the AWS Key Management Service ( AWS KMS ) key used to encrypt the secrets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &ProviderProperty{
//   	KeyArn: jsii.String("keyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-provider.html
//
type CfnCluster_ProviderProperty struct {
	// Amazon Resource Name (ARN) or alias of the KMS key.
	//
	// The KMS key must be symmetric and created in the same AWS Region as the cluster. If the KMS key was created in a different account, the [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) must have access to the KMS key. For more information, see [Allowing users in other accounts to use a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-provider.html#cfn-eks-cluster-provider-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

