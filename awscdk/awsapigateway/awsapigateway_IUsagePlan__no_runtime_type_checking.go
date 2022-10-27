//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUsagePlan) validateAddApiKeyParameters(apiKey IApiKey, options *AddApiKeyOptions) error {
	return nil
}

