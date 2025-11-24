package mixinsawsbedrock


// Specifies configurations for the storage location of the images extracted from multimodal documents in your data source.
//
// These images can be retrieved and returned to the end user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   supplementalDataStorageConfigurationProperty := &SupplementalDataStorageConfigurationProperty{
//   	SupplementalDataStorageLocations: []interface{}{
//   		&SupplementalDataStorageLocationProperty{
//   			S3Location: &S3LocationProperty{
//   				Uri: jsii.String("uri"),
//   			},
//   			SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastorageconfiguration.html
//
type CfnKnowledgeBasePropsMixin_SupplementalDataStorageConfigurationProperty struct {
	// List of supplemental data storage locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-supplementaldatastorageconfiguration.html#cfn-bedrock-knowledgebase-supplementaldatastorageconfiguration-supplementaldatastoragelocations
	//
	SupplementalDataStorageLocations interface{} `field:"optional" json:"supplementalDataStorageLocations" yaml:"supplementalDataStorageLocations"`
}

