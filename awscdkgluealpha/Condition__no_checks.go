//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateCondition_CrawlerParameters(crawler interfacesawsglue.ICrawlerRef, crawlState CrawlerState, options *ConditionOptions) error {
	return nil
}

func validateCondition_JobParameters(job interfacesawsglue.IJobRef, state JobState, options *ConditionOptions) error {
	return nil
}

