//go:build no_runtime_type_checking

package regioninfo

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegionInfo) validateAdotLambdaLayerArnParameters(type_ *string, version *string, architecture *string) error {
	return nil
}

func (r *jsiiProxy_RegionInfo) validateCloudwatchLambdaInsightsArnParameters(insightsVersion *string) error {
	return nil
}

func (r *jsiiProxy_RegionInfo) validateServicePrincipalParameters(service *string) error {
	return nil
}

func validateRegionInfo_GetParameters(name *string) error {
	return nil
}

func validateRegionInfo_LimitedRegionMapParameters(factName *string, partitions *[]*string) error {
	return nil
}

func validateRegionInfo_RegionMapParameters(factName *string) error {
	return nil
}

