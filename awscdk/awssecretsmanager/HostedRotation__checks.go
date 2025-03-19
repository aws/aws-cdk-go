//go:build !no_runtime_type_checking

package awssecretsmanager

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HostedRotation) validateBindParameters(secret ISecret, scope constructs.Construct) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateHostedRotation_MariaDbMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_MariaDbSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_MongoDbMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_MongoDbSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_MysqlMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_MysqlSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_OracleMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_OracleSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_PostgreSqlMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_PostgreSqlSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_RedshiftMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_RedshiftSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_SqlServerMultiUserParameters(options *MultiUserHostedRotationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHostedRotation_SqlServerSingleUserParameters(options *SingleUserHostedRotationOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

