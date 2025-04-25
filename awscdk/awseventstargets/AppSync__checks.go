//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

func (a *jsiiProxy_AppSync) validateBindParameters(rule awsevents.IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewAppSyncParameters(appsyncApi awsappsync.IGraphqlApi, props *AppSyncGraphQLApiProps) error {
	if appsyncApi == nil {
		return fmt.Errorf("parameter appsyncApi is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

