//go:build !no_runtime_type_checking

package awsbackup

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateBackupResource_FromArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromConstructParameters(construct constructs.Construct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromDynamoDbTableParameters(table interfacesawsdynamodb.ITableRef) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromEc2InstanceParameters(instance interfacesawsec2.IInstanceRef) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromEfsFileSystemParameters(fileSystem interfacesawsefs.IFileSystemRef) error {
	if fileSystem == nil {
		return fmt.Errorf("parameter fileSystem is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromRdsDatabaseClusterParameters(cluster interfacesawsrds.IDBClusterRef) error {
	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromRdsDatabaseInstanceParameters(instance interfacesawsrds.IDBInstanceRef) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromRdsServerlessClusterParameters(cluster interfacesawsrds.IDBClusterRef) error {
	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	return nil
}

func validateBackupResource_FromTagParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewBackupResourceParameters(tagCondition *TagCondition) error {
	if err := _jsii_.ValidateStruct(tagCondition, func() string { return "parameter tagCondition" }); err != nil {
		return err
	}

	return nil
}

