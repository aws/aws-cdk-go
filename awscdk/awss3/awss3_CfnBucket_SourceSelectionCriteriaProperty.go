package awss3


// A container that describes additional filters for identifying the source objects that you want to replicate.
//
// You can choose to enable or disable the replication of these objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceSelectionCriteriaProperty := &sourceSelectionCriteriaProperty{
//   	replicaModifications: &replicaModificationsProperty{
//   		status: jsii.String("status"),
//   	},
//   	sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   		status: jsii.String("status"),
//   	},
//   }
//
type CfnBucket_SourceSelectionCriteriaProperty struct {
	// A filter that you can specify for selection for modifications on replicas.
	ReplicaModifications interface{} `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.
	SseKmsEncryptedObjects interface{} `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

