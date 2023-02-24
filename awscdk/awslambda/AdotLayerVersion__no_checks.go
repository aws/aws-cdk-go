//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateAdotLayerVersion_FromGenericLayerVersionParameters(version AdotLambdaLayerGenericVersion) error {
	return nil
}

func validateAdotLayerVersion_FromJavaAutoInstrumentationLayerVersionParameters(version AdotLambdaLayerJavaAutoInstrumentationVersion) error {
	return nil
}

func validateAdotLayerVersion_FromJavaScriptSdkLayerVersionParameters(version AdotLambdaLayerJavaScriptSdkVersion) error {
	return nil
}

func validateAdotLayerVersion_FromJavaSdkLayerVersionParameters(version AdotLambdaLayerJavaSdkVersion) error {
	return nil
}

func validateAdotLayerVersion_FromPythonSdkLayerVersionParameters(version AdotLambdaLayerPythonSdkVersion) error {
	return nil
}

