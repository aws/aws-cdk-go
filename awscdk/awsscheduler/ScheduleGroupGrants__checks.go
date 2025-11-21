//go:build !no_runtime_type_checking

package awsscheduler

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsscheduler"
)

func (s *jsiiProxy_ScheduleGroupGrants) validateDeleteSchedulesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduleGroupGrants) validateReadSchedulesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduleGroupGrants) validateWriteSchedulesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateScheduleGroupGrants_FromScheduleGroupParameters(resource interfacesawsscheduler.IScheduleGroupRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

