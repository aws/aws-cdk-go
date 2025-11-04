//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (a *jsiiProxy_AppSyncTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewAppSyncTargetParameters(graphqlApi awsappsync.GraphqlApi) error {
	if graphqlApi == nil {
		return fmt.Errorf("parameter graphqlApi is required, but nil was provided")
	}

	return nil
}

