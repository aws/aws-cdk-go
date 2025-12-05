//go:build !no_runtime_type_checking

package awsfsx

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnS3AccessPointAttachment) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnS3AccessPointAttachment_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnS3AccessPointAttachment_IsCfnResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnS3AccessPointAttachment_IsCfnS3AccessPointAttachmentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnS3AccessPointAttachment_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnS3AccessPointAttachment) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnS3AccessPointAttachment) validateSetOntapConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty:
		val := val.(*CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty:
		val_ := val.(CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnS3AccessPointAttachment_S3AccessPointOntapConfigurationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnS3AccessPointAttachment) validateSetOpenZfsConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty:
		val := val.(*CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty:
		val_ := val.(CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnS3AccessPointAttachment_S3AccessPointOpenZFSConfigurationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnS3AccessPointAttachment) validateSetS3AccessPointParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *CfnS3AccessPointAttachment_S3AccessPointProperty:
		val := val.(*CfnS3AccessPointAttachment_S3AccessPointProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnS3AccessPointAttachment_S3AccessPointProperty:
		val_ := val.(CfnS3AccessPointAttachment_S3AccessPointProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *CfnS3AccessPointAttachment_S3AccessPointProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnS3AccessPointAttachment) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCfnS3AccessPointAttachmentParameters(scope constructs.Construct, id *string, props *CfnS3AccessPointAttachmentProps) error {
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

