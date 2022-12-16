package awsamplifyuibuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AmplifyUIBuilder::Component`.
//
// The AWS::AmplifyUIBuilder::Component resource specifies a component within an Amplify app. A component is a user interface (UI) element that you can customize. Use `ComponentChild` to configure an instance of a `Component` . A `ComponentChild` instance inherits the configuration of the main `Component` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//   var overrides interface{}
//   var predicateProperty_ predicateProperty
//
//   cfnComponent := awscdk.Aws_amplifyuibuilder.NewCfnComponent(this, jsii.String("MyCfnComponent"), &cfnComponentProps{
//   	bindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &ComponentBindingPropertiesValueProperty{
//   			"bindingProperties": &ComponentBindingPropertiesValuePropertiesProperty{
//   				"bucket": jsii.String("bucket"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"field": jsii.String("field"),
//   				"key": jsii.String("key"),
//   				"model": jsii.String("model"),
//   				"predicates": []interface{}{
//   					&predicateProperty{
//   						"and": []interface{}{
//   							predicateProperty_,
//   						},
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operator": jsii.String("operator"),
//   						"or": []interface{}{
//   							predicateProperty_,
//   						},
//   					},
//   				},
//   				"userAttribute": jsii.String("userAttribute"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	componentType: jsii.String("componentType"),
//   	name: jsii.String("name"),
//   	overrides: overrides,
//   	properties: map[string]interface{}{
//   		"propertiesKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	variants: []interface{}{
//   		&componentVariantProperty{
//   			overrides: overrides,
//   			variantValues: map[string]*string{
//   				"variantValuesKey": jsii.String("variantValues"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	children: []interface{}{
//   		&componentChildProperty{
//   			componentType: jsii.String("componentType"),
//   			name: jsii.String("name"),
//   			properties: map[string]interface{}{
//   				"propertiesKey": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			events: map[string]interface{}{
//   				"eventsKey": &ComponentEventProperty{
//   					"action": jsii.String("action"),
//   					"parameters": &ActionParametersProperty{
//   						"anchor": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"fields": map[string]interface{}{
//   							"fieldsKey": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"global": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"id": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"model": jsii.String("model"),
//   						"state": &MutationActionSetStateParameterProperty{
//   							"componentName": jsii.String("componentName"),
//   							"property": jsii.String("property"),
//   							"set": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"target": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"type": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"url": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	collectionProperties: map[string]interface{}{
//   		"collectionPropertiesKey": &ComponentDataConfigurationProperty{
//   			"model": jsii.String("model"),
//
//   			// the properties below are optional
//   			"identifiers": []*string{
//   				jsii.String("identifiers"),
//   			},
//   			"predicate": &predicateProperty{
//   				"and": []interface{}{
//   					predicateProperty_,
//   				},
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operator": jsii.String("operator"),
//   				"or": []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   			"sort": []interface{}{
//   				&SortPropertyProperty{
//   					"direction": jsii.String("direction"),
//   					"field": jsii.String("field"),
//   				},
//   			},
//   		},
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &ComponentEventProperty{
//   			"action": jsii.String("action"),
//   			"parameters": &ActionParametersProperty{
//   				"anchor": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"fields": map[string]interface{}{
//   					"fieldsKey": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"global": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"id": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"model": jsii.String("model"),
//   				"state": &MutationActionSetStateParameterProperty{
//   					"componentName": jsii.String("componentName"),
//   					"property": jsii.String("property"),
//   					"set": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"target": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"type": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"url": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	schemaVersion: jsii.String("schemaVersion"),
//   	sourceId: jsii.String("sourceId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique ID for the Amplify app.
	AttrAppId() *string
	// The name of the backend environment that is a part of the Amplify app.
	AttrEnvironmentName() *string
	// The unique ID of the component.
	AttrId() *string
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	BindingProperties() interface{}
	SetBindingProperties(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A list of the component's `ComponentChild` instances.
	Children() interface{}
	SetChildren(val interface{})
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties() interface{}
	SetCollectionProperties(val interface{})
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType() *string
	SetComponentType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events() interface{}
	SetEvents(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the component.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides() interface{}
	SetOverrides(val interface{})
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties() interface{}
	SetProperties(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schema version of the component when it was imported.
	SchemaVersion() *string
	SetSchemaVersion(val *string)
	// The unique ID of the component in its original source system, such as Figma.
	SourceId() *string
	SetSourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// One or more key-value pairs to use when tagging the component.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants() interface{}
	SetVariants(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponent
type jsiiProxy_CfnComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponent) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) BindingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Children() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CollectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) ComponentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Variants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variants",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent(scope constructs.Construct, id *string, props *CfnComponentProps) CfnComponent {
	_init_.Initialize()

	if err := validateNewCfnComponentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComponent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent_Override(c CfnComponent, scope constructs.Construct, id *string, props *CfnComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponent)SetBindingProperties(val interface{}) {
	if err := j.validateSetBindingPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetChildren(val interface{}) {
	if err := j.validateSetChildrenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"children",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetCollectionProperties(val interface{}) {
	if err := j.validateSetCollectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetComponentType(val *string) {
	if err := j.validateSetComponentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentType",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetEvents(val interface{}) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetOverrides(val interface{}) {
	if err := j.validateSetOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetProperties(val interface{}) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetSchemaVersion(val *string) {
	_jsii_.Set(
		j,
		"schemaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetSourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetVariants(val interface{}) {
	if err := j.validateSetVariantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variants",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponent) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponent) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponent) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponent) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

