//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateAction_CrawlerParameters(crawler interfacesawsglue.ICrawlerRef, options *CrawlerActionOptions) error {
	return nil
}

func validateAction_JobParameters(job interfacesawsglue.IJobRef, options *JobActionOptions) error {
	return nil
}

