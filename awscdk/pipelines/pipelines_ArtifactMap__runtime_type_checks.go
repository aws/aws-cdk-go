//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"
)

func (a *jsiiProxy_ArtifactMap) validateToCodePipelineParameters(x FileSet) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

