//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSyncTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	return nil
}

func validateNewAppSyncTargetParameters(graphqlApi awsappsync.GraphqlApi) error {
	return nil
}

