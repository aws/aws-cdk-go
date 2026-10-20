//go:build !no_runtime_type_checking

package awsbedrockalphamixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (k *jsiiProxy_KnowledgeBaseTypeConfiguration) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KnowledgeBaseTypeConfiguration) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateKnowledgeBaseTypeConfiguration_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewKnowledgeBaseTypeConfigurationParameters(knowledgeBaseType awsbedrockalpha.KnowledgeBaseType) error {
	if knowledgeBaseType == nil {
		return fmt.Errorf("parameter knowledgeBaseType is required, but nil was provided")
	}

	return nil
}

