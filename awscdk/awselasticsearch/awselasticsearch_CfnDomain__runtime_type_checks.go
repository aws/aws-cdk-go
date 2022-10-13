//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awselasticsearch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnDomain) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDomain) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnDomain_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnDomain_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnDomain_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetAccessPoliciesParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetAdvancedOptionsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]*string:
		// ok
	case map[string]*string:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *map[string]*string; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetAdvancedSecurityOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_AdvancedSecurityOptionsInputProperty:
		val := val.(*CfnDomain_AdvancedSecurityOptionsInputProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_AdvancedSecurityOptionsInputProperty:
		val_ := val.(CfnDomain_AdvancedSecurityOptionsInputProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_AdvancedSecurityOptionsInputProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetCognitoOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_CognitoOptionsProperty:
		val := val.(*CfnDomain_CognitoOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_CognitoOptionsProperty:
		val_ := val.(CfnDomain_CognitoOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_CognitoOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetDomainEndpointOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_DomainEndpointOptionsProperty:
		val := val.(*CfnDomain_DomainEndpointOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_DomainEndpointOptionsProperty:
		val_ := val.(CfnDomain_DomainEndpointOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_DomainEndpointOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetEbsOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_EBSOptionsProperty:
		val := val.(*CfnDomain_EBSOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_EBSOptionsProperty:
		val_ := val.(CfnDomain_EBSOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_EBSOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetElasticsearchClusterConfigParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_ElasticsearchClusterConfigProperty:
		val := val.(*CfnDomain_ElasticsearchClusterConfigProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_ElasticsearchClusterConfigProperty:
		val_ := val.(CfnDomain_ElasticsearchClusterConfigProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_ElasticsearchClusterConfigProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetEncryptionAtRestOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_EncryptionAtRestOptionsProperty:
		val := val.(*CfnDomain_EncryptionAtRestOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_EncryptionAtRestOptionsProperty:
		val_ := val.(CfnDomain_EncryptionAtRestOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_EncryptionAtRestOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetLogPublishingOptionsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]interface{}:
		val := val.(*map[string]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnDomain_LogPublishingOptionProperty:
				v := v.(*CfnDomain_LogPublishingOptionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnDomain_LogPublishingOptionProperty:
				v_ := v.(CfnDomain_LogPublishingOptionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnDomain_LogPublishingOptionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	case map[string]interface{}:
		val_ := val.(map[string]interface{})
		val := &val_
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnDomain_LogPublishingOptionProperty:
				v := v.(*CfnDomain_LogPublishingOptionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnDomain_LogPublishingOptionProperty:
				v_ := v.(CfnDomain_LogPublishingOptionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnDomain_LogPublishingOptionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
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

func (j *jsiiProxy_CfnDomain) validateSetNodeToNodeEncryptionOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_NodeToNodeEncryptionOptionsProperty:
		val := val.(*CfnDomain_NodeToNodeEncryptionOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_NodeToNodeEncryptionOptionsProperty:
		val_ := val.(CfnDomain_NodeToNodeEncryptionOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_NodeToNodeEncryptionOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetSnapshotOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_SnapshotOptionsProperty:
		val := val.(*CfnDomain_SnapshotOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_SnapshotOptionsProperty:
		val_ := val.(CfnDomain_SnapshotOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_SnapshotOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnDomain) validateSetVpcOptionsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnDomain_VPCOptionsProperty:
		val := val.(*CfnDomain_VPCOptionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnDomain_VPCOptionsProperty:
		val_ := val.(CfnDomain_VPCOptionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnDomain_VPCOptionsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnDomainParameters(scope constructs.Construct, id *string, props *CfnDomainProps) error {
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

