//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IResource) validateAddCorsPreflightParameters(options *CorsOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IResource) validateAddMethodParameters(httpMethod *string, options *MethodOptions) error {
	if httpMethod == nil {
		return fmt.Errorf("parameter httpMethod is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IResource) validateAddProxyParameters(options *ProxyResourceOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IResource) validateAddResourceParameters(pathPart *string, options *ResourceOptions) error {
	if pathPart == nil {
		return fmt.Errorf("parameter pathPart is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IResource) validateGetResourceParameters(pathPart *string) error {
	if pathPart == nil {
		return fmt.Errorf("parameter pathPart is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IResource) validateResourceForPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

