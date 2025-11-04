//go:build no_runtime_type_checking

package awscdkpipessourcesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoDBSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

func validateDynamoDBSource_IsSourceWithDeadLetterTargetParameters(source awscdkpipesalpha.ISource) error {
	return nil
}

func validateNewDynamoDBSourceParameters(table awsdynamodb.ITableV2, parameters *DynamoDBSourceParameters) error {
	return nil
}

