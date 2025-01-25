package awsbedrock


// Configurations for a Redshift knowledge base.
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
	// Configurations for Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-queryengineconfiguration
	//
	QueryEngineConfiguration interface{} `field:"required" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// List of configurations for available Redshift query engine storage types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-storageconfigurations
	//
	StorageConfigurations interface{} `field:"required" json:"storageConfigurations" yaml:"storageConfigurations"`
	// Configurations for generating Redshift engine queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftconfiguration.html#cfn-bedrock-knowledgebase-redshiftconfiguration-querygenerationconfiguration
	//
	QueryGenerationConfiguration interface{} `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
}

