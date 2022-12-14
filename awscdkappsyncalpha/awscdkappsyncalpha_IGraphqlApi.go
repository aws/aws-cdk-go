// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/internal"
)

// Interface for GraphQL.
// Experimental.
type IGraphqlApi interface {
	awscdk.IResource
	// add a new DynamoDB data source to this API.
	// Experimental.
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource
	// add a new elasticsearch data source to this API.
	// Deprecated: - use `addOpenSearchDataSource`.
	AddElasticsearchDataSource(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) ElasticsearchDataSource
	// add a new http data source to this API.
	// Experimental.
	AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource
	// add a new Lambda data source to this API.
	// Experimental.
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource
	// add a new dummy data source to this API.
	//
	// Useful for pipeline resolvers
	// and for backend changes that don't require a data source.
	// Experimental.
	AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource
	// Add a new OpenSearch data source to this API.
	// Experimental.
	AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) OpenSearchDataSource
	// add a new Rds data source to this API.
	// Experimental.
	AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource
	// Add schema dependency if not imported.
	// Experimental.
	AddSchemaDependency(construct awscdk.CfnResource) *bool
	// creates a new resolver for this datasource and API using the given properties.
	// Experimental.
	CreateResolver(id *string, props *ExtendedResolverProps) Resolver
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	// Experimental.
	ApiId() *string
	// the ARN of the API.
	// Experimental.
	Arn() *string
}

// The jsii proxy for IGraphqlApi
type jsiiProxy_IGraphqlApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGraphqlApi) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
	if err := i.validateAddDynamoDbDataSourceParameters(id, table, options); err != nil {
		panic(err)
	}
	var returns DynamoDbDataSource

	_jsii_.Invoke(
		i,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddElasticsearchDataSource(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) ElasticsearchDataSource {
	if err := i.validateAddElasticsearchDataSourceParameters(id, domain, options); err != nil {
		panic(err)
	}
	var returns ElasticsearchDataSource

	_jsii_.Invoke(
		i,
		"addElasticsearchDataSource",
		[]interface{}{id, domain, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
	if err := i.validateAddHttpDataSourceParameters(id, endpoint, options); err != nil {
		panic(err)
	}
	var returns HttpDataSource

	_jsii_.Invoke(
		i,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
	if err := i.validateAddLambdaDataSourceParameters(id, lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaDataSource

	_jsii_.Invoke(
		i,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
	if err := i.validateAddNoneDataSourceParameters(id, options); err != nil {
		panic(err)
	}
	var returns NoneDataSource

	_jsii_.Invoke(
		i,
		"addNoneDataSource",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) OpenSearchDataSource {
	if err := i.validateAddOpenSearchDataSourceParameters(id, domain, options); err != nil {
		panic(err)
	}
	var returns OpenSearchDataSource

	_jsii_.Invoke(
		i,
		"addOpenSearchDataSource",
		[]interface{}{id, domain, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
	if err := i.validateAddRdsDataSourceParameters(id, serverlessCluster, secretStore, options); err != nil {
		panic(err)
	}
	var returns RdsDataSource

	_jsii_.Invoke(
		i,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddSchemaDependency(construct awscdk.CfnResource) *bool {
	if err := i.validateAddSchemaDependencyParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"addSchemaDependency",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) CreateResolver(id *string, props *ExtendedResolverProps) Resolver {
	if err := i.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		i,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IGraphqlApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

