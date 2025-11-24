package mixinsawsbedrock


// Contains details about the Managed Cluster configuration of the knowledge base in Amazon OpenSearch Service.
//
// For more information, see [Create a vector index in OpenSearch Managed Cluster](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-osm.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openSearchManagedClusterConfigurationProperty := &OpenSearchManagedClusterConfigurationProperty{
//   	DomainArn: jsii.String("domainArn"),
//   	DomainEndpoint: jsii.String("domainEndpoint"),
//   	FieldMapping: &OpenSearchManagedClusterFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		TextField: jsii.String("textField"),
//   		VectorField: jsii.String("vectorField"),
//   	},
//   	VectorIndexName: jsii.String("vectorIndexName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.html
//
type CfnKnowledgeBasePropsMixin_OpenSearchManagedClusterConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the OpenSearch domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterconfiguration-domainarn
	//
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// The endpoint URL the OpenSearch domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterconfiguration-domainendpoint
	//
	DomainEndpoint *string `field:"optional" json:"domainEndpoint" yaml:"domainEndpoint"`
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterconfiguration-vectorindexname
	//
	VectorIndexName *string `field:"optional" json:"vectorIndexName" yaml:"vectorIndexName"`
}

