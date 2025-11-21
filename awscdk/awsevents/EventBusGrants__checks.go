//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

func (e *jsiiProxy_EventBusGrants) validateAllPutEventsParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateEventBusGrants_FromEventBusParameters(resource interfacesawsevents.IEventBusRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

