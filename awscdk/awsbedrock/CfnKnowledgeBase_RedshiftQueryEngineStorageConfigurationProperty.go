package awsbedrock


// Contains configurations for Amazon Redshift data storage.
//
// Specify the data storage service to use in the `type` field and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftQueryEngineStorageConfigurationProperty := &RedshiftQueryEngineStorageConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   		TableNames: []*string{
//   			jsii.String("tableNames"),
//   		},
//   	},
//   	RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html
//
type CfnKnowledgeBase_RedshiftQueryEngineStorageConfigurationProperty struct {
	// The data storage service to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies configurations for storage in AWS Glue Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-awsdatacatalogconfiguration
	//
	AwsDataCatalogConfiguration interface{} `field:"optional" json:"awsDataCatalogConfiguration" yaml:"awsDataCatalogConfiguration"`
	// Specifies configurations for storage in Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-redshiftconfiguration
	//
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
}

