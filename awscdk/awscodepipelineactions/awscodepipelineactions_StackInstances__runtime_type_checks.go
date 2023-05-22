//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

func validateStackInstances_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath, regions *[]*string) error {
	if artifactPath == nil {
		return fmt.Errorf("parameter artifactPath is required, but nil was provided")
	}

	if regions == nil {
		return fmt.Errorf("parameter regions is required, but nil was provided")
	}

	return nil
}

func validateStackInstances_InAccountsParameters(accounts *[]*string, regions *[]*string) error {
	if accounts == nil {
		return fmt.Errorf("parameter accounts is required, but nil was provided")
	}

	if regions == nil {
		return fmt.Errorf("parameter regions is required, but nil was provided")
	}

	return nil
}

func validateStackInstances_InOrganizationalUnitsParameters(ous *[]*string, regions *[]*string) error {
	if ous == nil {
		return fmt.Errorf("parameter ous is required, but nil was provided")
	}

	if regions == nil {
		return fmt.Errorf("parameter regions is required, but nil was provided")
	}

	return nil
}

