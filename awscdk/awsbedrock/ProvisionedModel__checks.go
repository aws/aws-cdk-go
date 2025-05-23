//go:build !no_runtime_type_checking

package awsbedrock

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateProvisionedModel_FromProvisionedModelArnParameters(_scope constructs.Construct, _id *string, provisionedModelArn *string) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if provisionedModelArn == nil {
		return fmt.Errorf("parameter provisionedModelArn is required, but nil was provided")
	}

	return nil
}

