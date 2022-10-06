//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateNewStackOutputParameters(artifactFile awscodepipeline.ArtifactPath, outputName *string) error {
	if artifactFile == nil {
		return fmt.Errorf("parameter artifactFile is required, but nil was provided")
	}

	if outputName == nil {
		return fmt.Errorf("parameter outputName is required, but nil was provided")
	}

	return nil
}

