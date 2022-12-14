//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnDeliveryStream) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDeliveryStream) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnDeliveryStream_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnDeliveryStream_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnDeliveryStream_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetAmazonOpenSearchServerlessDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetAmazonopensearchserviceDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetDeliveryStreamEncryptionConfigurationInputParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty:
		val := val.(*CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty:
		val_ := val.(CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetElasticsearchDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetExtendedS3DestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetHttpEndpointDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetKinesisStreamSourceConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_KinesisStreamSourceConfigurationProperty:
		val := val.(*CfnDeliveryStream_KinesisStreamSourceConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_KinesisStreamSourceConfigurationProperty:
		val_ := val.(CfnDeliveryStream_KinesisStreamSourceConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_KinesisStreamSourceConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetRedshiftDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_RedshiftDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_RedshiftDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_RedshiftDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_RedshiftDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_RedshiftDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetS3DestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_S3DestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_S3DestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_S3DestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_S3DestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_S3DestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDeliveryStream) validateSetSplunkDestinationConfigurationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDeliveryStream_SplunkDestinationConfigurationProperty:
		val := val.(*CfnDeliveryStream_SplunkDestinationConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDeliveryStream_SplunkDestinationConfigurationProperty:
		val_ := val.(CfnDeliveryStream_SplunkDestinationConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDeliveryStream_SplunkDestinationConfigurationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnDeliveryStreamParameters(scope constructs.Construct, id *string, props *CfnDeliveryStreamProps) error {
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

