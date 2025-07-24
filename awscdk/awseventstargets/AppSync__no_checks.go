//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSync) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewAppSyncParameters(appsyncApi awsappsync.IGraphqlApi, props *AppSyncGraphQLApiProps) error {
	return nil
}

