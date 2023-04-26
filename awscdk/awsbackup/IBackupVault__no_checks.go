//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBackupVault) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

