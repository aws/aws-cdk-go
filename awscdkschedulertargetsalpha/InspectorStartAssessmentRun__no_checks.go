//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorStartAssessmentRun) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewInspectorStartAssessmentRunParameters(template awsinspector.IAssessmentTemplate, props *ScheduleTargetBaseProps) error {
	return nil
}

