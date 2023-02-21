//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateAdotLayerVersion_FromGenericLayerVersionParameters(version AdotLambdaLayerGenericVersion) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateAdotLayerVersion_FromJavaAutoInstrumentationLayerVersionParameters(version AdotLambdaLayerJavaAutoInstrumentationVersion) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateAdotLayerVersion_FromJavaScriptSdkLayerVersionParameters(version AdotLambdaLayerJavaScriptSdkVersion) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateAdotLayerVersion_FromJavaSdkLayerVersionParameters(version AdotLambdaLayerJavaSdkVersion) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateAdotLayerVersion_FromPythonSdkLayerVersionParameters(version AdotLambdaLayerPythonSdkVersion) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

