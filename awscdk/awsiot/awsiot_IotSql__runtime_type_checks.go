//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiot

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v3"
)

func (i *jsiiProxy_IotSql) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateIotSql_FromStringAsVer20151008Parameters(sql *string) error {
	if sql == nil {
		return fmt.Errorf("parameter sql is required, but nil was provided")
	}

	return nil
}

func validateIotSql_FromStringAsVer20160323Parameters(sql *string) error {
	if sql == nil {
		return fmt.Errorf("parameter sql is required, but nil was provided")
	}

	return nil
}

func validateIotSql_FromStringAsVerNewestUnstableParameters(sql *string) error {
	if sql == nil {
		return fmt.Errorf("parameter sql is required, but nil was provided")
	}

	return nil
}

