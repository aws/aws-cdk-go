//go:build !no_runtime_type_checking

package awsdynamodb

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

func (t *jsiiProxy_TableGrants) validateActionsParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateFullAccessParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateMultiAccountReplicationFromParameters(sourceReplicaArn *string) error {
	if sourceReplicaArn == nil {
		return fmt.Errorf("parameter sourceReplicaArn is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateMultiAccountReplicationToParameters(destinationReplicaArn *string) error {
	if destinationReplicaArn == nil {
		return fmt.Errorf("parameter destinationReplicaArn is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateReadDataParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateReadWriteDataParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableGrants) validateWriteDataParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateTableGrants_FromTableParameters(table interfacesawsdynamodb.ITableRef) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	return nil
}

func validateNewTableGrantsParameters(props *TableGrantsProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

