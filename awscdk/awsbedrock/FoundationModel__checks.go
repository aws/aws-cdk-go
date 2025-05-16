//go:build !no_runtime_type_checking

package awsbedrock

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateFoundationModel_FromFoundationModelIdParameters(scope constructs.Construct, _id *string, foundationModelId FoundationModelIdentifier) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if foundationModelId == nil {
		return fmt.Errorf("parameter foundationModelId is required, but nil was provided")
	}

	return nil
}

