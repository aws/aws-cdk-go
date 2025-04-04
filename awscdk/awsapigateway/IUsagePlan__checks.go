//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IUsagePlan) validateAddApiKeyParameters(apiKey IApiKey, options *AddApiKeyOptions) error {
	if apiKey == nil {
		return fmt.Errorf("parameter apiKey is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

