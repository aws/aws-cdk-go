//go:build !no_runtime_type_checking

package awssqs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssqs"
)

func (q *jsiiProxy_QueueGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.EncryptedPermissionsOptions) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QueueGrants) validateConsumeMessagesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QueueGrants) validatePurgeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QueueGrants) validateSendMessagesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateQueueGrants_FromQueueParameters(resource interfacesawssqs.IQueueRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

