//go:build !no_runtime_type_checking

package awscdkclilibalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AwsCdkCli) validateBootstrapParameters(options *BootstrapOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateDeployParameters(options *DeployOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateDestroyParameters(options *DestroyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateListParameters(options *ListOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateSynthParameters(options *SynthOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAwsCdkCli_FromCdkAppDirectoryParameters(props *CdkAppDirectoryProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAwsCdkCli_FromCloudAssemblyDirectoryProducerParameters(producer ICloudAssemblyDirectoryProducer) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	return nil
}

