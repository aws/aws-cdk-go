//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnSAMLProvider) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnSAMLProvider) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnSAMLProvider_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnSAMLProvider_IsCfnResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnSAMLProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnSAMLProvider) validateSetPrivateKeyListParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *[]interface{}:
		val := val.(*[]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnSAMLProvider_SAMLPrivateKeyProperty:
				v := v.(*CfnSAMLProvider_SAMLPrivateKeyProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnSAMLProvider_SAMLPrivateKeyProperty:
				v_ := v.(CfnSAMLProvider_SAMLPrivateKeyProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnSAMLProvider_SAMLPrivateKeyProperty; received %#v (a %T)", idx_97dfc6, v, v)
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
			case *CfnSAMLProvider_SAMLPrivateKeyProperty:
				v := v.(*CfnSAMLProvider_SAMLPrivateKeyProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnSAMLProvider_SAMLPrivateKeyProperty:
				v_ := v.(CfnSAMLProvider_SAMLPrivateKeyProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnSAMLProvider_SAMLPrivateKeyProperty; received %#v (a %T)", idx_97dfc6, v, v)
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

func (j *jsiiProxy_CfnSAMLProvider) validateSetTagsRawParameters(val *[]*awscdk.CfnTag) error {
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewCfnSAMLProviderParameters(scope constructs.Construct, id *string, props *CfnSAMLProviderProps) error {
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

