//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateCodePipelineFileSet_FromArtifactParameters(artifact awscodepipeline.Artifact) error {
	if artifact == nil {
		return fmt.Errorf("parameter artifact is required, but nil was provided")
	}

	return nil
}

