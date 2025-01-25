//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IWorkflow) validateAddCustomScheduledTriggerParameters(id *string, options *CustomScheduledTriggerOptions) error {
	return nil
}

func (i *jsiiProxy_IWorkflow) validateAddDailyScheduledTriggerParameters(id *string, options *DailyScheduleTriggerOptions) error {
	return nil
}

func (i *jsiiProxy_IWorkflow) validateAddOnDemandTriggerParameters(id *string, options *OnDemandTriggerOptions) error {
	return nil
}

func (i *jsiiProxy_IWorkflow) validateAddWeeklyScheduledTriggerParameters(id *string, options *WeeklyScheduleTriggerOptions) error {
	return nil
}

