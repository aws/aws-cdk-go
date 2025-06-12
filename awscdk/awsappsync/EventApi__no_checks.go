//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventApi) validateAddChannelNamespaceParameters(id *string, options *ChannelNamespaceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *AppSyncDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddEventBridgeDataSourceParameters(id *string, eventBus awsevents.IEventBus, options *AppSyncDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddHttpDataSourceParameters(id *string, endpoint *string, options *AppSyncHttpDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddLambdaDataSourceParameters(id *string, lambdaFunction awslambda.IFunction, options *AppSyncDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddOpenSearchDataSourceParameters(id *string, domain awsopensearchservice.IDomain, options *AppSyncDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateAddRdsDataSourceParameters(id *string, serverlessCluster interface{}, secretStore awssecretsmanager.ISecret, options *AppSyncDataSourceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGrantParameters(grantee awsiam.IGrantable, resources AppSyncEventResource) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGrantPublishAndSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApi) validateGrantSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEventApi_FromEventApiAttributesParameters(scope constructs.Construct, id *string, attrs *EventApiAttributes) error {
	return nil
}

func validateEventApi_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventApi_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEventApi_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEventApiParameters(scope constructs.Construct, id *string, props *EventApiProps) error {
	return nil
}

