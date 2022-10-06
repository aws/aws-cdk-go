//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

func (d *jsiiProxy_DeployCdkStackAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeployCdkStackAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDeployCdkStackAction_FromStackArtifactParameters(scope constructs.Construct, artifact cxapi.CloudFormationStackArtifact, options *CdkStackActionFromArtifactOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if artifact == nil {
		return fmt.Errorf("parameter artifact is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewDeployCdkStackActionParameters(props *DeployCdkStackActionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

