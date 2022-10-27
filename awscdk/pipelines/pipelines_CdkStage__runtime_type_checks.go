//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CdkStage) validateAddApplicationParameters(appStage awscdk.Stage, options *AddStageOptions) error {
	if appStage == nil {
		return fmt.Errorf("parameter appStage is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CdkStage) validateAddManualApprovalActionParameters(options *AddManualApprovalOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CdkStage) validateAddStackArtifactDeploymentParameters(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions) error {
	if stackArtifact == nil {
		return fmt.Errorf("parameter stackArtifact is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CdkStage) validateDeploysStackParameters(artifactId *string) error {
	if artifactId == nil {
		return fmt.Errorf("parameter artifactId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdkStage) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdkStage) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateCdkStage_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCdkStageParameters(scope constructs.Construct, id *string, props *CdkStageProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

