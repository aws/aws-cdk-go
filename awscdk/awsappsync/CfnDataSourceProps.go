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
//   	MetricsConfig: jsii.String("metricsConfig"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html
//
type CfnDataSourceProps struct {
	// Unique AWS AppSync GraphQL API identifier where this data source will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Friendly name for you to identify your AppSync data source after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	//
	// - *AWS_LAMBDA* : The data source is an AWS Lambda function.
	// - *AMAZON_DYNAMODB* : The data source is an Amazon DynamoDB table.
	// - *AMAZON_ELASTICSEARCH* : The data source is an Amazon OpenSearch Service domain.
	// - *AMAZON_EVENTBRIDGE* : The data source is an Amazon EventBridge event bus.
	// - *AMAZON_OPENSEARCH_SERVICE* : The data source is an Amazon OpenSearch Service domain.
	// - *AMAZON_BEDROCK_RUNTIME* : The data source is the Amazon Bedrock runtime.
	// - *NONE* : There is no data source. This type is used when you wish to invoke a GraphQL operation without connecting to a data source, such as performing data transformation with resolvers or triggering a subscription to be invoked from a mutation.
	// - *HTTP* : The data source is an HTTP endpoint.
	// - *RELATIONAL_DATABASE* : The data source is a relational database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Region and TableName for an Amazon DynamoDB table in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-dynamodbconfig
	//
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-elasticsearchconfig
	//
	// Deprecated: this property has been deprecated.
	ElasticsearchConfig interface{} `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// An EventBridge configuration that contains a valid ARN of an event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-eventbridgeconfig
	//
	EventBridgeConfig interface{} `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// Endpoints for an HTTP data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-httpconfig
	//
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// An ARN of a Lambda function in valid ARN format.
	//
	// This can be the ARN of a Lambda function that exists in the current account or in another account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-lambdaconfig
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Enables or disables enhanced data source metrics for specified data sources.
	//
	// Note that `MetricsConfig` won't be used unless the `dataSourceLevelMetricsBehavior` value is set to `PER_DATA_SOURCE_METRICS` . If the `dataSourceLevelMetricsBehavior` is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig` will be ignored. However, you can still set its value.
	//
	// `MetricsConfig` can be `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-metricsconfig
	//
	MetricsConfig *string `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-opensearchserviceconfig
	//
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// Relational Database configuration of the relational database data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-relationaldatabaseconfig
	//
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// The AWS Identity and Access Management service role ARN for the data source.
	//
	// The system assumes this role when accessing the data source.
	//
	// Required if `Type` is specified as `AWS_LAMBDA` , `AMAZON_DYNAMODB` , `AMAZON_ELASTICSEARCH` , `AMAZON_EVENTBRIDGE` , `AMAZON_OPENSEARCH_SERVICE` , `RELATIONAL_DATABASE` , or AMAZON_BEDROCK_RUNTIME.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html#cfn-appsync-datasource-servicerolearn
	//
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

