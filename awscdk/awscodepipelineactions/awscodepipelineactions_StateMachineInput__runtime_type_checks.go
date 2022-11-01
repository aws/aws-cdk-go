//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateStateMachineInput_FilePathParameters(inputFile awscodepipeline.ArtifactPath) error {
	if inputFile == nil {
		return fmt.Errorf("parameter inputFile is required, but nil was provided")
	}

	return nil
}

func validateStateMachineInput_LiteralParameters(object *map[string]interface{}) error {
	if object == nil {
		return fmt.Errorf("parameter object is required, but nil was provided")
	}

	return nil
}

