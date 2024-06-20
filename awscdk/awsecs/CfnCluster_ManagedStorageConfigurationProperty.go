package awsecs


// The managed storage configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedStorageConfigurationProperty := &ManagedStorageConfigurationProperty{
//   	FargateEphemeralStorageKmsKeyId: jsii.String("fargateEphemeralStorageKmsKeyId"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-managedstorageconfiguration.html
//
type CfnCluster_ManagedStorageConfigurationProperty struct {
	// Specify the AWS Key Management Service key ID for the Fargate ephemeral storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-managedstorageconfiguration.html#cfn-ecs-cluster-managedstorageconfiguration-fargateephemeralstoragekmskeyid
	//
	FargateEphemeralStorageKmsKeyId *string `field:"optional" json:"fargateEphemeralStorageKmsKeyId" yaml:"fargateEphemeralStorageKmsKeyId"`
	// Specify a AWS Key Management Service key ID to encrypt the managed storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-managedstorageconfiguration.html#cfn-ecs-cluster-managedstorageconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

