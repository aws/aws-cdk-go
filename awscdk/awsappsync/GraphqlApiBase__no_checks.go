//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GraphqlApiBase) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddElasticsearchDataSourceParameters(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddEventBridgeDataSourceParameters(id *string, eventBus awsevents.IEventBus, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddHttpDataSourceParameters(id *string, endpoint *string, options *HttpDataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddLambdaDataSourceParameters(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddNoneDataSourceParameters(id *string, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddOpenSearchDataSourceParameters(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddRdsDataSourceParameters(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddSchemaDependencyParameters(construct awscdk.CfnResource) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateCreateResolverParameters(id *string, props *ExtendedResolverProps) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGrantParameters(grantee awsiam.IGrantable, resources IamResource) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGrantMutationParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGrantSubscriptionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateGraphqlApiBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGraphqlApiBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGraphqlApiBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGraphqlApiBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

