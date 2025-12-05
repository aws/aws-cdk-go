//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnCapacityProvider) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnCapacityProvider) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_ArnForCapacityProviderParameters(resource interfacesawslambda.ICapacityProviderRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_FromCapacityProviderArnParameters(scope constructs.Construct, id *string, arn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_FromCapacityProviderNameParameters(scope constructs.Construct, id *string, capacityProviderName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if capacityProviderName == nil {
		return fmt.Errorf("parameter capacityProviderName is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_IsCfnCapacityProviderParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_IsCfnResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnCapacityProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnCapacityProvider) validateSetCapacityProviderScalingConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnCapacityProvider_CapacityProviderScalingConfigProperty:
		val := val.(*CfnCapacityProvider_CapacityProviderScalingConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnCapacityProvider_CapacityProviderScalingConfigProperty:
		val_ := val.(CfnCapacityProvider_CapacityProviderScalingConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnCapacityProvider_CapacityProviderScalingConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnCapacityProvider) validateSetInstanceRequirementsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnCapacityProvider_InstanceRequirementsProperty:
		val := val.(*CfnCapacityProvider_InstanceRequirementsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnCapacityProvider_InstanceRequirementsProperty:
		val_ := val.(CfnCapacityProvider_InstanceRequirementsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnCapacityProvider_InstanceRequirementsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnCapacityProvider) validateSetPermissionsConfigParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnCapacityProvider_CapacityProviderPermissionsConfigProperty:
		val := val.(*CfnCapacityProvider_CapacityProviderPermissionsConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnCapacityProvider_CapacityProviderPermissionsConfigProperty:
		val_ := val.(CfnCapacityProvider_CapacityProviderPermissionsConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnCapacityProvider_CapacityProviderPermissionsConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnCapacityProvider) validateSetTagsParameters(val *[]*awscdk.CfnTag) error {
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func (j *jsiiProxy_CfnCapacityProvider) validateSetVpcConfigParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnCapacityProvider_CapacityProviderVpcConfigProperty:
		val := val.(*CfnCapacityProvider_CapacityProviderVpcConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnCapacityProvider_CapacityProviderVpcConfigProperty:
		val_ := val.(CfnCapacityProvider_CapacityProviderVpcConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnCapacityProvider_CapacityProviderVpcConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnCapacityProviderParameters(scope constructs.Construct, id *string, props *CfnCapacityProviderProps) error {
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

