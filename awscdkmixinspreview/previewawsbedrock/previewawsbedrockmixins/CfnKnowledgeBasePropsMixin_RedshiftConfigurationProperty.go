package previewawsbedrockmixins


// Contains configurations for an Amazon Redshift database.
//
// For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftConfigurationProperty := &RedshiftConfigurationProperty{
//   	QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   		ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   			AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   				DatabaseUser: jsii.String("databaseUser"),
//   				Type: jsii.String("type"),
//   				UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   			},
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   			AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   				Type: jsii.String("type"),
//   				UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   			},
//   			WorkgroupArn: jsii.String("workgroupArn"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	QueryGenerationConfiguration: &QueryGenerationConfigurationProperty{
//   		ExecutionTimeoutSeconds: jsii.Number(123),
//   		GenerationContext: &QueryGenerationContextProperty{
//   			CuratedQueries: []interface{}{
//   				&CuratedQueryProperty{
//   					NaturalLanguage: jsii.String("naturalLanguage"),
//   					Sql: jsii.String("sql"),
//   				},
//   			},
//   			Tables: []interface{}{
//   				&QueryGenerationTableProperty{
//   					Columns: []interface{}{
//   						&QueryGenerationColumnProperty{
//   							Description: jsii.String("description"),
//   							Inclusion: jsii.String("inclusion"),
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Description: jsii.String("description"),
//   					Inclusion: jsii.String("inclusion"),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	StorageConfigurations: []interface{}{
//   		&RedshiftQueryEngineStorageConfigurationProperty{
//   			AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   				TableNames: []*string{
//   					jsii.String("tableNames"),
//   				},
//   			},
//   			RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html
//
type CfnKnowledgeBasePropsMixin_RedshiftConfigurationProperty struct {
	// Specifies configurations for an Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-queryengineconfiguration
	//
	QueryEngineConfiguration interface{} `field:"optional" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// Specifies configurations for generating queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-querygenerationconfiguration
	//
	QueryGenerationConfiguration interface{} `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
	// Specifies configurations for Amazon Redshift database storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-storageconfigurations
	//
	StorageConfigurations interface{} `field:"optional" json:"storageConfigurations" yaml:"storageConfigurations"`
}

