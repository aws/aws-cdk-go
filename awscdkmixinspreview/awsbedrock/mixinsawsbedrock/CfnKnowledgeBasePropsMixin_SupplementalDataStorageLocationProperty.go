package mixinsawsbedrock


// Contains information about a storage location for images extracted from multimodal documents in your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   supplementalDataStorageLocationProperty := &SupplementalDataStorageLocationProperty{
//   	S3Location: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   	SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html
//
type CfnKnowledgeBasePropsMixin_SupplementalDataStorageLocationProperty struct {
	// Contains information about the Amazon S3 location for the extracted images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html#cfn-bedrock-knowledgebase-supplementaldatastoragelocation-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
	// Supplemental data storage location type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.html#cfn-bedrock-knowledgebase-supplementaldatastoragelocation-supplementaldatastoragelocationtype
	//
	SupplementalDataStorageLocationType *string `field:"optional" json:"supplementalDataStorageLocationType" yaml:"supplementalDataStorageLocationType"`
}

