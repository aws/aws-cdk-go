//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateStackSetTemplate_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath) error {
	if artifactPath == nil {
		return fmt.Errorf("parameter artifactPath is required, but nil was provided")
	}

	return nil
}

