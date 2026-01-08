//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGraphqlApi) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddElasticsearchDataSourceParameters(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddEventBridgeDataSourceParameters(id *string, eventBus awsevents.IEventBus, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddHttpDataSourceParameters(id *string, endpoint *string, options *HttpDataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddLambdaDataSourceParameters(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddNoneDataSourceParameters(id *string, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddOpenSearchDataSourceParameters(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddRdsDataSourceParameters(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddRdsDataSourceV2Parameters(id *string, serverlessCluster awsrds.IDatabaseCluster, secretStore awssecretsmanager.ISecret, options *DataSourceOptions) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateAddSchemaDependencyParameters(construct awscdk.CfnResource) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateCreateResolverParameters(id *string, props *ExtendedResolverProps) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateGrantParameters(grantee awsiam.IGrantable, resources IamResource) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateGrantMutationParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateGrantSubscriptionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGraphqlApi) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

