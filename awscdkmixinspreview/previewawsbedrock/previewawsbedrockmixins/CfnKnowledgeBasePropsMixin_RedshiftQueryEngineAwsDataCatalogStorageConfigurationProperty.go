package previewawsbedrockmixins


// Contains configurations for storage in AWS Glue Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftQueryEngineAwsDataCatalogStorageConfigurationProperty := &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   	TableNames: []*string{
//   		jsii.String("tableNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineawsdatacatalogstorageconfiguration.html
//
type CfnKnowledgeBasePropsMixin_RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty struct {
	// A list of names of the tables to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineawsdatacatalogstorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineawsdatacatalogstorageconfiguration-tablenames
	//
	TableNames *[]*string `field:"optional" json:"tableNames" yaml:"tableNames"`
}

