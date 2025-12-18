//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateAmazonManagedComponent_AwsCliV2Parameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_FromAmazonManagedComponentAttributesParameters(scope constructs.Construct, id *string, attrs *AmazonManagedComponentAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_FromAmazonManagedComponentNameParameters(scope constructs.Construct, id *string, amazonManagedComponentName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if amazonManagedComponentName == nil {
		return fmt.Errorf("parameter amazonManagedComponentName is required, but nil was provided")
	}

	return nil
}

func validateAmazonManagedComponent_HelloWorldParameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_Python3Parameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_RebootParameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_StigBuildParameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateAmazonManagedComponent_UpdateOsParameters(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

