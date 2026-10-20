//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock"
)

func validateKnowledgeBaseReflection_OfParameters(knowledgeBaseRef interfacesawsbedrock.IKnowledgeBaseRef) error {
	if knowledgeBaseRef == nil {
		return fmt.Errorf("parameter knowledgeBaseRef is required, but nil was provided")
	}

	return nil
}

