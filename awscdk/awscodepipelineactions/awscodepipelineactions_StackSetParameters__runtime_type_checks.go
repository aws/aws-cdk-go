//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateStackSetParameters_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath) error {
	if artifactPath == nil {
		return fmt.Errorf("parameter artifactPath is required, but nil was provided")
	}

	return nil
}

func validateStackSetParameters_FromLiteralParameters(parameters *map[string]*string) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}

	return nil
}

