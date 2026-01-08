//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorStartAssessmentRun) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (i *jsiiProxy_InspectorStartAssessmentRun) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewInspectorStartAssessmentRunParameters(template interfacesawsinspector.IAssessmentTemplateRef, props *ScheduleTargetBaseProps) error {
	return nil
}

