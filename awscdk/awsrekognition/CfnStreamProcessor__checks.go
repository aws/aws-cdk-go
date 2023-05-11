//go:build !no_runtime_type_checking

package awsrekognition

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnStreamProcessor) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnStreamProcessor) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnStreamProcessor_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnStreamProcessor_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnStreamProcessor_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetBoundingBoxRegionsOfInterestParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *[]interface{}:
		val := val.(*[]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *CfnStreamProcessor_BoundingBoxProperty:
				v := v.(*CfnStreamProcessor_BoundingBoxProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnStreamProcessor_BoundingBoxProperty:
				v_ := v.(CfnStreamProcessor_BoundingBoxProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_BoundingBoxProperty; received %#v (a %T)", idx_97dfc6, v, v)
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
			case *CfnStreamProcessor_BoundingBoxProperty:
				v := v.(*CfnStreamProcessor_BoundingBoxProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnStreamProcessor_BoundingBoxProperty:
				v_ := v.(CfnStreamProcessor_BoundingBoxProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_BoundingBoxProperty; received %#v (a %T)", idx_97dfc6, v, v)
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

func (j *jsiiProxy_CfnStreamProcessor) validateSetConnectedHomeSettingsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_ConnectedHomeSettingsProperty:
		val := val.(*CfnStreamProcessor_ConnectedHomeSettingsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_ConnectedHomeSettingsProperty:
		val_ := val.(CfnStreamProcessor_ConnectedHomeSettingsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_ConnectedHomeSettingsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetDataSharingPreferenceParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_DataSharingPreferenceProperty:
		val := val.(*CfnStreamProcessor_DataSharingPreferenceProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_DataSharingPreferenceProperty:
		val_ := val.(CfnStreamProcessor_DataSharingPreferenceProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_DataSharingPreferenceProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetFaceSearchSettingsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_FaceSearchSettingsProperty:
		val := val.(*CfnStreamProcessor_FaceSearchSettingsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_FaceSearchSettingsProperty:
		val_ := val.(CfnStreamProcessor_FaceSearchSettingsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_FaceSearchSettingsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetKinesisDataStreamParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_KinesisDataStreamProperty:
		val := val.(*CfnStreamProcessor_KinesisDataStreamProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_KinesisDataStreamProperty:
		val_ := val.(CfnStreamProcessor_KinesisDataStreamProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_KinesisDataStreamProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetKinesisVideoStreamParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_KinesisVideoStreamProperty:
		val := val.(*CfnStreamProcessor_KinesisVideoStreamProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_KinesisVideoStreamProperty:
		val_ := val.(CfnStreamProcessor_KinesisVideoStreamProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_KinesisVideoStreamProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetNotificationChannelParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_NotificationChannelProperty:
		val := val.(*CfnStreamProcessor_NotificationChannelProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_NotificationChannelProperty:
		val_ := val.(CfnStreamProcessor_NotificationChannelProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_NotificationChannelProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetPolygonRegionsOfInterestParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetRoleArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnStreamProcessor) validateSetS3DestinationParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnStreamProcessor_S3DestinationProperty:
		val := val.(*CfnStreamProcessor_S3DestinationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnStreamProcessor_S3DestinationProperty:
		val_ := val.(CfnStreamProcessor_S3DestinationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnStreamProcessor_S3DestinationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnStreamProcessorParameters(scope constructs.Construct, id *string, props *CfnStreamProcessorProps) error {
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

