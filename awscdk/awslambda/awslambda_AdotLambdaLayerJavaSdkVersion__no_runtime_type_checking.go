//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdotLambdaLayerJavaSdkVersion) validateLayerArnParameters(scope constructs.IConstruct, architecture Architecture) error {
	return nil
}

