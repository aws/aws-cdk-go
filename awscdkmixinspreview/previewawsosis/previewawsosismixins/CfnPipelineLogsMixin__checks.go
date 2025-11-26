//go:build !no_runtime_type_checking

package previewawsosismixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnPipelineLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPipelineLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnPipelineLogsMixin_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCfnPipelineLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	if logType == nil {
		return fmt.Errorf("parameter logType is required, but nil was provided")
	}

	if logDelivery == nil {
		return fmt.Errorf("parameter logDelivery is required, but nil was provided")
	}

	return nil
}

