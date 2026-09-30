//go:build !no_runtime_type_checking

package awscustomerprofiles

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscustomerprofiles"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnRecommenderSchema) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateAddResourceDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	if strength == "" {
		return fmt.Errorf("parameter strength is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateCfnPropertyNameParameters(cdkPropertyName *string) error {
	if cdkPropertyName == nil {
		return fmt.Errorf("parameter cdkPropertyName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateRemoveResourceDependencyParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if newTarget == nil {
		return fmt.Errorf("parameter newTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnRecommenderSchema) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnRecommenderSchema_ArnForRecommenderSchemaParameters(resource interfacesawscustomerprofiles.IRecommenderSchemaRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateCfnRecommenderSchema_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnRecommenderSchema_IsCfnRecommenderSchemaParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnRecommenderSchema_IsCfnResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnRecommenderSchema_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnRecommenderSchema) validateSetDomainNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnRecommenderSchema) validateSetFieldsParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]interface{}:
		val := val.(*map[string]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case awscdk.IResolvable:
				// ok
			case *[]interface{}:
				v := v.(*[]interface{})
				for idx_4c9448, v := range *v {
					switch v.(type) {
					case awscdk.IResolvable:
						// ok
					case *CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v := v.(*CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					case CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v_ := v.(CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						v := &v_
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					default:
						if !_jsii_.IsAnonymousProxy(v) {
							return fmt.Errorf("parameter val[%#v][%#v] must be one of the allowed types: awscdk.IResolvable, *CfnRecommenderSchema_RecommenderSchemaFieldProperty; received %#v (a %T)", idx_97dfc6, idx_4c9448, v, v)
						}
					}
				}
			case []interface{}:
				v_ := v.([]interface{})
				v := &v_
				for idx_4c9448, v := range *v {
					switch v.(type) {
					case awscdk.IResolvable:
						// ok
					case *CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v := v.(*CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					case CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v_ := v.(CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						v := &v_
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					default:
						if !_jsii_.IsAnonymousProxy(v) {
							return fmt.Errorf("parameter val[%#v][%#v] must be one of the allowed types: awscdk.IResolvable, *CfnRecommenderSchema_RecommenderSchemaFieldProperty; received %#v (a %T)", idx_97dfc6, idx_4c9448, v, v)
						}
					}
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *[]interface{}; received %#v (a %T)", idx_97dfc6, v, v)
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
			case *[]interface{}:
				v := v.(*[]interface{})
				for idx_4c9448, v := range *v {
					switch v.(type) {
					case awscdk.IResolvable:
						// ok
					case *CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v := v.(*CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					case CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v_ := v.(CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						v := &v_
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					default:
						if !_jsii_.IsAnonymousProxy(v) {
							return fmt.Errorf("parameter val[%#v][%#v] must be one of the allowed types: awscdk.IResolvable, *CfnRecommenderSchema_RecommenderSchemaFieldProperty; received %#v (a %T)", idx_97dfc6, idx_4c9448, v, v)
						}
					}
				}
			case []interface{}:
				v_ := v.([]interface{})
				v := &v_
				for idx_4c9448, v := range *v {
					switch v.(type) {
					case awscdk.IResolvable:
						// ok
					case *CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v := v.(*CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					case CfnRecommenderSchema_RecommenderSchemaFieldProperty:
						v_ := v.(CfnRecommenderSchema_RecommenderSchemaFieldProperty)
						v := &v_
						if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v][%#v]", idx_97dfc6, idx_4c9448) }); err != nil {
							return err
						}
					default:
						if !_jsii_.IsAnonymousProxy(v) {
							return fmt.Errorf("parameter val[%#v][%#v] must be one of the allowed types: awscdk.IResolvable, *CfnRecommenderSchema_RecommenderSchemaFieldProperty; received %#v (a %T)", idx_97dfc6, idx_4c9448, v, v)
						}
					}
				}
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: awscdk.IResolvable, *[]interface{}; received %#v (a %T)", idx_97dfc6, v, v)
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

func (j *jsiiProxy_CfnRecommenderSchema) validateSetRecommenderSchemaNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnRecommenderSchema) validateSetTagsParameters(val *[]*CfnRecommenderSchema_TagsItemsProperty) error {
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewCfnRecommenderSchemaParameters(scope constructs.Construct, id *string, props *CfnRecommenderSchemaProps) error {
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

