//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"
)

func validateConfigurationContent_FromFileParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateConfigurationContent_FromInlineParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateConfigurationContent_FromInlineJsonParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateConfigurationContent_FromInlineTextParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateConfigurationContent_FromInlineYamlParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

