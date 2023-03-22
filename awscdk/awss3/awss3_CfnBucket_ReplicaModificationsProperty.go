package awss3


// A filter that you can specify for selection for modifications on replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaModificationsProperty := &replicaModificationsProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_ReplicaModificationsProperty struct {
	// Specifies whether Amazon S3 replicates modifications on replicas.
	//
	// *Allowed values* : `Enabled` | `Disabled`.
	Status *string `field:"required" json:"status" yaml:"status"`
}

