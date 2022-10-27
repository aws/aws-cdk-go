//go:build !no_runtime_type_checking

package awsconnect

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnInstanceStorageConfig) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnInstanceStorageConfig_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnInstanceStorageConfig_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnInstanceStorageConfig_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetInstanceArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetKinesisFirehoseConfigParameters(val interface{}) error {
	switch val.(type) {
	case *CfnInstanceStorageConfig_KinesisFirehoseConfigProperty:
		val := val.(*CfnInstanceStorageConfig_KinesisFirehoseConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnInstanceStorageConfig_KinesisFirehoseConfigProperty:
		val_ := val.(CfnInstanceStorageConfig_KinesisFirehoseConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnInstanceStorageConfig_KinesisFirehoseConfigProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetKinesisStreamConfigParameters(val interface{}) error {
	switch val.(type) {
	case *CfnInstanceStorageConfig_KinesisStreamConfigProperty:
		val := val.(*CfnInstanceStorageConfig_KinesisStreamConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnInstanceStorageConfig_KinesisStreamConfigProperty:
		val_ := val.(CfnInstanceStorageConfig_KinesisStreamConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnInstanceStorageConfig_KinesisStreamConfigProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetKinesisVideoStreamConfigParameters(val interface{}) error {
	switch val.(type) {
	case *CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty:
		val := val.(*CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty:
		val_ := val.(CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnInstanceStorageConfig_KinesisVideoStreamConfigProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetResourceTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetS3ConfigParameters(val interface{}) error {
	switch val.(type) {
	case *CfnInstanceStorageConfig_S3ConfigProperty:
		val := val.(*CfnInstanceStorageConfig_S3ConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnInstanceStorageConfig_S3ConfigProperty:
		val_ := val.(CfnInstanceStorageConfig_S3ConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnInstanceStorageConfig_S3ConfigProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnInstanceStorageConfig) validateSetStorageTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCfnInstanceStorageConfigParameters(scope awscdk.Construct, id *string, props *CfnInstanceStorageConfigProps) error {
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

