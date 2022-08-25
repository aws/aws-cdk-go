package awsappsync


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &cfnDataSourceProps{
//   	apiId: jsii.String("apiId"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
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
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
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
	// Required if `Type` is specified as `AWS_LAMBDA` , `AMAZON_DYNAMODB` , `AMAZON_ELASTICSEARCH` , or `AMAZON_OPENSEARCH_SERVICE` .
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

