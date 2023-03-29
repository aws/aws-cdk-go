package awsappsync


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	ApiId: jsii.String("apiId"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   }
//
type CfnDataSourceProps struct {
	// Unique AWS AppSync GraphQL API identifier where this data source will be created.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Friendly name for you to identify your AppSync data source after creation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	//
	// - *AWS_LAMBDA* : The data source is an AWS Lambda function.
	// - *AMAZON_DYNAMODB* : The data source is an Amazon DynamoDB table.
	// - *AMAZON_ELASTICSEARCH* : The data source is an Amazon OpenSearch Service domain.
	// - *AMAZON_EVENTBRIDGE* : The data source is an Amazon EventBridge event bus.
	// - *AMAZON_OPENSEARCH_SERVICE* : The data source is an Amazon OpenSearch Service domain.
	// - *NONE* : There is no data source. This type is used when you wish to invoke a GraphQL operation without connecting to a data source, such as performing data transformation with resolvers or triggering a subscription to be invoked from a mutation.
	// - *HTTP* : The data source is an HTTP endpoint.
	// - *RELATIONAL_DATABASE* : The data source is a relational database.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Region and TableName for an Amazon DynamoDB table in your account.
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	//
	// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service . This property is deprecated. For new data sources, use *OpenSearchServiceConfig* to specify an OpenSearch Service data source.
	ElasticsearchConfig interface{} `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// An EventBridge configuration that contains a valid ARN of an event bus.
	EventBridgeConfig interface{} `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// Endpoints for an HTTP data source.
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// An ARN of a Lambda function in valid ARN format.
	//
	// This can be the ARN of a Lambda function that exists in the current account or in another account.
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// Relational Database configuration of the relational database data source.
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// The AWS Identity and Access Management service role ARN for the data source.
	//
	// The system assumes this role when accessing the data source.
	//
	// Required if `Type` is specified as `AWS_LAMBDA` , `AMAZON_DYNAMODB` , `AMAZON_ELASTICSEARCH` , `AMAZON_EVENTBRIDGE` , or `AMAZON_OPENSEARCH_SERVICE` .
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

