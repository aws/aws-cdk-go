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
//   sourceSelectionCriteriaProperty := &SourceSelectionCriteriaProperty{
//   	ReplicaModifications: &ReplicaModificationsProperty{
//   		Status: jsii.String("status"),
//   	},
//   	SseKmsEncryptedObjects: &SseKmsEncryptedObjectsProperty{
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html
//
type CfnBucket_SourceSelectionCriteriaProperty struct {
	// A filter that you can specify for selection for modifications on replicas.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html#cfn-s3-bucket-sourceselectioncriteria-replicamodifications
	//
	ReplicaModifications interface{} `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html#cfn-s3-bucket-sourceselectioncriteria-ssekmsencryptedobjects
	//
	SseKmsEncryptedObjects interface{} `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

