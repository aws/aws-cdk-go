//go:build !no_runtime_type_checking

package awsschedulertargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

func (e *jsiiProxy_EcsRunFargateTask) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EcsRunFargateTask) validateBindParameters(schedule awsscheduler.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EcsRunFargateTask) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewEcsRunFargateTaskParameters(cluster awsecs.ICluster, props *FargateTaskProps) error {
	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

