package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Kms Keys for encryption ECS managed storage.
//
// Example:
//   var key key
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ManagedStorageConfiguration: &ManagedStorageConfiguration{
//   		FargateEphemeralStorageKmsKey: key,
//   	},
//   })
//
type ManagedStorageConfiguration struct {
	// KMS Key used to encrypt ECS Fargate ephemeral Storage.
	//
	// The configured KMS Key's policy will be modified to allow ECS to use the Key to encrypt the ephemeral Storage for this cluster.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-storage-encryption.html
	//
	// Default: No encryption will be applied.
	//
	FargateEphemeralStorageKmsKey awskms.IKey `field:"optional" json:"fargateEphemeralStorageKmsKey" yaml:"fargateEphemeralStorageKmsKey"`
}

