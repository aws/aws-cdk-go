//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IHostedZone) validateGrantDelegationParameters(grantee awsiam.IGrantable) error {
	return nil
}

