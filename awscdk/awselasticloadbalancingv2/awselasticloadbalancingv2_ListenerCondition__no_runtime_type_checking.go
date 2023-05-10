//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func validateListenerCondition_HostHeadersParameters(values *[]*string) error {
	return nil
}

func validateListenerCondition_HttpHeaderParameters(name *string, values *[]*string) error {
	return nil
}

func validateListenerCondition_HttpRequestMethodsParameters(values *[]*string) error {
	return nil
}

func validateListenerCondition_PathPatternsParameters(values *[]*string) error {
	return nil
}

func validateListenerCondition_QueryStringsParameters(values *[]*QueryStringCondition) error {
	return nil
}

func validateListenerCondition_SourceIpsParameters(values *[]*string) error {
	return nil
}

