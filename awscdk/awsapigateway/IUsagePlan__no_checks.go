//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUsagePlan) validateAddApiKeyParameters(apiKey IApiKeyRef, options *AddApiKeyOptions) error {
	return nil
}

func (i *jsiiProxy_IUsagePlan) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

