//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func validateBackupResource_FromArnParameters(arn *string) error {
	return nil
}

func validateBackupResource_FromConstructParameters(construct constructs.Construct) error {
	return nil
}

func validateBackupResource_FromDynamoDbTableParameters(table awsdynamodb.ITable) error {
	return nil
}

func validateBackupResource_FromEc2InstanceParameters(instance awsec2.IInstance) error {
	return nil
}

func validateBackupResource_FromEfsFileSystemParameters(fileSystem awsefs.IFileSystem) error {
	return nil
}

func validateBackupResource_FromRdsDatabaseClusterParameters(cluster awsrds.IDatabaseCluster) error {
	return nil
}

func validateBackupResource_FromRdsDatabaseInstanceParameters(instance awsrds.IDatabaseInstance) error {
	return nil
}

func validateBackupResource_FromRdsServerlessClusterParameters(cluster awsrds.IServerlessCluster) error {
	return nil
}

func validateBackupResource_FromTagParameters(key *string, value *string) error {
	return nil
}

func validateNewBackupResourceParameters(tagCondition *TagCondition) error {
	return nil
}

