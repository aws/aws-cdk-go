//go:build !no_runtime_type_checking

package regioninfo

import (
	"fmt"
)

func (r *jsiiProxy_RegionInfo) validateAdotLambdaLayerArnParameters(type_ *string, version *string, architecture *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RegionInfo) validateCloudwatchLambdaInsightsArnParameters(insightsVersion *string) error {
	if insightsVersion == nil {
		return fmt.Errorf("parameter insightsVersion is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RegionInfo) validateServicePrincipalParameters(service *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateRegionInfo_GetParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateRegionInfo_LimitedRegionMapParameters(factName *string, partitions *[]*string) error {
	if factName == nil {
		return fmt.Errorf("parameter factName is required, but nil was provided")
	}

	if partitions == nil {
		return fmt.Errorf("parameter partitions is required, but nil was provided")
	}

	return nil
}

func validateRegionInfo_RegionMapParameters(factName *string) error {
	if factName == nil {
		return fmt.Errorf("parameter factName is required, but nil was provided")
	}

	return nil
}

