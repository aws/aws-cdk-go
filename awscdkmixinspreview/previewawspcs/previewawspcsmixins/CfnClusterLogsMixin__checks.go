//go:build !no_runtime_type_checking

package previewawspcsmixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnClusterLogsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnClusterLogsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnClusterLogsMixin_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCfnClusterLogsMixinParameters(logType *string, logDelivery previewawslogs.ILogsDelivery) error {
	if logType == nil {
		return fmt.Errorf("parameter logType is required, but nil was provided")
	}

	if logDelivery == nil {
		return fmt.Errorf("parameter logDelivery is required, but nil was provided")
	}

	return nil
}

