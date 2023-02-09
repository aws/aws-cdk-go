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
//   extendedDataSourceProps := &extendedDataSourceProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	dynamoDbConfig: &dynamoDBConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		deltaSyncConfig: &deltaSyncConfigProperty{
//   			baseTableTtl: jsii.String("baseTableTtl"),
//   			deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   			deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   		},
//   		useCallerCredentials: jsii.Boolean(false),
//   		versioned: jsii.Boolean(false),
//   	},
//   	elasticsearchConfig: &elasticsearchConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	httpConfig: &httpConfigProperty{
//   		endpoint: jsii.String("endpoint"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			authorizationType: jsii.String("authorizationType"),
//
//   			// the properties below are optional
//   			awsIamConfig: &awsIamConfigProperty{
//   				signingRegion: jsii.String("signingRegion"),
//   				signingServiceName: jsii.String("signingServiceName"),
//   			},
//   		},
//   	},
//   	lambdaConfig: &lambdaConfigProperty{
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	openSearchServiceConfig: &openSearchServiceConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	relationalDatabaseConfig: &relationalDatabaseConfigProperty{
//   		relationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   		// the properties below are optional
//   		rdsHttpEndpointConfig: &rdsHttpEndpointConfigProperty{
//   			awsRegion: jsii.String("awsRegion"),
//   			awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   			dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   			// the properties below are optional
//   			databaseName: jsii.String("databaseName"),
//   			schema: jsii.String("schema"),
//   		},
//   	},
//   }
//
// Experimental.
type ExtendedDataSourceProps struct {
	// the type of the AppSync datasource.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// configuration for DynamoDB Datasource.
	// Experimental.
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// configuration for Elasticsearch data source.
	// Deprecated: - use `openSearchConfig`.
	ElasticsearchConfig interface{} `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// configuration for HTTP Datasource.
	// Experimental.
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// configuration for Lambda Datasource.
	// Experimental.
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// configuration for OpenSearch data source.
	// Experimental.
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// configuration for RDS Datasource.
	// Experimental.
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
}

