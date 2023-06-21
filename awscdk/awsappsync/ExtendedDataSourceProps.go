package awsappsync


// props used by implementations of BaseDataSource to provide configuration.
//
// Should not be used directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extendedDataSourceProps := &ExtendedDataSourceProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DynamoDbConfig: &DynamoDBConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		DeltaSyncConfig: &DeltaSyncConfigProperty{
//   			BaseTableTtl: jsii.String("baseTableTtl"),
//   			DeltaSyncTableName: jsii.String("deltaSyncTableName"),
//   			DeltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   		},
//   		UseCallerCredentials: jsii.Boolean(false),
//   		Versioned: jsii.Boolean(false),
//   	},
//   	ElasticsearchConfig: &ElasticsearchConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	EventBridgeConfig: &EventBridgeConfigProperty{
//   		EventBusArn: jsii.String("eventBusArn"),
//   	},
//   	HttpConfig: &HttpConfigProperty{
//   		Endpoint: jsii.String("endpoint"),
//
//   		// the properties below are optional
//   		AuthorizationConfig: &AuthorizationConfigProperty{
//   			AuthorizationType: jsii.String("authorizationType"),
//
//   			// the properties below are optional
//   			AwsIamConfig: &AwsIamConfigProperty{
//   				SigningRegion: jsii.String("signingRegion"),
//   				SigningServiceName: jsii.String("signingServiceName"),
//   			},
//   		},
//   	},
//   	LambdaConfig: &LambdaConfigProperty{
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	OpenSearchServiceConfig: &OpenSearchServiceConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	RelationalDatabaseConfig: &RelationalDatabaseConfigProperty{
//   		RelationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   		// the properties below are optional
//   		RdsHttpEndpointConfig: &RdsHttpEndpointConfigProperty{
//   			AwsRegion: jsii.String("awsRegion"),
//   			AwsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   			DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   			// the properties below are optional
//   			DatabaseName: jsii.String("databaseName"),
//   			Schema: jsii.String("schema"),
//   		},
//   	},
//   }
//
type ExtendedDataSourceProps struct {
	// the type of the AppSync datasource.
	Type *string `field:"required" json:"type" yaml:"type"`
	// configuration for DynamoDB Datasource.
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// configuration for Elasticsearch data source.
	// Deprecated: - use `openSearchConfig`.
	ElasticsearchConfig interface{} `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// configuration for EventBridge Datasource.
	EventBridgeConfig interface{} `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// configuration for HTTP Datasource.
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// configuration for Lambda Datasource.
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// configuration for OpenSearch data source.
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// configuration for RDS Datasource.
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
}

