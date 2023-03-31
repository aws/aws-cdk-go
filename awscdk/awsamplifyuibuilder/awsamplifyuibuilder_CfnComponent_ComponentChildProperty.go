package awsamplifyuibuilder


// The `ComponentChild` property specifies a nested UI configuration within a parent `Component` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentChildProperty := &componentChildProperty{
//   	componentType: jsii.String("componentType"),
//   	name: jsii.String("name"),
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
//   }
//
type CfnComponent_ComponentChildProperty struct {
	// The type of the child component.
	ComponentType *string `field:"required" json:"componentType" yaml:"componentType"`
	// The name of the child component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the properties of the child component.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// The list of `ComponentChild` instances for this component.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// Describes the events that can be raised on the child component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}

