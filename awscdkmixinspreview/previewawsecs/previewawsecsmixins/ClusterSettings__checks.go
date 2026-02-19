//go:build !no_runtime_type_checking

package previewawsecsmixins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ClusterSettings) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClusterSettings) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateClusterSettings_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewClusterSettingsParameters(settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) error {
	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}
	for idx_cde0fb, v := range *settings {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter settings[%#v]", idx_cde0fb) }); err != nil {
			return err
		}
	}

	return nil
}

