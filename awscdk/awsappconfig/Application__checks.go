//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_Application) validateAddEnvironmentParameters(id *string, options *EnvironmentOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateAddExistingEnvironmentParameters(environment IEnvironment) error {
	if environment == nil {
		return fmt.Errorf("parameter environment is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Application) validateAddExtensionParameters(extension IExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Application) validateAddHostedConfigurationParameters(id *string, options *HostedConfigurationOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateAddSourcedConfigurationParameters(id *string, options *SourcedConfigurationOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Application) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	if actionPoint == "" {
		return fmt.Errorf("parameter actionPoint is required, but nil was provided")
	}

	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Application) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateApplication_AddAgentToEcsParameters(taskDef awsecs.TaskDefinition) error {
	if taskDef == nil {
		return fmt.Errorf("parameter taskDef is required, but nil was provided")
	}

	return nil
}

func validateApplication_FromApplicationArnParameters(scope constructs.Construct, id *string, applicationArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if applicationArn == nil {
		return fmt.Errorf("parameter applicationArn is required, but nil was provided")
	}

	return nil
}

func validateApplication_FromApplicationIdParameters(scope constructs.Construct, id *string, applicationId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if applicationId == nil {
		return fmt.Errorf("parameter applicationId is required, but nil was provided")
	}

	return nil
}

func validateApplication_GetLambdaLayerVersionArnParameters(region *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func validateApplication_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateApplication_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateApplication_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Application) validateSetExtensibleParameters(val ExtensibleBase) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationParameters(scope constructs.Construct, id *string, props *ApplicationProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

