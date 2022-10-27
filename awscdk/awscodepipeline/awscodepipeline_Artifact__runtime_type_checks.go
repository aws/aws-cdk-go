//go:build !no_runtime_type_checking

package awscodepipeline

import (
	"fmt"
)

func (a *jsiiProxy_Artifact) validateAtPathParameters(fileName *string) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Artifact) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Artifact) validateGetParamParameters(jsonFile *string, keyName *string) error {
	if jsonFile == nil {
		return fmt.Errorf("parameter jsonFile is required, but nil was provided")
	}

	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Artifact) validateSetMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateArtifact_ArtifactParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

