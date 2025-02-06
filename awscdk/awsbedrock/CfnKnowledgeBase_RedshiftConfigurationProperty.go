package awsbedrock


// Contains configurations for an Amazon Redshift database.
//
// For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftConfigurationProperty := &RedshiftConfigurationProperty{
//   	QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   			AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				DatabaseUser: jsii.String("databaseUser"),
//   				UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   			},
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   			AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   			},
//   			WorkgroupArn: jsii.String("workgroupArn"),
//   		},
//   	},
//   	StorageConfigurations: []interface{}{
//   		&RedshiftQueryEngineStorageConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   				TableNames: []*string{
//   					jsii.String("tableNames"),
//   				},
//   			},
//   			RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
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
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Columns: []interface{}{
//   						&QueryGenerationColumnProperty{
//   							Description: jsii.String("description"),
//   							Inclusion: jsii.String("inclusion"),
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Description: jsii.String("description"),
//   					Inclusion: jsii.String("inclusion"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html
//
type CfnKnowledgeBase_RedshiftConfigurationProperty struct {
	// Specifies configurations for an Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-queryengineconfiguration
	//
	QueryEngineConfiguration interface{} `field:"required" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// Specifies configurations for Amazon Redshift database storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-storageconfigurations
	//
	StorageConfigurations interface{} `field:"required" json:"storageConfigurations" yaml:"storageConfigurations"`
	// Specifies configurations for generating queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-querygenerationconfiguration
	//
	QueryGenerationConfiguration interface{} `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
}

