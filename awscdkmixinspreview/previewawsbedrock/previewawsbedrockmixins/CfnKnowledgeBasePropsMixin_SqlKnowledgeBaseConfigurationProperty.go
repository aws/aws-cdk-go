package previewawsbedrockmixins


// Contains configurations for a knowledge base connected to an SQL database.
//
// Specify the SQL database type in the `type` field and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sqlKnowledgeBaseConfigurationProperty := &SqlKnowledgeBaseConfigurationProperty{
//   	RedshiftConfiguration: &RedshiftConfigurationProperty{
//   		QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   			ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   				AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   					DatabaseUser: jsii.String("databaseUser"),
//   					Type: jsii.String("type"),
//   					UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   				},
//   				ClusterIdentifier: jsii.String("clusterIdentifier"),
//   			},
//   			ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   				AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   					Type: jsii.String("type"),
//   					UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   				},
//   				WorkgroupArn: jsii.String("workgroupArn"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		QueryGenerationConfiguration: &QueryGenerationConfigurationProperty{
//   			ExecutionTimeoutSeconds: jsii.Number(123),
//   			GenerationContext: &QueryGenerationContextProperty{
//   				CuratedQueries: []interface{}{
//   					&CuratedQueryProperty{
//   						NaturalLanguage: jsii.String("naturalLanguage"),
//   						Sql: jsii.String("sql"),
//   					},
//   				},
//   				Tables: []interface{}{
//   					&QueryGenerationTableProperty{
//   						Columns: []interface{}{
//   							&QueryGenerationColumnProperty{
//   								Description: jsii.String("description"),
//   								Inclusion: jsii.String("inclusion"),
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Description: jsii.String("description"),
//   						Inclusion: jsii.String("inclusion"),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		StorageConfigurations: []interface{}{
//   			&RedshiftQueryEngineStorageConfigurationProperty{
//   				AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   					TableNames: []*string{
//   						jsii.String("tableNames"),
//   					},
//   				},
//   				RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-sqlknowledgebaseconfiguration.html
//
type CfnKnowledgeBasePropsMixin_SqlKnowledgeBaseConfigurationProperty struct {
	// Specifies configurations for a knowledge base connected to an Amazon Redshift database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-sqlknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-sqlknowledgebaseconfiguration-redshiftconfiguration
	//
	RedshiftConfiguration interface{} `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// The type of SQL database to connect to the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-sqlknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-sqlknowledgebaseconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

