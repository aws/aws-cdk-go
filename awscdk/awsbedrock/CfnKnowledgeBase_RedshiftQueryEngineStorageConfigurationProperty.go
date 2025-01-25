package awsbedrock


// Configurations for available Redshift query engine storage types.
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
	// Redshift query engine storage type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configurations for Redshift query engine AWS Data Catalog backed storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-awsdatacatalogconfiguration
	//
	AwsDataCatalogConfiguration interface{} `field:"optional" json:"awsDataCatalogConfiguration" yaml:"awsDataCatalogConfiguration"`
	// Configurations for Redshift query engine Redshift backed storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration-redshiftconfiguration
	//
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
}

