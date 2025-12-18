package awsbedrock


// Contains information about a storage location for multimedia content (images, audio, and video) extracted from multimodal documents in your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   supplementalDataStorageLocationProperty := &SupplementalDataStorageLocationProperty{
//   	SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//
//   	// the properties below are optional
//   	S3Location: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html
//
type CfnKnowledgeBase_SupplementalDataStorageLocationProperty struct {
	// Supplemental data storage location type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html#cfn-bedrock-knowledgebase-supplementaldatastoragelocation-supplementaldatastoragelocationtype
	//
	SupplementalDataStorageLocationType *string `field:"required" json:"supplementalDataStorageLocationType" yaml:"supplementalDataStorageLocationType"`
	// Contains information about the Amazon S3 location for the extracted multimedia content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html#cfn-bedrock-knowledgebase-supplementaldatastoragelocation-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

