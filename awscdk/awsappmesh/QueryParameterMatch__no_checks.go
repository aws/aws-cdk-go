//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueryParameterMatch) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateQueryParameterMatch_ValueIsParameters(queryParameterName *string, queryParameterValue *string) error {
	return nil
}

