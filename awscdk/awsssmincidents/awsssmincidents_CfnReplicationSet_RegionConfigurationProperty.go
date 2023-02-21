package awsssmincidents


// The `RegionConfiguration` property specifies the Region and KMS key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionConfigurationProperty := &regionConfigurationProperty{
//   	sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   }
//
type CfnReplicationSet_RegionConfigurationProperty struct {
	// The KMS key ID to use to encrypt your replication set.
	SseKmsKeyId *string `field:"required" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
}

