//go:build !no_runtime_type_checking

package awswafv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnWebACL) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnWebACL) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnWebACL_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnWebACL_IsCfnResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnWebACL_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetAssociationConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_AssociationConfigProperty:
		val := val.(*CfnWebACL_AssociationConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_AssociationConfigProperty:
		val_ := val.(CfnWebACL_AssociationConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_AssociationConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetCaptchaConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_CaptchaConfigProperty:
		val := val.(*CfnWebACL_CaptchaConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_CaptchaConfigProperty:
		val_ := val.(CfnWebACL_CaptchaConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_CaptchaConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetChallengeConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_ChallengeConfigProperty:
		val := val.(*CfnWebACL_ChallengeConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_ChallengeConfigProperty:
		val_ := val.(CfnWebACL_ChallengeConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_ChallengeConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetCustomResponseBodiesParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]interface{}:
		val := val.(*map[string]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnWebACL_CustomResponseBodyProperty:
				v := v.(*CfnWebACL_CustomResponseBodyProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnWebACL_CustomResponseBodyProperty:
				v_ := v.(CfnWebACL_CustomResponseBodyProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_CustomResponseBodyProperty; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	case map[string]interface{}:
		val_ := val.(map[string]interface{})
		val := &val_
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnWebACL_CustomResponseBodyProperty:
				v := v.(*CfnWebACL_CustomResponseBodyProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnWebACL_CustomResponseBodyProperty:
				v_ := v.(CfnWebACL_CustomResponseBodyProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_CustomResponseBodyProperty; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *map[string]interface{}; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetDataProtectionConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_DataProtectionConfigProperty:
		val := val.(*CfnWebACL_DataProtectionConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_DataProtectionConfigProperty:
		val_ := val.(CfnWebACL_DataProtectionConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_DataProtectionConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetDefaultActionParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_DefaultActionProperty:
		val := val.(*CfnWebACL_DefaultActionProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_DefaultActionProperty:
		val_ := val.(CfnWebACL_DefaultActionProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_DefaultActionProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetOnSourceDDoSProtectionConfigParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_OnSourceDDoSProtectionConfigProperty:
		val := val.(*CfnWebACL_OnSourceDDoSProtectionConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_OnSourceDDoSProtectionConfigProperty:
		val_ := val.(CfnWebACL_OnSourceDDoSProtectionConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_OnSourceDDoSProtectionConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetRulesParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *[]interface{}:
		val := val.(*[]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnWebACL_RuleProperty:
				v := v.(*CfnWebACL_RuleProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnWebACL_RuleProperty:
				v_ := v.(CfnWebACL_RuleProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_RuleProperty; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	case []interface{}:
		val_ := val.([]interface{})
		val := &val_
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnWebACL_RuleProperty:
				v := v.(*CfnWebACL_RuleProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnWebACL_RuleProperty:
				v_ := v.(CfnWebACL_RuleProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_RuleProperty; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *[]interface{}; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetScopeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetTagsRawParameters(val *[]*awscdk.CfnTag) error {
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func (j *jsiiProxy_CfnWebACL) validateSetVisibilityConfigParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnWebACL_VisibilityConfigProperty:
		val := val.(*CfnWebACL_VisibilityConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnWebACL_VisibilityConfigProperty:
		val_ := val.(CfnWebACL_VisibilityConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnWebACL_VisibilityConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnWebACLParameters(scope constructs.Construct, id *string, props *CfnWebACLProps) error {
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

