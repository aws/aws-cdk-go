package awsamplifyuibuilder


// The `ComponentEvent` property specifies the configuration of an event.
//
// You can bind an event and a corresponding action to a `Component` or a `ComponentChild` . A button click is an example of an event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentEventProperty := &componentEventProperty{
//   	action: jsii.String("action"),
//   	parameters: &actionParametersProperty{
//   		anchor: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		fields: map[string]interface{}{
//   			"fieldsKey": &componentPropertyProperty{
//   				"bindingProperties": &componentPropertyBindingPropertiesProperty{
//   					"property": jsii.String("property"),
//
//   					// the properties below are optional
//   					"field": jsii.String("field"),
//   				},
//   				"bindings": map[string]interface{}{
//   					"bindingsKey": &FormBindingElementProperty{
//   						"element": jsii.String("element"),
//   						"property": jsii.String("property"),
//   					},
//   				},
//   				"collectionBindingProperties": &componentPropertyBindingPropertiesProperty{
//   					"property": jsii.String("property"),
//
//   					// the properties below are optional
//   					"field": jsii.String("field"),
//   				},
//   				"componentName": jsii.String("componentName"),
//   				"concat": []interface{}{
//   					componentPropertyProperty_,
//   				},
//   				"condition": &componentConditionPropertyProperty{
//   					"else": componentPropertyProperty_,
//   					"field": jsii.String("field"),
//   					"operand": jsii.String("operand"),
//   					"operandType": jsii.String("operandType"),
//   					"operator": jsii.String("operator"),
//   					"property": jsii.String("property"),
//   					"then": componentPropertyProperty_,
//   				},
//   				"configured": jsii.Boolean(false),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"event": jsii.String("event"),
//   				"importedValue": jsii.String("importedValue"),
//   				"model": jsii.String("model"),
//   				"property": jsii.String("property"),
//   				"type": jsii.String("type"),
//   				"userAttribute": jsii.String("userAttribute"),
//   				"value": jsii.String("value"),
//   			},
//   		},
//   		global: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		id: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		model: jsii.String("model"),
//   		state: &mutationActionSetStateParameterProperty{
//   			componentName: jsii.String("componentName"),
//   			property: jsii.String("property"),
//   			set: &componentPropertyProperty{
//   				bindingProperties: &componentPropertyBindingPropertiesProperty{
//   					property: jsii.String("property"),
//
//   					// the properties below are optional
//   					field: jsii.String("field"),
//   				},
//   				bindings: map[string]interface{}{
//   					"bindingsKey": &FormBindingElementProperty{
//   						"element": jsii.String("element"),
//   						"property": jsii.String("property"),
//   					},
//   				},
//   				collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   					property: jsii.String("property"),
//
//   					// the properties below are optional
//   					field: jsii.String("field"),
//   				},
//   				componentName: jsii.String("componentName"),
//   				concat: []interface{}{
//   					componentPropertyProperty_,
//   				},
//   				condition: &componentConditionPropertyProperty{
//   					else: componentPropertyProperty_,
//   					field: jsii.String("field"),
//   					operand: jsii.String("operand"),
//   					operandType: jsii.String("operandType"),
//   					operator: jsii.String("operator"),
//   					property: jsii.String("property"),
//   					then: componentPropertyProperty_,
//   				},
//   				configured: jsii.Boolean(false),
//   				defaultValue: jsii.String("defaultValue"),
//   				event: jsii.String("event"),
//   				importedValue: jsii.String("importedValue"),
//   				model: jsii.String("model"),
//   				property: jsii.String("property"),
//   				type: jsii.String("type"),
//   				userAttribute: jsii.String("userAttribute"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		target: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		type: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		url: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnComponent_ComponentEventProperty struct {
	// The action to perform when a specific event is raised.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Describes information about the action.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

