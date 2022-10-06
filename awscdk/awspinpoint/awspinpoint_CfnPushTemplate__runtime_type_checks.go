//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awspinpoint

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CfnPushTemplate) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnPushTemplate) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnPushTemplate_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnPushTemplate_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnPushTemplate_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetAdmParameters(val interface{}) error {
	switch val.(type) {
	case *CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val := val.(*CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val_ := val.(CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnPushTemplate_AndroidPushNotificationTemplateProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetApnsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnPushTemplate_APNSPushNotificationTemplateProperty:
		val := val.(*CfnPushTemplate_APNSPushNotificationTemplateProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnPushTemplate_APNSPushNotificationTemplateProperty:
		val_ := val.(CfnPushTemplate_APNSPushNotificationTemplateProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnPushTemplate_APNSPushNotificationTemplateProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetBaiduParameters(val interface{}) error {
	switch val.(type) {
	case *CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val := val.(*CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val_ := val.(CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnPushTemplate_AndroidPushNotificationTemplateProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetDefaultParameters(val interface{}) error {
	switch val.(type) {
	case *CfnPushTemplate_DefaultPushNotificationTemplateProperty:
		val := val.(*CfnPushTemplate_DefaultPushNotificationTemplateProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnPushTemplate_DefaultPushNotificationTemplateProperty:
		val_ := val.(CfnPushTemplate_DefaultPushNotificationTemplateProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnPushTemplate_DefaultPushNotificationTemplateProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetGcmParameters(val interface{}) error {
	switch val.(type) {
	case *CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val := val.(*CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnPushTemplate_AndroidPushNotificationTemplateProperty:
		val_ := val.(CfnPushTemplate_AndroidPushNotificationTemplateProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnPushTemplate_AndroidPushNotificationTemplateProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnPushTemplate) validateSetTemplateNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCfnPushTemplateParameters(scope awscdk.Construct, id *string, props *CfnPushTemplateProps) error {
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

