package awsbedrock


// Contains the names of the fields to which to map information about the vector store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   neptuneAnalyticsFieldMappingProperty := &NeptuneAnalyticsFieldMappingProperty{
//   	MetadataField: jsii.String("metadataField"),
//   	TextField: jsii.String("textField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsfieldmapping.html
//
type CfnKnowledgeBase_NeptuneAnalyticsFieldMappingProperty struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsfieldmapping.html#cfn-bedrock-knowledgebase-neptuneanalyticsfieldmapping-metadatafield
	//
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-neptuneanalyticsfieldmapping.html#cfn-bedrock-knowledgebase-neptuneanalyticsfieldmapping-textfield
	//
	TextField *string `field:"required" json:"textField" yaml:"textField"`
}

