//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

func (i *jsiiProxy_IEventApi) validateAddChannelNamespaceParameters(id *string, options *ChannelNamespaceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *AppSyncDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddEventBridgeDataSourceParameters(id *string, eventBus awsevents.IEventBus, options *AppSyncDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if eventBus == nil {
		return fmt.Errorf("parameter eventBus is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddHttpDataSourceParameters(id *string, endpoint *string, options *AppSyncHttpDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddLambdaDataSourceParameters(id *string, lambdaFunction awslambda.IFunction, options *AppSyncDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddOpenSearchDataSourceParameters(id *string, domain awsopensearchservice.IDomain, options *AppSyncDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateAddRdsDataSourceParameters(id *string, serverlessCluster interface{}, secretStore awssecretsmanager.ISecret, options *AppSyncDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if serverlessCluster == nil {
		return fmt.Errorf("parameter serverlessCluster is required, but nil was provided")
	}
	switch serverlessCluster.(type) {
	case awsrds.IDatabaseCluster:
		// ok
	case awsrds.IServerlessCluster:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(serverlessCluster) {
			return fmt.Errorf("parameter serverlessCluster must be one of the allowed types: awsrds.IDatabaseCluster, awsrds.IServerlessCluster; received %#v (a %T)", serverlessCluster, serverlessCluster)
		}
	}

	if secretStore == nil {
		return fmt.Errorf("parameter secretStore is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantParameters(grantee awsiam.IGrantable, resources AppSyncEventResource) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if resources == nil {
		return fmt.Errorf("parameter resources is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantPublishAndSubscribeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantSubscribeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

