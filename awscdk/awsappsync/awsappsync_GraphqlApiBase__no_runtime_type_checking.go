//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GraphqlApiBase) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *DataSourceOptions) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateAddElasticsearchDataSourceParameters(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) error {
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

func (g *jsiiProxy_GraphqlApiBase) validateCreateResolverParameters(props *ExtendedResolverProps) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (g *jsiiProxy_GraphqlApiBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateGraphqlApiBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGraphqlApiBase_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewGraphqlApiBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

