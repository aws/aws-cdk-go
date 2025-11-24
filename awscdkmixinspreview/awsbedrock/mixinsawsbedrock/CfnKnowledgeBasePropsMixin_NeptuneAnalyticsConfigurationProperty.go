package mixinsawsbedrock


// Contains details about the storage configuration of the knowledge base in Amazon Neptune Analytics.
//
// For more information, see [Create a vector index in Amazon Neptune Analytics](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-neptune.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   neptuneAnalyticsConfigurationProperty := &NeptuneAnalyticsConfigurationProperty{
//   	FieldMapping: &NeptuneAnalyticsFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		TextField: jsii.String("textField"),
//   	},
//   	GraphArn: jsii.String("graphArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsconfiguration.html
//
type CfnKnowledgeBasePropsMixin_NeptuneAnalyticsConfigurationProperty struct {
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsconfiguration.html#cfn-bedrock-knowledgebase-neptuneanalyticsconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The Amazon Resource Name (ARN) of the Neptune Analytics vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsconfiguration.html#cfn-bedrock-knowledgebase-neptuneanalyticsconfiguration-grapharn
	//
	GraphArn *string `field:"optional" json:"graphArn" yaml:"graphArn"`
}

