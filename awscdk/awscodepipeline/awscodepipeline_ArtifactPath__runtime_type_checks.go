//go:build !no_runtime_type_checking

package awscodepipeline

import (
	"fmt"
)

func validateArtifactPath_ArtifactPathParameters(artifactName *string, fileName *string) error {
	if artifactName == nil {
		return fmt.Errorf("parameter artifactName is required, but nil was provided")
	}

	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	return nil
}

func validateNewArtifactPathParameters(artifact Artifact, fileName *string) error {
	if artifact == nil {
		return fmt.Errorf("parameter artifact is required, but nil was provided")
	}

	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	return nil
}

