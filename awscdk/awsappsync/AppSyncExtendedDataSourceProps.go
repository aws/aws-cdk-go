package awsappsync


// Props used by implementations of BaseDataSource to provide configuration.
//
// Should not be used directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncExtendedDataSourceProps := &AppSyncExtendedDataSourceProps{
//   	Type: awscdk.Aws_appsync.AppSyncDataSourceType_LAMBDA,
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
type AppSyncExtendedDataSourceProps struct {
	// The type of the AppSync datasource.
	Type AppSyncDataSourceType `field:"required" json:"type" yaml:"type"`
	// Configuration for DynamoDB Datasource.
	// Default: - No config.
	//
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// Configuration for EventBridge Datasource.
	// Default: - No config.
	//
	EventBridgeConfig interface{} `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// Configuration for HTTP Datasource.
	// Default: - No config.
	//
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// Configuration for Lambda Datasource.
	// Default: - No config.
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Configuration for OpenSearch data source.
	// Default: - No config.
	//
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// Configuration for RDS Datasource.
	// Default: - No config.
	//
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
}

